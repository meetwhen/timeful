package db

import (
	"context"
	"io"
	"os"
	"strings"
	"testing"

	"github.com/jackc/pgx/v5/pgxpool"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"timeful/server/logger"
	"timeful/server/models"
	pgstore "timeful/server/postgres"
)

// initAccountAuthorityTestMongo connects the retained MongoDB store used to
// prove that a retained document is never served as account authority. It
// requires an isolated test database.
func initAccountAuthorityTestMongo(t *testing.T) {
	t.Helper()
	if logger.StdErr == nil {
		logger.Init(io.Discard)
	}
	if UsersCollection != nil {
		return
	}
	if os.Getenv("MONGODB_URI") == "" {
		t.Skip("MONGODB_URI is required for account authority tests")
	}
	database := os.Getenv("MONGODB_DATABASE")
	if database != "timeful-test" && !strings.HasPrefix(database, "timeful-test-") {
		t.Fatalf("MONGODB_DATABASE must be timeful-test or use a timeful-test- prefix; got %q", database)
	}
	Init()
}

func insertAccountAuthorityUser(t *testing.T, user models.User) {
	t.Helper()
	ctx := context.Background()
	if _, err := UsersCollection.InsertOne(ctx, user); err != nil {
		t.Fatalf("insert retained user: %v", err)
	}
	t.Cleanup(func() {
		if _, err := UsersCollection.DeleteOne(context.Background(), bson.M{"_id": user.Id}); err != nil {
			t.Errorf("delete retained user: %v", err)
		}
	})
}

// closedPostgresTestPool is a non-nil pool whose queries fail immediately, so a
// test can exercise the initialized-but-erroring PostgreSQL state without a
// server.
func closedPostgresTestPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	config, err := pgxpool.ParseConfig("postgres://timeful:timeful@127.0.0.1:1/timeful-test?sslmode=disable&connect_timeout=1")
	if err != nil {
		t.Fatal(err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		t.Fatal(err)
	}
	pool.Close()
	return pool
}

func accountsAuthorityTestPool(t *testing.T) *pgxpool.Pool {
	t.Helper()
	uri := os.Getenv("POSTGRES_APPLICATION_URI")
	if uri == "" {
		t.Skip("POSTGRES_APPLICATION_URI is required for the not-found case")
	}
	config, err := pgxpool.ParseConfig(uri)
	if err != nil {
		t.Fatal(err)
	}
	pool, err := pgxpool.NewWithConfig(context.Background(), config)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(pool.Close)
	return pool
}

// initAccountLookupPostgres connects the authoritative PostgreSQL store used by
// the account lookup tests, reusing an already initialized package pool when a
// sibling test has one.
func initAccountLookupPostgres(t *testing.T) {
	t.Helper()
	if os.Getenv("POSTGRES_APPLICATION_URI") == "" {
		t.Skip("POSTGRES_APPLICATION_URI is required for account lookup tests")
	}
	previous := pgstore.Pool
	var closePool func()
	if previous == nil {
		closePool = pgstore.Init()
	}
	t.Cleanup(func() {
		if closePool != nil {
			closePool()
		}
		pgstore.Pool = previous
	})
}

func deleteAccountLookupFixture(t *testing.T, externalUserID string) {
	t.Helper()
	repository, err := pgstore.DefaultRepository()
	if err != nil {
		t.Errorf("resolve repository for account cleanup: %v", err)
		return
	}
	ctx := context.Background()
	if err := repository.DeleteAccountByExternalUserID(ctx, externalUserID); err != nil {
		t.Errorf("delete account %s: %v", externalUserID, err)
	}
	if _, err := pgstore.Pool.Exec(ctx, `DELETE FROM platform_identities WHERE external_user_id = $1`, externalUserID); err != nil {
		t.Errorf("delete platform identity %s: %v", externalUserID, err)
	}
}

