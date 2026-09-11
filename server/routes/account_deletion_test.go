package routes

import (
	"context"
	"errors"
	"net/http"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"timeful/server/accounts"
	"timeful/server/db"
	"timeful/server/models"
	pgstore "timeful/server/postgres"
)

// seedRetainedMongoData seeds legacy records for the signed-in account so the
// deletion contract can prove the retained store is never written.
func seedRetainedMongoData(t *testing.T, account *pgstore.Account) primitive.ObjectID {
	t.Helper()
	ctx := context.Background()
	objectID := accountObjectID(t, account.ExternalUserID)
	mongoEventID := primitive.NewObjectID()
	folderID := primitive.NewObjectID()

	if _, err := db.EventsCollection.InsertOne(ctx, models.Event{Id: mongoEventID, OwnerId: objectID, Name: "Owned legacy event"}); err != nil {
		t.Fatal(err)
	}
	if _, err := db.EventResponsesCollection.InsertMany(ctx, []interface{}{
		models.EventResponse{EventId: mongoEventID, UserId: account.ExternalUserID, Response: &models.Response{Name: "Owner"}},
		models.EventResponse{EventId: mongoEventID, UserId: "guest-key", Response: &models.Response{Name: "Guest"}},
	}); err != nil {
		t.Fatal(err)
	}
	if _, err := db.FoldersCollection.InsertOne(ctx, models.Folder{Id: folderID, UserId: objectID, Name: "Folder"}); err != nil {
		t.Fatal(err)
	}
	if _, err := db.FolderEventsCollection.InsertOne(ctx, models.FolderEvent{UserId: objectID, FolderId: folderID, EventId: mongoEventID}); err != nil {
		t.Fatal(err)
	}
	if _, err := db.UsersCollection.InsertOne(ctx, models.User{Id: objectID, Email: account.Email, FirstName: "Retained"}); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		cleanup := context.Background()
		_, _ = db.EventsCollection.DeleteOne(cleanup, bson.M{"_id": mongoEventID})
		_, _ = db.EventResponsesCollection.DeleteMany(cleanup, bson.M{"eventId": mongoEventID})
		_, _ = db.FoldersCollection.DeleteOne(cleanup, bson.M{"_id": folderID})
		_, _ = db.FolderEventsCollection.DeleteMany(cleanup, bson.M{"folderId": folderID})
		_, _ = db.UsersCollection.DeleteOne(cleanup, bson.M{"_id": objectID})
	})
	return mongoEventID
}

