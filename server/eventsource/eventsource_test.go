package eventsource

import "testing"

func TestParse(t *testing.T) {
	tests := []struct {
		name       string
		id         string
		wantSource Source
		wantID     string
	}{
		{name: "Mongo long ID", id: "m_64f5e4d3c2b1a09876543210", wantSource: MongoDB, wantID: "64f5e4d3c2b1a09876543210"},
		{name: "Mongo short ID", id: "m_ABCD1234", wantSource: MongoDB, wantID: "ABCD1234"},
		{name: "PostgreSQL short ID", id: "ABCD1234", wantSource: PostgreSQL, wantID: "ABCD1234"},
		{name: "legacy Mongo long ID", id: "64f5e4d3c2b1a09876543210", wantSource: MongoDB, wantID: "64f5e4d3c2b1a09876543210"},
		{name: "invalid legacy PostgreSQL ID", id: "p_ABCD1234", wantSource: Unknown},
		{name: "non-Crockford ID remains Mongo", id: "ABCIO123", wantSource: MongoDB, wantID: "ABCIO123"},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			gotSource, gotID := Parse(test.id)
			if gotSource != test.wantSource || gotID != test.wantID {
				t.Fatalf("Parse(%q) = (%v, %q), want (%v, %q)", test.id, gotSource, gotID, test.wantSource, test.wantID)
			}
		})
	}
}
