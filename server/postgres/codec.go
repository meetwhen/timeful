package postgres

import (
	"bytes"
	"encoding/json"
	"fmt"
)

var emptyPayload = json.RawMessage(`{}`)

func encodePayload(payload json.RawMessage) ([]byte, error) {
	if len(payload) == 0 {
		return append([]byte(nil), emptyPayload...), nil
	}
	if !json.Valid(payload) {
		return nil, fmt.Errorf("payload is not valid JSON")
	}
	if !bytes.HasPrefix(bytes.TrimSpace(payload), []byte("{")) {
		return nil, fmt.Errorf("payload must be a JSON object")
	}
	var value map[string]json.RawMessage
	if err := json.Unmarshal(payload, &value); err != nil {
		return nil, fmt.Errorf("payload must be a JSON object: %w", err)
	}
	return append([]byte(nil), payload...), nil
}

func decodePayload(payload []byte) json.RawMessage {
	return append(json.RawMessage(nil), payload...)
}