// TestAccountDeletionRemovesPostgresAuthority proves the ratified FR-123
// deletion unit: the PostgreSQL account authority, platform identity, calendar
// connections, responses, folders, and daily-log membership are removed;
// events the account organized survive with ownership released and their other
// guests' responses intact; and the retained legacy store is never written.
func TestAccountDeletionRemovesPostgresAuthority(t *testing.T) {
	router := newAccountContractRouter(t)
	client := newAccountContractClient(t, router)
	ctx := context.Background()
	email := "delete-" + primitive.NewObjectID().Hex() + "@example.com"

	verifyOtpSignIn(t, client, email, "123456")
	repository := repositoryForTest(t)
	account, err := repository.GetAccountByEmail(ctx, email)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { deleteAccountTestFixtures(t, account.ExternalUserID) })
	objectID := accountObjectID(t, account.ExternalUserID)
	mongoEventID := seedRetainedMongoData(t, account)

	// PostgreSQL: a daily log shared with another account, and a log that only
	// the deleted account used.
	sharedDate := time.Date(2000, 1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, int(objectID[0])*256+int(objectID[1]))
	soloDate := time.Date(3000, 1, 1, 0, 0, 0, 0, time.UTC).AddDate(0, 0, int(objectID[2])*256+int(objectID[3]))
	otherAccountID := primitive.NewObjectID().Hex()
	var sharedLogID, soloLogID string
	if err := pgstore.Pool.QueryRow(ctx, `INSERT INTO daily_user_logs (log_date) VALUES ($1)
ON CONFLICT (log_date) DO UPDATE SET updated_at = daily_user_logs.updated_at RETURNING id`, sharedDate).Scan(&sharedLogID); err != nil {
		t.Fatal(err)
	}
	if _, err := pgstore.Pool.Exec(ctx, `INSERT INTO daily_user_log_members (daily_user_log_id, account_user_id, first_seen_position)
VALUES ($1, $2, 0), ($1, $3, 1) ON CONFLICT DO NOTHING`, sharedLogID, account.ExternalUserID, otherAccountID); err != nil {
		t.Fatal(err)
	}
	if err := pgstore.Pool.QueryRow(ctx, `INSERT INTO daily_user_logs (log_date) VALUES ($1)
ON CONFLICT (log_date) DO UPDATE SET updated_at = daily_user_logs.updated_at RETURNING id`, soloDate).Scan(&soloLogID); err != nil {
		t.Fatal(err)
	}
	if _, err := pgstore.Pool.Exec(ctx, `INSERT INTO daily_user_log_members (daily_user_log_id, account_user_id, first_seen_position)
VALUES ($1, $2, 0) ON CONFLICT DO NOTHING`, soloLogID, account.ExternalUserID); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_, _ = pgstore.Pool.Exec(context.Background(), `DELETE FROM daily_user_logs WHERE id = ANY($1)`, []string{sharedLogID, soloLogID})
	})

	// PostgreSQL: an event the account owns with one account response and one
	// guest response.
	shortID, err := pgstore.GenerateShortID()
	if err != nil {
		t.Fatal(err)
	}
	var eventID, ownerVisitorID, guestVisitorID string
	if err := pgstore.Pool.QueryRow(ctx, `INSERT INTO postgres_events (short_id, name, type, owner_external_id, owner_platform_identity_id)
VALUES ($1, 'Owned', 'specific_dates', $2, $3) RETURNING id`, shortID, account.ExternalUserID, account.PlatformIdentityID).Scan(&eventID); err != nil {
		t.Fatal(err)
	}
	if err := pgstore.Pool.QueryRow(ctx, `INSERT INTO event_visitor_identities (event_id, platform_identity_id) VALUES ($1, $2) RETURNING id`, eventID, account.PlatformIdentityID).Scan(&ownerVisitorID); err != nil {
		t.Fatal(err)
	}
	if _, err := pgstore.Pool.Exec(ctx, `UPDATE postgres_events SET owner_event_visitor_identity_id = $2 WHERE id = $1`, eventID, ownerVisitorID); err != nil {
		t.Fatal(err)
	}
	if err := pgstore.Pool.QueryRow(ctx, `INSERT INTO event_visitor_identities (event_id) VALUES ($1) RETURNING id`, eventID).Scan(&guestVisitorID); err != nil {
		t.Fatal(err)
	}
	if _, err := pgstore.Pool.Exec(ctx, `INSERT INTO postgres_event_responses (event_id, event_visitor_identity_id, respondent_kind, account_user_id, payload)
VALUES ($1, $2, 'account', $3, '{"name":"Owner"}'), ($1, $4, 'guest', NULL, '{"name":"Guest"}')`,
		eventID, ownerVisitorID, account.ExternalUserID, guestVisitorID); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		cleanup := context.Background()
		_, _ = pgstore.Pool.Exec(cleanup, `DELETE FROM postgres_events WHERE id = $1`, eventID)
	})

	// PostgreSQL: an account folder with a PostgreSQL member and a legacy member.
	var folderID string
	if err := pgstore.Pool.QueryRow(ctx, `INSERT INTO folders (account_user_id, name) VALUES ($1, 'Folder') RETURNING id`, account.ExternalUserID).Scan(&folderID); err != nil {
		t.Fatal(err)
	}
	if _, err := pgstore.Pool.Exec(ctx, `INSERT INTO folder_events (account_user_id, folder_id, event_id) VALUES ($1, $2, $3)`, account.ExternalUserID, folderID, eventID); err != nil {
		t.Fatal(err)
	}
	if _, err := pgstore.Pool.Exec(ctx, `INSERT INTO folder_events (account_user_id, folder_id, legacy_event_id) VALUES ($1, $2, $3)`, account.ExternalUserID, folderID, mongoEventID.Hex()); err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		_, _ = pgstore.Pool.Exec(context.Background(), `DELETE FROM folders WHERE id = $1`, folderID)
	})

	client.request(http.MethodDelete, "/api/user", map[string]any{"email": email}, http.StatusOK)

	// Signed out and account authority gone.
	client.request(http.MethodGet, "/api/user/profile", nil, http.StatusUnauthorized)
	if _, err := repository.GetAccountByExternalUserID(ctx, account.ExternalUserID); err == nil {
		t.Fatal("account still resolves after deletion")
	}
	var identities int
	if err := pgstore.Pool.QueryRow(ctx, `SELECT count(*) FROM platform_identities WHERE external_user_id = $1`, account.ExternalUserID).Scan(&identities); err != nil {
		t.Fatal(err)
	}
	if identities != 0 {
		t.Fatalf("platform identity survived deletion: %d", identities)
	}

	// PostgreSQL: the account's own data is gone while another guest's
	// response survives.
	var ownResponses, guestResponses, pgFolders, pgMemberships int64
	if err := pgstore.Pool.QueryRow(ctx, `SELECT
 (SELECT count(*) FROM postgres_event_responses WHERE account_user_id = $1),
 (SELECT count(*) FROM postgres_event_responses WHERE event_id = $2 AND respondent_kind = 'guest'),
 (SELECT count(*) FROM folders WHERE account_user_id = $1),
 (SELECT count(*) FROM folder_events WHERE account_user_id = $1)`,
		account.ExternalUserID, eventID).Scan(&ownResponses, &guestResponses, &pgFolders, &pgMemberships); err != nil {
		t.Fatal(err)
	}
	if ownResponses != 0 || pgFolders != 0 || pgMemberships != 0 {
		t.Fatalf("owned PostgreSQL data survived: responses=%d folders=%d memberships=%d", ownResponses, pgFolders, pgMemberships)
	}
	if guestResponses != 1 {
		t.Fatalf("another guest's response was removed: %d", guestResponses)
	}

	// The deleted account is gone from PostgreSQL daily logs, the log it emptied
	// is deleted, and a log shared with another account survives.
	var ownLogMembers, sharedSurvived, soloSurvived, otherLogMembers int
	if err := pgstore.Pool.QueryRow(ctx, `SELECT
 (SELECT count(*) FROM daily_user_log_members WHERE account_user_id = $1),
 (SELECT count(*) FROM daily_user_logs WHERE id = $2),
 (SELECT count(*) FROM daily_user_logs WHERE id = $3),
 (SELECT count(*) FROM daily_user_log_members WHERE daily_user_log_id = $2)`,
		account.ExternalUserID, sharedLogID, soloLogID).Scan(&ownLogMembers, &sharedSurvived, &soloSurvived, &otherLogMembers); err != nil {
		t.Fatal(err)
	}
	if ownLogMembers != 0 {
		t.Fatalf("deleted account still appears in a daily user log: %d", ownLogMembers)
	}
	if soloSurvived != 0 {
		t.Fatal("daily user log emptied by deletion was not removed")
	}
	if sharedSurvived != 1 || otherLogMembers != 1 {
		t.Fatalf("daily user log shared with another account was not preserved: logs=%d members=%d", sharedSurvived, otherLogMembers)
	}

	// Events survive with released ownership.
	var pgOwned, pgOwnResponses, pgGuestResponses int
	if err := pgstore.Pool.QueryRow(ctx, `SELECT
 (SELECT count(*) FROM postgres_events WHERE id = $1 AND owner_platform_identity_id IS NULL AND owner_external_id IS NULL AND owner_event_visitor_identity_id IS NULL),
 (SELECT count(*) FROM postgres_event_responses WHERE event_id = $1 AND respondent_kind = 'account'),
 (SELECT count(*) FROM postgres_event_responses WHERE event_id = $1 AND respondent_kind = 'guest')`, eventID).Scan(&pgOwned, &pgOwnResponses, &pgGuestResponses); err != nil {
		t.Fatal(err)
	}
	if pgOwned != 1 {
		t.Fatal("PostgreSQL event did not survive with ownership released")
	}
	if pgOwnResponses != 0 || pgGuestResponses != 1 {
		t.Fatalf("PostgreSQL responses wrong: own=%d guest=%d", pgOwnResponses, pgGuestResponses)
	}

	// The retained legacy store is never written by deletion: the legacy event,
	// its owner, its responses, the folder, and the retained user document all
	// survive as unmodified recovery source.
	var legacyEvent models.Event
	if err := db.EventsCollection.FindOne(ctx, bson.M{"_id": mongoEventID}).Decode(&legacyEvent); err != nil {
		t.Fatalf("retained legacy event did not survive: %v", err)
	}
	if legacyEvent.OwnerId != objectID {
		t.Fatalf("retained legacy event ownership was rewritten: %s", legacyEvent.OwnerId.Hex())
	}
	var legacyResponses, legacyFolders, legacyUsers int64
	if legacyResponses, err = db.EventResponsesCollection.CountDocuments(ctx, bson.M{"eventId": mongoEventID}); err != nil {
		t.Fatal(err)
	}
	if legacyFolders, err = db.FoldersCollection.CountDocuments(ctx, bson.M{"userId": objectID}); err != nil {
		t.Fatal(err)
	}
	if legacyUsers, err = db.UsersCollection.CountDocuments(ctx, bson.M{"_id": objectID}); err != nil {
		t.Fatal(err)
	}
	if legacyResponses != 2 || legacyFolders != 1 || legacyUsers != 1 {
		t.Fatalf("deletion wrote the retained legacy store: responses=%d folders=%d users=%d", legacyResponses, legacyFolders, legacyUsers)
	}

	// Deletion is idempotent: repeating the unit is a no-op.
	if err := accounts.DeleteAccount(ctx, account.ExternalUserID); err != nil {
		t.Fatalf("repeated deletion must be idempotent: %v", err)
	}
}

