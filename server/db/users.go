package db

import (
	"context"
	"errors"
	"strings"

	"github.com/jackc/pgx/v5"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"timeful/server/logger"
	"timeful/server/models"
	pgstore "timeful/server/postgres"
)

// GetUserById returns the authoritative PostgreSQL account profile. A
// PostgreSQL lookup that fails for any reason other than a genuine no-row
// result yields nil instead of an inferred account, so a transient database
// failure can never be mistaken for account authority.
func GetUserById(userId string) *models.User {
	account, err := accountByExternalUserID(userId)
	if err != nil {
		logger.StdErr.Printf("account lookup failed for %s: %v", userId, err)
		return nil
	}
	return AccountUser(account)
}

// GetUserByEmail resolves the authoritative account by case-insensitive email
// and returns its profile. A PostgreSQL failure other than a genuine no-row
// result yields nil instead of an inferred account.
func GetUserByEmail(email string) *models.User {
	emailQuery := strings.TrimSpace(email)
	if emailQuery == "" {
		return nil
	}
	account, err := accountByEmail(emailQuery)
	if err != nil {
		logger.StdErr.Printf("account lookup failed for %s: %v", emailQuery, err)
		return nil
	}
	return AccountUser(account)
}

// AccountUser builds the internal user shape from the authoritative PostgreSQL
// account. It never reads or writes the retained store.
func AccountUser(account *pgstore.Account) *models.User {
	if account == nil {
		return nil
	}
	user := &models.User{}
	if objectID, err := primitive.ObjectIDFromHex(account.ExternalUserID); err == nil {
		user.Id = objectID
	}
	user.Email = account.Email
	user.FirstName = account.FirstName
	user.LastName = account.LastName
	user.Picture = account.Picture
	user.HasCustomName = account.HasCustomName
	user.TimezoneOffset = account.TimezoneOffset
	user.NumEventsCreated = account.NumEventsCreated
	return user
}

// accountByExternalUserID resolves the authoritative PostgreSQL account. A
// deliberately uninitialized pool and a missing account row both return no
// account; every other PostgreSQL failure is returned so callers never mistake
// the error for absence.
func accountByExternalUserID(userId string) (*pgstore.Account, error) {
	if userId == "" {
		return nil, nil
	}
	repository, err := pgstore.DefaultRepository()
	if err != nil {
		if errors.Is(err, pgstore.ErrPoolUninitialized) {
			return nil, nil
		}
		return nil, err
	}
	account, err := repository.GetAccountByExternalUserID(context.Background(), userId)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return account, nil
}

// accountByEmail applies the same outcome classification as
// accountByExternalUserID for a case-insensitive email lookup.
func accountByEmail(email string) (*pgstore.Account, error) {
	repository, err := pgstore.DefaultRepository()
	if err != nil {
		if errors.Is(err, pgstore.ErrPoolUninitialized) {
			return nil, nil
		}
		return nil, err
	}
	account, err := repository.GetAccountByEmail(context.Background(), email)
	if errors.Is(err, pgx.ErrNoRows) {
		return nil, nil
	}
	if err != nil {
		return nil, err
	}
	return account, nil
}
