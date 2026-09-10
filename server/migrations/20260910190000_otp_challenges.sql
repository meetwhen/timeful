-- +goose Up
-- OTP challenges leave the MongoDB otpCodes collection and become
-- PostgreSQL-owned per the retained-data contract. Each challenge receives a
-- fresh identity and is keyed by email, so at most one challenge is active per
-- email and a new send replaces the existing one. The one-time code is stored
-- only as a salted one-way hash, never as plaintext, and verified by hash.
CREATE TABLE otp_challenges (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    email TEXT NOT NULL UNIQUE CHECK (email <> ''),
    code_hash TEXT NOT NULL CHECK (code_hash <> ''),
    expires_at TIMESTAMPTZ NOT NULL,
    -- Preserves the legacy five-attempt lockout, counted before each compare.
    attempts INTEGER NOT NULL DEFAULT 0 CHECK (attempts >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

-- Expired challenges are swept explicitly instead of by a MongoDB TTL index; the
-- index keeps the sweep and the verification predicate bounded.
CREATE INDEX otp_challenges_expires_at_idx ON otp_challenges (expires_at);

-- +goose Down
DROP TABLE otp_challenges;
