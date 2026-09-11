package main

import (
	"context"
	"errors"

	"github.com/jackc/pgx/v5"
	"github.com/jackc/pgx/v5/pgconn"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// querier is satisfied by *pgxpool.Pool and pgx.Tx so migration helpers run
// either statement-by-statement or inside one unit transaction.
type querier interface {
	Exec(context.Context, string, ...any) (pgconn.CommandTag, error)
	Query(context.Context, string, ...any) (pgx.Rows, error)
	QueryRow(context.Context, string, ...any) pgx.Row
}

const (
	// ledgerKindDailyUserLog is the migration kind recorded for one historical
	// daily user log unit. The unit is one retained dailyuserlogs document with
	// its membership list.
	ledgerKindDailyUserLog = "daily-user-log"

	// The retained-data contract defines missing-owner-account for a retained
	// unit whose owner has no PostgreSQL account. A daily-log member that cannot
	// be resolved quarantines the whole log.
	reasonMissingOwnerAccount = "missing-owner-account"
)

// isNoRows reports whether a query matched no row.
func isNoRows(err error) bool {
	return errors.Is(err, pgx.ErrNoRows)
}

// pageFilter advances a MongoDB _id cursor without treating a legitimate zero
// ObjectID as "no page read yet".
func pageFilter(lastID primitive.ObjectID, hasCursor bool) bson.M {
	if !hasCursor {
		return bson.M{}
	}
	return bson.M{"_id": bson.M{"$gt": lastID}}
}

// orderedDistinct collapses duplicate account identifiers while preserving the
// first occurrence, matching the legacy userIds array semantics where an
// account appears at most once and keeps its first-seen position.
func orderedDistinct(accountUserIDs []string) []string {
	seen := make(map[string]bool, len(accountUserIDs))
	distinct := make([]string, 0, len(accountUserIDs))
	for _, accountUserID := range accountUserIDs {
		if accountUserID == "" || seen[accountUserID] {
			continue
		}
		seen[accountUserID] = true
		distinct = append(distinct, accountUserID)
	}
	return distinct
}

// mergeOrdered appends incoming identifiers after existing ones, dropping any
// identifier already present. Two retained documents can share a date, and
// their membership merges into one target log in source order.
func mergeOrdered(existing, incoming []string) []string {
	return orderedDistinct(append(append([]string{}, existing...), incoming...))
}
