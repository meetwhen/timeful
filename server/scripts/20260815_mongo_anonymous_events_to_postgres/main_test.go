package main

import (
	"encoding/json"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"
	"timeful/server/models"
)

func TestBuildEventUsesPostgresPayloadShape(t *testing.T) {
	shortID := "6df78"
	responses := 5
	daysOnly := false
	timezone := "Asia/Karachi"
	event := models.Event{
		Id:              primitive.NewObjectIDFromTimestamp(time.Date(2026, 8, 1, 2, 3, 4, 0, time.UTC)),
		ShortId:         &shortID,
		Name:            "Demo",
		Type:            models.SPECIFIC_DATES,
		DaysOnly:        &daysOnly,
		EventTimezone:   &timezone,
		SlotGeneration:  &models.SlotGeneration{StartTimeLocal: "09:00:00", EndTimeLocal: "10:00:00", TimeIncrementMinutes: 15},
		TimedRecurrence: &models.TimedRecurrence{Kind: "specific_dates"},
		ActiveSlots:     []primitive.DateTime{primitive.NewDateTimeFromTime(time.Date(2026, 8, 2, 9, 0, 0, 0, time.UTC))},
		NumResponses:    &responses,
	}

	migrated, err := buildEvent(event)
	if err != nil {
		t.Fatal(err)
	}
	if migrated.numResponses != 5 || migrated.scheduleVersion != 1 {
		t.Fatalf("event metadata = responses=%d scheduleVersion=%d", migrated.numResponses, migrated.scheduleVersion)
	}
	if !migrated.createdAt.Equal(event.Id.Timestamp().UTC()) || !migrated.updatedAt.Equal(migrated.createdAt) {
		t.Fatalf("timestamps = %s, %s", migrated.createdAt, migrated.updatedAt)
	}

	var payload map[string]json.RawMessage
	if err := json.Unmarshal(migrated.payload, &payload); err != nil {
		t.Fatal(err)
	}
	for _, field := range []string{"shortId", "numResponses"} {
		if string(payload[field]) != "null" {
			t.Fatalf("payload %s = %s, want null", field, payload[field])
		}
	}
	var payloadID, payloadOwnerID string
	if err := json.Unmarshal(payload["_id"], &payloadID); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(payload["ownerId"], &payloadOwnerID); err != nil {
		t.Fatal(err)
	}
	if payloadID != primitive.NilObjectID.Hex() || payloadOwnerID != primitive.NilObjectID.Hex() {
		t.Fatalf("payload identities = _id=%q ownerId=%q", payloadID, payloadOwnerID)
	}
	if _, exists := payload["activeSlots"]; !exists {
		t.Fatal("payload is missing activeSlots")
	}
}

func TestBuildResponseMigratesGuestIdentityAndSlots(t *testing.T) {
	available := primitive.NewDateTimeFromTime(time.Date(2026, 8, 6, 0, 30, 0, 0, time.UTC))
	ifNeeded := primitive.NewDateTimeFromTime(time.Date(2026, 8, 6, 0, 45, 0, 0, time.UTC))
	stored := models.EventResponse{
		Id:     primitive.NewObjectIDFromTimestamp(time.Date(2026, 8, 3, 0, 0, 0, 0, time.UTC)),
		UserId: "guest-opaque-id",
		Response: &models.Response{
			Name:               " Avery Chen ",
			GuestId:            "guest-opaque-id",
			GuestEditToken:     "secret-token",
			GuestEditPolicy:    "open",
			GuestOwnershipMode: "token",
			Availability:       []primitive.DateTime{available, available},
			IfNeeded:           []primitive.DateTime{available, ifNeeded, ifNeeded},
		},
	}

	migrated, err := buildResponse(stored)
	if err != nil {
		t.Fatal(err)
	}
	if migrated.respondentKind != "guest" || migrated.canonicalGuestName == nil || *migrated.canonicalGuestName != "Avery Chen" {
		t.Fatalf("guest identity = kind=%q canonical=%v", migrated.respondentKind, migrated.canonicalGuestName)
	}
	if migrated.guestEditToken == nil || *migrated.guestEditToken != "secret-token" {
		t.Fatalf("guest token = %v", migrated.guestEditToken)
	}

	var payload map[string]json.RawMessage
	if err := json.Unmarshal(migrated.payload, &payload); err != nil {
		t.Fatal(err)
	}
	if _, exists := payload["guestEditToken"]; exists {
		t.Fatal("guest token must be stored only in its PostgreSQL column")
	}
	var payloadAvailability, payloadIfNeeded []primitive.DateTime
	if err := json.Unmarshal(payload["availability"], &payloadAvailability); err != nil {
		t.Fatal(err)
	}
	if err := json.Unmarshal(payload["ifNeeded"], &payloadIfNeeded); err != nil {
		t.Fatal(err)
	}
	if len(payloadAvailability) != 1 || len(payloadIfNeeded) != 1 || payloadIfNeeded[0] != ifNeeded {
		t.Fatalf("normalized slots = availability=%v ifNeeded=%v", payloadAvailability, payloadIfNeeded)
	}
}
