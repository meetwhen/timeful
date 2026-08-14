// Package eventsource classifies public event identifiers before a storage
// implementation attempts legacy MongoDB ID resolution.
package eventsource

import "strings"

type Source uint8

const (
	MongoDB Source = iota
	PostgreSQL
)

const PostgreSQLIDPrefix = "p_"

// Classify returns PostgreSQL for every reserved PostgreSQL namespace value,
// including malformed values. This prevents a malformed p_ ID from reaching
// MongoDB's short-ID or ObjectID resolution paths.
func Classify(id string) Source {
	if strings.HasPrefix(id, PostgreSQLIDPrefix) {
		return PostgreSQL
	}

	return MongoDB
}