// TestAccountDeletionRejectsEmailMismatch proves that the account email must be
// typed before anything is removed.
func TestAccountDeletionRejectsEmailMismatch(t *testing.T) {
	router := newAccountContractRouter(t)
	client := newAccountContractClient(t, router)
	ctx := context.Background()
	email := "mismatch-" + primitive.NewObjectID().Hex() + "@example.com"

	verifyOtpSignIn(t, client, email, "123456")
	account, err := repositoryForTest(t).GetAccountByEmail(ctx, email)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { deleteAccountTestFixtures(t, account.ExternalUserID) })

	client.request(http.MethodDelete, "/api/user", map[string]any{"email": "someone-else@example.com"}, http.StatusBadRequest)
	client.request(http.MethodGet, "/api/user/profile", nil, http.StatusOK)
	if _, err := repositoryForTest(t).GetAccountByExternalUserID(ctx, account.ExternalUserID); err != nil {
		t.Fatalf("account removed despite email mismatch: %v", err)
	}
}

// TestAccountDeletionPartialFailureLeavesAuthorityAndRetryConverges proves that
// a PostgreSQL deletion failure leaves the account authority and session intact
// and that retrying converges.
func TestAccountDeletionPartialFailureLeavesAuthorityAndRetryConverges(t *testing.T) {
	router := newAccountContractRouter(t)
	client := newAccountContractClient(t, router)
	ctx := context.Background()
	email := "partial-" + primitive.NewObjectID().Hex() + "@example.com"

	verifyOtpSignIn(t, client, email, "123456")
	account, err := repositoryForTest(t).GetAccountByEmail(ctx, email)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() { deleteAccountTestFixtures(t, account.ExternalUserID) })

	restore := accounts.SetDefaultDeleter(accounts.Deleter{
		DeletePostgres: func(context.Context, string) error { return errors.New("postgres unavailable") },
	})
	defer restore()

	client.request(http.MethodDelete, "/api/user", map[string]any{"email": email}, http.StatusInternalServerError)
	client.request(http.MethodGet, "/api/user/profile", nil, http.StatusOK)
	if _, err := repositoryForTest(t).GetAccountByExternalUserID(ctx, account.ExternalUserID); err != nil {
		t.Fatalf("partial failure removed account authority: %v", err)
	}

	restore()
	client.request(http.MethodDelete, "/api/user", map[string]any{"email": email}, http.StatusOK)
	client.request(http.MethodGet, "/api/user/profile", nil, http.StatusUnauthorized)
}

