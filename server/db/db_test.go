package db_test

import (
	"strings"
	"testing"
	"time"

	"timeful/server/db"
)

func TestGetDailyUserLogByDate(t *testing.T) {
	db.GetDailyUserLogByDate(time.Now(), 7)
}

const validShortIdLetters = "0123456789ABCDEFGHJKMNPQRSTVWXYZ"

func TestGenerateShortEventId(t *testing.T) {
	db.Init()

	seen := make(map[string]bool)
	for i := 0; i < 200; i++ {
		id := db.GenerateShortEventId()
		if len(id) != 8 {
			t.Fatalf("GenerateShortEventId() = %q, want 8 characters", id)
		}
		for _, char := range id {
			if !strings.ContainsRune(validShortIdLetters, char) {
				t.Fatalf("GenerateShortEventId() = %q contains invalid character %q", id, char)
			}
		}
		if seen[id] {
			t.Fatalf("GenerateShortEventId() returned duplicate id %q", id)
		}
		seen[id] = true
	}
}
