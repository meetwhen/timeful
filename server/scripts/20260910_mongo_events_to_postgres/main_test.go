package main

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"timeful/server/scripts/internal/legacybson"
)

func TestClassifyEventKinds(t *testing.T) {
	signup := true
	cases := []struct {
		name  string
		event legacybson.Event
		want  string
	}{
		{"signup", legacybson.Event{IsSignUpForm: &signup}, eventKindSignup},
		{"group", legacybson.Event{Type: legacybson.GROUP}, eventKindGroup},
		{"dow", legacybson.Event{Type: legacybson.DOW}, eventKindDayOfWeek},
		{"specific dates", legacybson.Event{Type: legacybson.SPECIFIC_DATES}, eventKindSpecificDates},
		{"empty type", legacybson.Event{}, eventKindSpecificDates},
	}
	for _, testCase := range cases {
		if got := classifyEvent(testCase.event); got != testCase.want {
			t.Fatalf("%s: classifyEvent = %q, want %q", testCase.name, got, testCase.want)
		}
	}
}

func TestBuildEventPayloadDropsIdentityAndTableOwnedFields(t *testing.T) {
	shortID := "ABCD1234"
	numResponses := 2
	blocks := []legacybson.SignUpBlock{{Id: primitive.NewObjectID(), Name: "Block"}}
	event := legacybson.Event{
		Id:              primitive.NewObjectID(),
		ShortId:         &shortID,
		OwnerId:         primitive.NewObjectID(),
		Name:            "Payload",
		NumResponses:    &numResponses,
		SignUpBlocks:    &blocks,
		SignUpResponses: map[string]*legacybson.SignUpResponse{"guest": {Name: "Guest"}},
		ResponsesMap:    map[string]*legacybson.Response{"guest": {Name: "Guest"}},
	}
	data, err := buildEventPayload(event)
	if err != nil {
		t.Fatal(err)
	}
	var value map[string]any
	if err := json.Unmarshal(data, &value); err != nil {
		t.Fatal(err)
	}
	if value["_id"] != primitive.NilObjectID.Hex() || value["ownerId"] != primitive.NilObjectID.Hex() {
		t.Fatalf("payload retained identity: %s", data)
	}
	for _, key := range []string{"shortId", "numResponses", "signUpBlocks", "responses", "signUpResponses"} {
		if value[key] != nil {
			t.Fatalf("payload retained %q: %v", key, value[key])
		}
	}
	if value["name"] != "Payload" {
		t.Fatalf("payload lost event name: %s", data)
	}
}

func TestPrepareResponseClassifiesGuestAndQuarantinesCredential(t *testing.T) {
	responseID := primitive.NewObjectID()
	stored := legacybson.EventResponse{
		Id: responseID,
		Response: &legacybson.Response{
			Name:               "Ada",
			GuestId:            "guest-ada",
			GuestEditToken:     "secret",
			GuestOwnershipMode: "token",
			GuestEditPolicy:    "protected",
		},
	}
	m := &migrator{}
	var unit unitEvent
	m.prepareResponse(nil, primitive.NilObjectID, stored, &unit)
	if len(unit.responses) != 1 {
		t.Fatalf("responses = %d, want 1", len(unit.responses))
	}
	if unit.responses[0].respondentKind != respondentKindGuest {
		t.Fatalf("respondent kind = %q", unit.responses[0].respondentKind)
	}
	if unit.responses[0].canonicalGuestName == nil || *unit.responses[0].canonicalGuestName != "Ada" {
		t.Fatalf("canonical name = %v", unit.responses[0].canonicalGuestName)
	}
	if len(unit.quarantine) != 1 || unit.quarantine[0].reason != reasonLegacyGuestCredential {
		t.Fatalf("quarantine = %#v", unit.quarantine)
	}
	if !strings.Contains(string(unit.responses[0].payload), `"availability"`) {
		t.Fatalf("payload missing availability: %s", unit.responses[0].payload)
	}
}

