package main

import (
	"context"
	"fmt"
	"log"
	"os"

	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"go.mongodb.org/mongo-driver/mongo/options"
	"timeful/server/db"
)

const shortIdIndexName = "shortId_1"

// shortIdRow decodes only the fields this migration touches, keeping the
// script compile-safe regardless of the runtime model's evolution.
type shortIdRow struct {
	Id      primitive.ObjectID `bson:"_id"`
	ShortId *string            `bson:"shortId"`
}

// duplicateShortIds returns the shortId values currently assigned to more
// than one event, bucketed by the events that carry them.
func duplicateShortIds(ctx context.Context) (map[string][]shortIdRow, error) {
	pipeline := bson.A{
		bson.M{"$match": bson.M{"shortId": bson.M{"$type": "string"}}},
		bson.M{"$group": bson.M{"_id": "$shortId", "count": bson.M{"$sum": 1}}},
		bson.M{"$match": bson.M{"count": bson.M{"$gt": 1}}},
	}
	cursor, err := db.EventsCollection.Aggregate(ctx, pipeline)
	if err != nil {
		return nil, err
	}
	duplicateValues := make([]string, 0)
	for cursor.Next(ctx) {
		var doc struct {
			Id    string `bson:"_id"`
			Count int    `bson:"count"`
		}
		if err := cursor.Decode(&doc); err != nil {
			return nil, err
		}
		duplicateValues = append(duplicateValues, doc.Id)
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	groups := make(map[string][]shortIdRow)
	for _, value := range duplicateValues {
		rows, err := db.EventsCollection.Find(ctx, bson.M{"shortId": value}, options.Find().SetSort(bson.D{{Key: "_id", Value: 1}}))
		if err != nil {
			return nil, err
		}
		var events []shortIdRow
		if err := rows.All(ctx, &events); err != nil {
			return nil, err
		}
		groups[value] = events
	}
	return groups, nil
}

func indexExists(ctx context.Context) (bool, error) {
	indexes, err := db.EventsCollection.Indexes().List(ctx)
	if err != nil {
		return false, err
	}
	for indexes.Next(ctx) {
		var doc struct {
			Name string `bson:"name"`
		}
		if err := indexes.Decode(&doc); err != nil {
			return false, err
		}
		if doc.Name == shortIdIndexName {
			return true, nil
		}
	}
	return false, indexes.Err()
}

func preflight(ctx context.Context) (int, error) {
	groups, err := duplicateShortIds(ctx)
	if err != nil {
		return 0, err
	}
	if len(groups) > 0 {
		fmt.Printf("duplicate shortIds found (%d groups):\n", len(groups))
		for value, events := range groups {
			ids := make([]string, 0, len(events))
			for _, event := range events {
				ids = append(ids, event.Id.Hex())
			}
			fmt.Printf("  %q -> %v\n", value, ids)
		}
	}

	exists, err := indexExists(ctx)
	if err != nil {
		return 0, err
	}
	fmt.Printf("unique shortId index present: %v\n", exists)
	return len(groups), nil
}

func dedupe(ctx context.Context, groups map[string][]shortIdRow) int {
	regenerated := 0
	for _, events := range groups {
		// Keep the oldest event's shortId, regenerate the rest.
		for _, event := range events[1:] {
			for {
				shortId := db.GenerateShortEventId()
				result, err := db.EventsCollection.UpdateOne(ctx, bson.M{"_id": event.Id}, bson.M{"$set": bson.M{"shortId": shortId}})
				if err != nil {
					log.Fatal(err)
				}
				if result.ModifiedCount == 1 {
					fmt.Printf("event %s: shortId regenerated -> %s\n", event.Id.Hex(), shortId)
					regenerated++
					break
				}
			}
		}
	}
	return regenerated
}

func main() {
	disconnect := db.Init()
	defer disconnect()

	ctx := context.Background()

	groups, err := duplicateShortIds(ctx)
	if err != nil {
		log.Fatal(err)
	}
	duplicates, err := preflight(ctx)
	if err != nil {
		log.Fatal(err)
	}

	if len(os.Args) != 2 || os.Args[1] != "--apply" {
		fmt.Println("shortId unique index migration preflight passed; rerun with --apply to dedupe and create the index")
		return
	}

	regenerated := 0
	if duplicates > 0 {
		regenerated = dedupe(ctx, groups)
	} else {
		fmt.Println("no duplicate shortIds to resolve")
	}

	exists, err := indexExists(ctx)
	if err != nil {
		log.Fatal(err)
	}
	if !exists {
		_, err = db.EventsCollection.Indexes().CreateOne(ctx, mongo.IndexModel{
			Keys:    bson.D{{Key: "shortId", Value: 1}},
			Options: options.Index().SetName(shortIdIndexName).SetUnique(true).SetPartialFilterExpression(bson.M{"shortId": bson.M{"$type": "string"}}),
		})
		if err != nil {
			log.Fatal(err)
		}
		fmt.Println("created unique index", shortIdIndexName)
	} else {
		fmt.Println("unique index", shortIdIndexName, "already present")
	}

	fmt.Printf("shortId unique index migration complete: regenerated=%d\n", regenerated)
}
