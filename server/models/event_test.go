package models

import (
	"encoding/json"
	"testing"
	"time"
)

func boolPtrEvent(value bool) *bool { return &value }
func intPtrEvent(value int) *int    { return &value }

func TestEventIDFieldsKeep24HexWireFormat(t *testing.T) {
	event := Event{Id: "507f1f77bcf86cd799439011", OwnerId: "507f1f77bcf86cd799439012"}

	payload, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal event: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("decode event: %v", err)
	}
	if decoded["_id"] != "507f1f77bcf86cd799439011" {
		t.Fatalf("_id = %v", decoded["_id"])
	}
	if decoded["ownerId"] != "507f1f77bcf86cd799439012" {
		t.Fatalf("ownerId = %v", decoded["ownerId"])
	}
}

func TestEventZeroIDsSurfaceGuestSentinel(t *testing.T) {
	payload, err := json.Marshal(Event{})
	if err != nil {
		t.Fatalf("marshal event: %v", err)
	}
	var decoded map[string]any
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("decode event: %v", err)
	}
	if decoded["_id"] != "000000000000000000000000" {
		t.Fatalf("_id = %v, want the zero ObjectID sentinel", decoded["_id"])
	}
	if decoded["ownerId"] != "000000000000000000000000" {
		t.Fatalf("ownerId = %v, want the zero ObjectID sentinel", decoded["ownerId"])
	}

	signUp := SignUpResponse{}
	signUpPayload, err := json.Marshal(signUp)
	if err != nil {
		t.Fatalf("marshal signup response: %v", err)
	}
	var signUpDecoded map[string]any
	if err := json.Unmarshal(signUpPayload, &signUpDecoded); err != nil {
		t.Fatalf("decode signup response: %v", err)
	}
	if signUpDecoded["userId"] != "000000000000000000000000" {
		t.Fatalf("signup userId = %v, want the zero ObjectID sentinel", signUpDecoded["userId"])
	}
}

func TestEventMarshalJSONSuppressesLegacyTimedScheduleColumns(t *testing.T) {
	event := Event{
		Id:               "507f1f77bcf86cd799439011",
		OwnerId:          "507f1f77bcf86cd799439012",
		Duration:         float32PtrEvent(1.5),
		Dates:            []DateTime{NewDateTimeFromTime(time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC))},
		TimeIncrement:    intPtrEvent(15),
		HasSpecificTimes: boolPtrEvent(true),
		Times:            []DateTime{NewDateTimeFromTime(time.Date(2026, 1, 5, 9, 0, 0, 0, time.UTC))},
		StartOnMonday:    boolPtrEvent(true),
	}

	payload, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal event: %v", err)
	}
	var decoded map[string]json.RawMessage
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("decode event: %v", err)
	}
	for _, legacyField := range []string{"duration", "dates", "timeIncrement", "hasSpecificTimes", "times", "startOnMonday"} {
		if _, exists := decoded[legacyField]; exists {
			t.Fatalf("expected %q to be suppressed for timed events", legacyField)
		}
	}
}

func TestEventMarshalJSONKeepsLegacyAttendeesKeyAsNull(t *testing.T) {
	for name, event := range map[string]Event{
		"timed":    {},
		"daysOnly": {DaysOnly: boolPtrEvent(true)},
	} {
		t.Run(name, func(t *testing.T) {
			payload, err := json.Marshal(event)
			if err != nil {
				t.Fatalf("marshal event: %v", err)
			}
			var decoded map[string]json.RawMessage
			if err := json.Unmarshal(payload, &decoded); err != nil {
				t.Fatalf("decode event: %v", err)
			}
			attendees, exists := decoded["attendees"]
			if !exists {
				t.Fatalf("expected attendees key in %s", payload)
			}
			if string(attendees) != "null" {
				t.Fatalf("attendees = %s, want null", attendees)
			}
		})
	}
}

func TestEventMarshalJSONKeepsLegacyColumnsForDaysOnlyEvents(t *testing.T) {
	event := Event{
		Id:            "507f1f77bcf86cd799439011",
		OwnerId:       "507f1f77bcf86cd799439012",
		DaysOnly:      boolPtrEvent(true),
		Duration:      float32PtrEvent(1.5),
		Dates:         []DateTime{NewDateTimeFromTime(time.Date(2026, 1, 5, 0, 0, 0, 0, time.UTC))},
		TimeIncrement: intPtrEvent(15),
	}

	payload, err := json.Marshal(event)
	if err != nil {
		t.Fatalf("marshal event: %v", err)
	}
	var decoded map[string]json.RawMessage
	if err := json.Unmarshal(payload, &decoded); err != nil {
		t.Fatalf("decode event: %v", err)
	}
	if _, exists := decoded["duration"]; !exists {
		t.Fatal("expected the day-only event to keep duration")
	}
	if string(decoded["dates"]) != `["2026-01-05T00:00:00Z"]` {
		t.Fatalf("dates = %s", decoded["dates"])
	}
	if _, exists := decoded["timeIncrement"]; !exists {
		t.Fatal("expected the day-only event to keep timeIncrement")
	}
}

func float32PtrEvent(value float32) *float32 { return &value }
