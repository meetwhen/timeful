package eventsource

import "testing"

func TestClassify(t *testing.T) {
	tests := []struct {
		name string
		id   string
		want Source
	}{
		{name: "Mongo long ID", id: "64f5e4d3c2b1a09876543210", want: MongoDB},
		{name: "Mongo short ID", id: "ABCD1234", want: MongoDB},
		{name: "PostgreSQL long ID", id: "p_01J3NYJ4ABCD1234EFGH5678JK", want: PostgreSQL},
		{name: "PostgreSQL short ID", id: "p_ABCD1234", want: PostgreSQL},
		{name: "malformed PostgreSQL namespace", id: "p_", want: PostgreSQL},
	}

	for _, test := range tests {
		t.Run(test.name, func(t *testing.T) {
			if got := Classify(test.id); got != test.want {
				t.Fatalf("Classify(%q) = %v, want %v", test.id, got, test.want)
			}
		})
	}
}
