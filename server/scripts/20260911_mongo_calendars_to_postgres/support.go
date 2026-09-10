package main

import (
	"context"
	"crypto/aes"
	"crypto/cipher"
	"encoding/base64"
	"errors"
	"os"

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
	// Ledger and quarantine kinds. The calendar unit is one owner account that
	// covers every connection, credential, sub-calendar, and preference it owns.
	ledgerKindCalendarAccount = "calendar-account"

	// The retained-data contract defines missing-owner-account for a calendar
	// unit whose owner has no PostgreSQL account. An unreadable legacy
	// credential fails the run rather than being silently substituted.
	reasonMissingOwnerAccount = "missing-owner-account"

	// legacyEncryptionKeyEnvironment is the key the legacy AES-CFB Apple
	// password helper used. It is the same ENCRYPTION_KEY the GCM codec reads.
	legacyEncryptionKeyEnvironment = "ENCRYPTION_KEY"
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

// decryptLegacyCFB decrypts a value produced by the legacy server/utils AES-CFB
// helper so the backfill can re-encrypt the plaintext with the AES-256-GCM
// envelope. The legacy format is base64(iv || CFB-ciphertext). It returns an
// error instead of panicking or substituting an empty credential.
func decryptLegacyCFB(encoded string) (string, error) {
	block, err := aes.NewCipher([]byte(os.Getenv(legacyEncryptionKeyEnvironment)))
	if err != nil {
		return "", err
	}
	ciphertext, err := base64.StdEncoding.DecodeString(encoded)
	if err != nil {
		return "", err
	}
	if len(ciphertext) < aes.BlockSize {
		return "", errors.New("legacy credential ciphertext is too short")
	}
	iv := ciphertext[:aes.BlockSize]
	plaintext := make([]byte, len(ciphertext)-aes.BlockSize)
	stream := cipher.NewCFBDecrypter(block, iv)
	stream.XORKeyStream(plaintext, ciphertext[aes.BlockSize:])
	return string(plaintext), nil
}
