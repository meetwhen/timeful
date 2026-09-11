package models

import (
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"strings"
	"testing"
	"time"
)

func marshalJSONString(t *testing.T, value any) string {
	t.Helper()
	payload, err := json.Marshal(value)
	if err != nil {
		t.Fatalf("marshal %T: %v", value, err)
	}
	return string(payload)
}

func TestZeroIDMarshalsAs24HexObjectIDSentinel(t *testing.T) {
	if got := marshalJSONString(t, ZeroID()); got != `"000000000000000000000000"` {
		t.Fatalf("zero ID = %s, want the zero ObjectID sentinel", got)
	}

	var unset ID
	if got := marshalJSONString(t, unset); got != `"000000000000000000000000"` {
		t.Fatalf("unset ID = %s, want the zero ObjectID sentinel", got)
	}
	if !unset.IsZero() {
		t.Fatal("expected the unset ID to be zero")
	}
	if !ZeroID().IsZero() {
		t.Fatal("expected the zero sentinel to be zero")
	}
}

func TestIDUnmarshalAcceptsDriverForms(t *testing.T) {
	var value ID
	if err := json.Unmarshal([]byte(`"507f1f77bcf86cd799439011"`), &value); err != nil {
		t.Fatalf("unmarshal 24-hex ID: %v", err)
	}
	if value != ID("507f1f77bcf86cd799439011") {
		t.Fatalf("decoded ID = %q", value)
	}

	if err := json.Unmarshal([]byte(`"507F1F77BCF86CD799439011"`), &value); err != nil {
		t.Fatalf("unmarshal uppercase ID: %v", err)
	}
	if value != ID("507f1f77bcf86cd799439011") {
		t.Fatalf("uppercase ID did not canonicalize: %q", value)
	}

	if err := json.Unmarshal([]byte(`""`), &value); err != nil {
		t.Fatalf("unmarshal empty ID: %v", err)
	}
	if !value.IsZero() {
		t.Fatalf("empty string did not decode as zero: %q", value)
	}

	if err := json.Unmarshal([]byte(`null`), &value); err != nil {
		t.Fatalf("unmarshal null ID: %v", err)
	}

	for _, invalid := range []string{`"abc"`, `"507f1f77bcf86cd79943901g"`, `"507f1f77bcf86cd7994390110"`, `"not-an-id"`} {
		var decoded ID
		if err := json.Unmarshal([]byte(invalid), &decoded); err == nil {
			t.Fatalf("expected %s to be rejected", invalid)
		}
	}
}

func TestIDTextEncodingKeepsSentinelForMapKeys(t *testing.T) {
	encoded := marshalJSONString(t, map[ID]string{ZeroID(): "zero", "507f1f77bcf86cd799439011": "event"})
	want := `{"000000000000000000000000":"zero","507f1f77bcf86cd799439011":"event"}`
	if encoded != want {
		t.Fatalf("map = %s, want %s", encoded, want)
	}

	var decoded map[ID]string
	if err := json.Unmarshal([]byte(encoded), &decoded); err != nil {
		t.Fatalf("decode map: %v", err)
	}
	if decoded[ZeroID()] != "zero" || decoded[ID("507f1f77bcf86cd799439011")] != "event" {
		t.Fatalf("decoded map = %v", decoded)
	}

	var emptyKey map[ID]string
	if err := json.Unmarshal([]byte(`{"":"x"}`), &emptyKey); err != nil {
		t.Fatalf("decode empty map key: %v", err)
	}
	if emptyKey[ZeroID()] != "x" {
		t.Fatalf("empty map key did not decode as the zero sentinel: %v", emptyKey)
	}
}

func TestIDUnmarshalAcceptsExtendedJSON(t *testing.T) {
	var value ID
	if err := json.Unmarshal([]byte(`{"$oid":"507f1f77bcf86cd799439011"}`), &value); err != nil {
		t.Fatalf("unmarshal extended JSON ID: %v", err)
	}
	if value != ID("507f1f77bcf86cd799439011") {
		t.Fatalf("decoded extended ID = %q", value)
	}

	if err := json.Unmarshal([]byte(`{"$oid":""}`), &value); err != nil {
		t.Fatalf("unmarshal empty extended ID: %v", err)
	}
	if !value.IsZero() {
		t.Fatalf("empty extended ID did not decode as zero: %q", value)
	}

	for _, invalid := range []string{`{}`, `{"$oid":42}`, `{"other":"507f1f77bcf86cd799439011"}`} {
		var decoded ID
		if err := json.Unmarshal([]byte(invalid), &decoded); err == nil {
			t.Fatalf("expected %s to be rejected", invalid)
		}
	}
}

func TestParseIDValidates24Hex(t *testing.T) {
	parsed, ok := ParseID("507F1F77BCF86CD799439011")
	if !ok {
		t.Fatal("expected valid 24-hex identifier")
	}
	if parsed != ID("507f1f77bcf86cd799439011") {
		t.Fatalf("parsed ID = %q", parsed)
	}

	for _, invalid := range []string{"", "abc", "507f1f77bcf86cd79943901g", strings.Repeat("0", 23), strings.Repeat("0", 25)} {
		if _, ok := ParseID(invalid); ok {
			t.Fatalf("expected %q to be rejected", invalid)
		}
	}
}

func TestNewIDKeeps24HexTimePrefixedObjectIDShape(t *testing.T) {
	id := NewID()
	if id.IsZero() {
		t.Fatal("expected a non-zero generated identifier")
	}
	if len(id) != 24 {
		t.Fatalf("generated ID %q has length %d", id, len(id))
	}
	if _, ok := ParseID(id.Hex()); !ok {
		t.Fatalf("generated ID %q is not 24-hex", id)
	}

	raw, err := hex.DecodeString(string(id))
	if err != nil {
		t.Fatalf("decode generated ID: %v", err)
	}
	if len(raw) != 12 {
		t.Fatalf("generated ID has %d bytes, want 12", len(raw))
	}
	prefixSeconds := int64(binary.BigEndian.Uint32(raw[:4]))
	if delta := time.Since(time.Unix(prefixSeconds, 0)); delta < -time.Minute || delta > time.Minute {
		t.Fatalf("generated ID time prefix is %v off, want within a minute", delta)
	}

	if NewID() == NewID() {
		t.Fatal("expected generated identifiers to be unique")
	}
}
