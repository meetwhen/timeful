package postgres

import (
	"context"
	"testing"
)

func TestPingReturnsErrorWhenPoolIsUninitialized(t *testing.T) {
	previousPool := Pool
	Pool = nil
	t.Cleanup(func() { Pool = previousPool })

	if err := Ping(context.Background()); err == nil {
		t.Fatal("Ping() error = nil, want an uninitialized pool error")
	}
}

func TestEnvironmentInt(t *testing.T) {
	t.Setenv(maxConnectionsEnvironment, "12")
	if got := environmentInt(maxConnectionsEnvironment, 10); got != 12 {
		t.Fatalf("environmentInt() = %d, want 12", got)
	}
}
