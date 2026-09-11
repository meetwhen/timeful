package respondents

import (
	"timeful/server/models"
)

// ResolveStoredUserID accepts either an explicit embedded user id or a legacy
// outer storage key that encoded the signed-in user identity.
func ResolveStoredUserID(explicit models.ID, storedKey string) (models.ID, bool) {
	if !explicit.IsZero() {
		return explicit, true
	}

	objectID, ok := models.ParseID(storedKey)
	if !ok {
		return "", false
	}

	return objectID, true
}
