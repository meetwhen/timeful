package db

import (
	"context"
	"errors"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
)

// DeleteAccountData removes every retained MongoDB record that belongs to the
// deleted visitor. It is idempotent, so a retry after a partial failure
// converges without removing any other account's data:
//
//   - the visitor's own event responses,
//   - the visitor's folders and folder memberships,
//   - friend requests sent to or by the visitor,
//   - ownership of events the visitor organized, which stay reachable and keep
//     their other guests' responses,
//   - the retained calendar-integration document.
//
// Historical daily-logs are PostgreSQL-owned and their membership cleanup runs
// in the PostgreSQL deletion transaction.
//
// PostgreSQL account authority, events, responses, visitor identities, and
// platform identities are removed by the PostgreSQL deletion transaction.
func DeleteAccountData(ctx context.Context, externalUserID string) error {
	objectID, err := primitive.ObjectIDFromHex(externalUserID)
	if err != nil {
		return errors.New("account identifier must be a hexadecimal object identifier")
	}
	if _, err := EventResponsesCollection.DeleteMany(ctx, bson.M{"userId": externalUserID}); err != nil {
		return err
	}
	if _, err := FoldersCollection.DeleteMany(ctx, bson.M{"userId": objectID}); err != nil {
		return err
	}
	if _, err := FolderEventsCollection.DeleteMany(ctx, bson.M{"userId": objectID}); err != nil {
		return err
	}
	if _, err := FriendRequestsCollection.DeleteMany(ctx, bson.M{"$or": bson.A{
		bson.M{"from": objectID},
		bson.M{"to": objectID},
	}}); err != nil {
		return err
	}
	if _, err := EventsCollection.UpdateMany(ctx, bson.M{"ownerId": objectID}, bson.M{
		"$unset": bson.M{"ownerId": ""},
	}); err != nil {
		return err
	}
	_, err = UsersCollection.DeleteOne(ctx, bson.M{"_id": objectID})
	return err
}
