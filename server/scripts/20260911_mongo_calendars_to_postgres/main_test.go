package main

import (
	"encoding/json"
	"strings"
	"testing"
	"time"

	"go.mongodb.org/mongo-driver/bson/primitive"

	"timeful/server/models"
	"timeful/server/utils"
)

const testEncryptionKey = "0123456789abcdef0123456789abcdef"

func stringPointer(value string) *string { return &value }
func boolPointer(value bool) *bool       { return &value }

// TestConvertCalendarConnectionMapsRetainedFields proves the retained
// connection fields, including absent-versus-explicit-false enabled, survive the
// conversion to the PostgreSQL shape.
func TestConvertCalendarConnectionMapsRetainedFields(t *testing.T) {
	subCalendars := map[string]models.SubCalendar{
		"primary": {Name: "Primary", Enabled: boolPointer(false)},
		"hidden":  {Name: "Hidden"},
	}
	expiresAt := primitive.NewDateTimeFromTime(time.UnixMilli(1700000000000).UTC())
	account := models.CalendarAccount{
		CalendarType: models.GoogleCalendarType,
		Email:        "person@example.com",
		Picture:      "https://example.com/p.png",
		OAuth2CalendarAuth: &models.OAuth2CalendarAuth{
			AccessToken:           "access",
			RefreshToken:          "refresh",
			Scope:                 "scope",
			AccessTokenExpireDate: expiresAt,
		},
		SubCalendars: &subCalendars,
	}

	converted, err := convertCalendarConnection("person@example.com_google", account)
	if err != nil {
		t.Fatal(err)
	}
	if converted.CalendarKey != "person@example.com_google" || converted.CalendarType != "google" {
		t.Fatalf("connection identity = %q %q", converted.CalendarKey, converted.CalendarType)
	}
	if converted.Email != "person@example.com" || converted.Picture != "https://example.com/p.png" {
		t.Fatalf("connection fields = %q %q", converted.Email, converted.Picture)
	}
	if converted.Enabled != nil {
		t.Fatalf("omitted enabled must stay nil, got %v", *converted.Enabled)
	}
	if converted.OAuth2 == nil || converted.OAuth2.AccessToken != "access" || converted.OAuth2.RefreshToken != "refresh" || converted.OAuth2.Scope != "scope" {
		t.Fatalf("oauth credentials = %#v", converted.OAuth2)
	}
	if converted.OAuth2.AccessTokenExpiresAt == nil || !converted.OAuth2.AccessTokenExpiresAt.Equal(expiresAt.Time().UTC()) {
		t.Fatalf("oauth expiry = %v, want %v", converted.OAuth2.AccessTokenExpiresAt, expiresAt.Time().UTC())
	}
	if len(converted.SubCalendars) != 2 {
		t.Fatalf("sub-calendars = %#v", converted.SubCalendars)
	}
	if converted.SubCalendars["primary"].Enabled == nil || *converted.SubCalendars["primary"].Enabled {
		t.Fatalf("explicit false sub-calendar enabled was lost: %#v", converted.SubCalendars["primary"])
	}
	if converted.SubCalendars["hidden"].Enabled != nil {
		t.Fatalf("omitted sub-calendar enabled must stay nil")
	}
}

// TestConvertCalendarConnectionDecryptsLegacyApplePassword proves the backfill
// reads the legacy AES-CFB value and hands the repository plaintext so it can
// re-encrypt with the GCM envelope.
func TestConvertCalendarConnectionDecryptsLegacyApplePassword(t *testing.T) {
	t.Setenv("ENCRYPTION_KEY", testEncryptionKey)
	encrypted, err := utils.Encrypt("app-password")
	if err != nil {
		t.Fatal(err)
	}
	account := models.CalendarAccount{
		CalendarType:      models.AppleCalendarType,
		Email:             "apple@example.com",
		AppleCalendarAuth: &models.AppleCalendarAuth{Email: "apple@example.com", Password: encrypted},
	}

	converted, err := convertCalendarConnection("apple@example.com_apple", account)
	if err != nil {
		t.Fatal(err)
	}
	if converted.Apple == nil || converted.Apple.Password != "app-password" {
		t.Fatalf("apple password = %#v, want plaintext", converted.Apple)
	}
}

// TestConvertCalendarPreferencesAbsentVersusPresent covers the retained
// preference fields, including absent versus present primary account key.
func TestConvertCalendarPreferencesAbsentVersusPresent(t *testing.T) {
	preferences, err := convertCalendarPreferences(models.User{})
	if err != nil {
		t.Fatal(err)
	}
	if preferences != nil {
		t.Fatalf("absent preferences must not create a row: %#v", preferences)
	}

	options := &models.CalendarOptions{
		BufferTime:   models.BufferTimeOptions{Enabled: true, Time: 15},
		WorkingHours: models.WorkingHoursOptions{Enabled: true, StartTime: 9, EndTime: 17},
	}
	preferences, err = convertCalendarPreferences(models.User{
		PrimaryAccountKey: stringPointer("person@example.com_google"),
		TokenOrigin:       models.WEB,
		CalendarOptions:   options,
	})
	if err != nil {
		t.Fatal(err)
	}
	if preferences == nil || preferences.PrimaryAccountKey == nil || *preferences.PrimaryAccountKey != "person@example.com_google" {
		t.Fatalf("primary account key = %#v", preferences)
	}
	if preferences.TokenOrigin == nil || *preferences.TokenOrigin != "web" {
		t.Fatalf("token origin = %#v", preferences.TokenOrigin)
	}
	var decoded models.CalendarOptions
	if err := json.Unmarshal(preferences.CalendarOptions, &decoded); err != nil {
		t.Fatal(err)
	}
	if decoded.BufferTime.Time != 15 || decoded.WorkingHours.EndTime != 17 {
		t.Fatalf("calendar options = %#v", decoded)
	}
}

func TestReconciliationReportFormatting(t *testing.T) {
	report := reconciliationReport{
		Units:              2,
		Accounts:           4,
		SubCalendars:       3,
		CredentialRows:     4,
		Preferences:        2,
		Quarantined:        1,
		QuarantineByReason: map[string]int{reasonMissingOwnerAccount: 1},
		Mismatches:         []string{"example mismatch"},
	}
	output := report.String()
	for _, want := range []string{
		"units=2 calendar_accounts=4 sub_calendars=3 credential_rows=4 preferences=2",
		"quarantined=1",
		"missing-owner-account=1",
		"example mismatch",
	} {
		if !strings.Contains(output, want) {
			t.Fatalf("report missing %q:\n%s", want, output)
		}
	}
}