// TestGetUserByIdReturnsPostgresProfile proves that both account lookups return
// the authoritative PostgreSQL profile, including the usage counter, and never
// any calendar integration fields.
func TestGetUserByIdReturnsPostgresProfile(t *testing.T) {
	initAccountLookupPostgres(t)
	ctx := context.Background()

	repository, err := pgstore.DefaultRepository()
	if err != nil {
		t.Fatal(err)
	}

	email := "lookup-" + primitive.NewObjectID().Hex() + "@example.com"
	externalUserID := primitive.NewObjectID().Hex()
	account, created, err := repository.FindOrCreateAccountByEmail(ctx, email, externalUserID, pgstore.Account{
		Email:            email,
		FirstName:        "Postgres",
		LastName:         "Profile",
		Picture:          "https://postgres.example/picture.png",
		TimezoneOffset:   90,
		NumEventsCreated: 7,
	})
	if err != nil {
		t.Fatal(err)
	}
	if !created {
		t.Fatal("expected the lookup test to create a fresh account")
	}
	t.Cleanup(func() { deleteAccountLookupFixture(t, account.ExternalUserID) })

	assertProfile := func(lookup string, got *models.User) {
		t.Helper()
		if got == nil {
			t.Fatalf("%s returned no user", lookup)
		}
		if got.Id.Hex() != account.ExternalUserID || got.Email != email {
			t.Fatalf("%s returned the wrong account: %#v", lookup, got)
		}
		if got.FirstName != "Postgres" || got.LastName != "Profile" {
			t.Fatalf("%s returned the wrong name: %#v", lookup, got)
		}
		if got.Picture != "https://postgres.example/picture.png" {
			t.Fatalf("%s returned the wrong picture: %q", lookup, got.Picture)
		}
		if got.TimezoneOffset != 90 {
			t.Fatalf("%s returned the wrong timezone offset: %d", lookup, got.TimezoneOffset)
		}
		if got.NumEventsCreated != 7 {
			t.Fatalf("%s returned the wrong usage counter: %d", lookup, got.NumEventsCreated)
		}
		if got.CalendarAccounts != nil || got.CalendarOptions != nil || got.PrimaryAccountKey != nil || got.TokenOrigin != "" {
			t.Fatalf("%s returned calendar integration fields: %#v", lookup, got)
		}
	}

	assertProfile("GetUserByEmail", GetUserByEmail(email))
	assertProfile("GetUserById", GetUserById(account.ExternalUserID))
}

// TestGetUserByIdIgnoresRetainedDocument proves that a retained MongoDB user
// document is never served as account authority once the account no longer
// exists in PostgreSQL.
func TestGetUserByIdIgnoresRetainedDocument(t *testing.T) {
	initAccountAuthorityTestMongo(t)
	initAccountLookupPostgres(t)

	email := "retained-" + primitive.NewObjectID().Hex() + "@example.com"
	user := models.User{Id: primitive.NewObjectID(), Email: email, FirstName: "Retained", LastName: "Legacy"}
	insertAccountAuthorityUser(t, user)

	if got := GetUserById(user.Id.Hex()); got != nil {
		t.Fatalf("retained document must not be served by identifier, got %#v", got)
	}
	if got := GetUserByEmail(email); got != nil {
		t.Fatalf("retained document must not be served by email, got %#v", got)
	}
}

// TestGetUserByIdPostgresErrorReturnsNil proves that a PostgreSQL lookup failure
// never falls back to an inferred account profile.
func TestGetUserByIdPostgresErrorReturnsNil(t *testing.T) {
	initAccountAuthorityTestMongo(t)
	previousPool := pgstore.Pool
	t.Cleanup(func() { pgstore.Pool = previousPool })
	pgstore.Pool = closedPostgresTestPool(t)

	email := "retained-error-" + primitive.NewObjectID().Hex() + "@example.com"
	user := models.User{Id: primitive.NewObjectID(), Email: email, FirstName: "MongoAuthority", LastName: "ShouldNotWin"}
	insertAccountAuthorityUser(t, user)

	if got := GetUserById(user.Id.Hex()); got != nil {
		t.Fatalf("PostgreSQL error must yield no account by identifier, got %#v", got)
	}
	if got := GetUserByEmail(email); got != nil {
		t.Fatalf("PostgreSQL error must yield no account by email, got %#v", got)
	}
}

// TestGetUserByIdUnknownAccountReturnsNil proves that a genuine not-found
// returns no account instead of an inferred profile.
func TestGetUserByIdUnknownAccountReturnsNil(t *testing.T) {
	initAccountAuthorityTestMongo(t)
	previousPool := pgstore.Pool
	t.Cleanup(func() { pgstore.Pool = previousPool })
	pgstore.Pool = accountsAuthorityTestPool(t)

	email := "retained-legacy-" + primitive.NewObjectID().Hex() + "@example.com"
	user := models.User{Id: primitive.NewObjectID(), Email: email, FirstName: "LegacyOnly", LastName: "PreBackfill"}
	insertAccountAuthorityUser(t, user)

	if got := GetUserById(user.Id.Hex()); got != nil {
		t.Fatalf("genuine not-found must yield no account by identifier, got %#v", got)
	}
	if got := GetUserByEmail(email); got != nil {
		t.Fatalf("genuine not-found must yield no account by email, got %#v", got)
	}
}