func TestPrepareResponseQuarantinesInvalidGuestName(t *testing.T) {
	stored := legacybson.EventResponse{
		Id:       primitive.NewObjectID(),
		Response: &legacybson.Response{Name: "0123456789abcdef01234567"},
	}
	m := &migrator{}
	var unit unitEvent
	m.prepareResponse(nil, primitive.NilObjectID, stored, &unit)
	if len(unit.responses) != 0 {
		t.Fatalf("invalid guest name must not migrate: %#v", unit.responses)
	}
	if len(unit.quarantine) != 1 || unit.quarantine[0].reason != reasonInvalidGuestName {
		t.Fatalf("quarantine = %#v", unit.quarantine)
	}
}

func TestPrepareResponseQuarantinesMissingIdentity(t *testing.T) {
	stored := legacybson.EventResponse{Id: primitive.NewObjectID(), Response: &legacybson.Response{}}
	m := &migrator{}
	var unit unitEvent
	m.prepareResponse(nil, primitive.NilObjectID, stored, &unit)
	if len(unit.responses) != 0 {
		t.Fatalf("missing identity must not migrate: %#v", unit.responses)
	}
	if len(unit.quarantine) != 1 || unit.quarantine[0].reason != reasonMissingResponseIdentity {
		t.Fatalf("quarantine = %#v", unit.quarantine)
	}
}

func TestPrepareSignupDataRewritesBlockMembership(t *testing.T) {
	blockID := primitive.NewObjectID()
	missingBlockID := primitive.NewObjectID()
	signup := true
	event := legacybson.Event{
		IsSignUpForm: &signup,
		SignUpBlocks: &[]legacybson.SignUpBlock{{Id: blockID, Name: "Morning"}},
		SignUpResponses: map[string]*legacybson.SignUpResponse{
			"Dana": {Name: "Dana", SignUpBlockIds: []primitive.ObjectID{blockID, missingBlockID}},
		},
	}
	m := &migrator{}
	var unit unitEvent
	unit.legacyID = event.Id.Hex()
	if err := m.prepareSignupData(event, &unit); err != nil {
		t.Fatal(err)
	}
	if len(unit.blocks) != 1 || unit.blocks[0].position != 1 {
		t.Fatalf("blocks = %#v", unit.blocks)
	}
	if len(unit.signupResponses) != 1 {
		t.Fatalf("signup responses = %#v", unit.signupResponses)
	}
	if len(unit.signupResponses[0].blockLegacyIDs) != 1 || unit.signupResponses[0].blockLegacyIDs[0] != blockID.Hex() {
		t.Fatalf("block membership = %#v", unit.signupResponses[0].blockLegacyIDs)
	}
	foundOrphan := false
	for _, record := range unit.quarantine {
		if record.reason == reasonOrphanMembership {
			foundOrphan = true
		}
	}
	if !foundOrphan {
		t.Fatalf("missing orphan membership quarantine: %#v", unit.quarantine)
	}
}

func TestReconciliationReportFormatting(t *testing.T) {
	report := reconciliationReport{
		Events:             1,
		Quarantined:        2,
		QuarantineByReason: map[string]int{reasonInvalidGuestName: 2},
	}
	output := report.String()
	if !strings.Contains(output, "mismatches: none") {
		t.Fatalf("clean report = %q", output)
	}
	report.Mismatches = []string{"events drifted"}
	if !strings.Contains(report.String(), "events drifted") {
		t.Fatalf("mismatch not reported: %q", report.String())
	}
}

func TestSignupInstantAndHelpers(t *testing.T) {
	instant := primitive.NewDateTimeFromTime(time.UnixMilli(1_700_000_000_000).UTC())
	value := signupInstant(&instant)
	if value == nil || !value.Equal(instant.Time().UTC()) {
		t.Fatalf("signup instant = %v", value)
	}
	if signupInstant(nil) != nil {
		t.Fatal("nil instant must stay nil")
	}
	if optionalString("") != nil {
		t.Fatal("empty string must map to nil")
	}
	if got := optionalString("x"); got == nil || *got != "x" {
		t.Fatalf("optional string = %v", got)
	}
	if boolValue(nil) {
		t.Fatal("nil bool must be false")
	}
}
