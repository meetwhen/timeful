package models

import (
	"crypto/rand"
	"encoding/binary"
	"encoding/hex"
	"encoding/json"
	"errors"
	"strings"
	"time"
)

const zeroIDValue = "000000000000000000000000"

// ZeroID returns the 24-hex sentinel the API emits for a zero identifier. It is
// surfaced to clients for an absent account identity (for example the guest
// userId) and for unowned events, so the wire format must keep it. It is
// exposed as a function so callers cannot overwrite the sentinel.
func ZeroID() ID { return zeroIDValue }

// ID is a canonical 24-hex identifier. Its JSON representation is a
// 24-character lowercase hexadecimal string whose empty value serializes as the
// zero sentinel.
type ID string

// Hex returns the 24-hex form, mapping the empty value to the zero sentinel.
func (id ID) Hex() string {
	if id == "" {
		return zeroIDValue
	}
	return string(id)
}

// String returns the 24-hex form used on the wire.
func (id ID) String() string { return id.Hex() }

// IsZero reports whether the identifier is the empty or zero sentinel value.
func (id ID) IsZero() bool { return id == "" || id == zeroIDValue }

// MarshalJSON emits the 24-hex string, including the zero sentinel.
func (id ID) MarshalJSON() ([]byte, error) { return json.Marshal(id.Hex()) }

// MarshalText emits the 24-hex string so ID keeps working as a JSON map key.
func (id ID) MarshalText() ([]byte, error) { return []byte(id.Hex()), nil }

// UnmarshalJSON accepts a 24-hex string, an empty string (the zero sentinel),
// null, or the extended JSON {"$oid":"..."} form. A twelve-byte raw identifier
// form is not accepted: a string-backed identifier cannot represent it without
// corruption and no producer emits it.
func (id *ID) UnmarshalJSON(data []byte) error {
	if string(data) == "null" {
		return nil
	}
	var raw string
	if err := json.Unmarshal(data, &raw); err != nil {
		var extended struct {
			OID *string `json:"$oid"`
		}
		if err := json.Unmarshal(data, &extended); err != nil {
			return err
		}
		if extended.OID == nil {
			return errors.New("not an extended JSON identifier")
		}
		raw = *extended.OID
	}
	return id.setHex(raw)
}

// UnmarshalText accepts the 24-hex form when ID is a JSON map key. It rejects
// the empty string; only UnmarshalJSON maps empty values to the zero sentinel.
func (id *ID) UnmarshalText(data []byte) error {
	value, ok := ParseID(string(data))
	if !ok {
		return errors.New("invalid 24-hex identifier")
	}
	*id = value
	return nil
}

func (id *ID) setHex(raw string) error {
	if raw == "" {
		*id = zeroIDValue
		return nil
	}
	if len(raw) != len(zeroIDValue) {
		return errors.New("invalid 24-hex identifier length")
	}
	if _, err := hex.DecodeString(raw); err != nil {
		return err
	}
	*id = ID(strings.ToLower(raw))
	return nil
}

// ParseID validates a 24-hex string and returns its canonical lowercase form.
func ParseID(value string) (ID, bool) {
	if len(value) != len(zeroIDValue) {
		return "", false
	}
	if _, err := hex.DecodeString(value); err != nil {
		return "", false
	}
	return ID(strings.ToLower(value)), true
}

// NewID returns a fresh canonical 24-hex identifier. The first four bytes carry
// the current Unix time so identifiers keep a roughly time-ordered shape.
func NewID() ID {
	var value [12]byte
	binary.BigEndian.PutUint32(value[:4], uint32(time.Now().Unix()))
	_, _ = rand.Read(value[4:])
	return ID(hex.EncodeToString(value[:]))
}