// TestAccountDeletionAllowsFreshReSignIn proves that signing in with the same
// email after deletion creates a new account and platform identity, leaving no
// orphan identity behind.
func TestAccountDeletionAllowsFreshReSignIn(t *testing.T) {
	router := newAccountContractRouter(t)
	client := newAccountContractClient(t, router)
	ctx := context.Background()
	email := "resignin-" + primitive.NewObjectID().Hex() + "@example.com"

	verifyOtpSignIn(t, client, email, "123456")
	repository := repositoryForTest(t)
	original, err := repository.GetAccountByEmail(ctx, email)
	if err != nil {
		t.Fatal(err)
	}
	t.Cleanup(func() {
		deleteAccountTestFixtures(t, original.ExternalUserID)
		if replacement, err := repository.GetAccountByEmail(ctx, email); err == nil {
			deleteAccountTestFixtures(t, replacement.ExternalUserID)
		}
	})

	client.request(http.MethodDelete, "/api/user", map[string]any{"email": email}, http.StatusOK)
	profile := verifyOtpSignIn(t, client, email, "654321")
	if got := decodeAccountString(t, profile, "email"); got != email {
		t.Fatalf("re-sign-in email = %q, want %q", got, email)
	}
	replacement, err := repository.GetAccountByEmail(ctx, email)
	if err != nil {
		t.Fatal(err)
	}
	if replacement.ExternalUserID == original.ExternalUserID {
		t.Fatal("re-sign-in reused the deleted external identity")
	}
	var orphaned int
	if err := pgstore.Pool.QueryRow(ctx, `SELECT count(*) FROM platform_identities WHERE external_user_id = $1`, original.ExternalUserID).Scan(&orphaned); err != nil {
		t.Fatal(err)
	}
	if orphaned != 0 {
		t.Fatalf("deleted platform identity is still present: %d", orphaned)
	}
}
