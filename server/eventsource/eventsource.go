// Package eventsource classifies public event identifiers before a storage
// implementation attempts legacy MongoDB ID resolution.
package eventsource

import (
	"regexp"
	"strings"
)

type Source uint8

const (
	Unknown Source = iota
	MongoDB
	PostgreSQL
)

const MongoDBIDPrefix = "m_"

var crockfordShortID = regexp.MustCompile(`^[0-9A-HJKMNPQRSTVWXYZ]{8}$`)

// Parse validates a public event identifier and returns its storage source and
// unwrapped storage identifier. MongoDB identifiers can be explicitly
// namespaced; bare eight-character Crockford identifiers belong to PostgreSQL.
func Parse(id string) (Source, string) {
	if strings.HasPrefix(id, MongoDBIDPrefix) {
		return MongoDB, strings.TrimPrefix(id, MongoDBIDPrefix)
	}
	if crockfordShortID.MatchString(id) {
		return PostgreSQL, id
	}
	if strings.HasPrefix(id, "p_") || id == "" {
		return Unknown, ""
	}

	return MongoDB, id
}

func MongoPublicID(id string) string { return MongoDBIDPrefix + id }
