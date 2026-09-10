-- +goose Up
-- Historical daily user logs leave the MongoDB dailyuserlogs collection and
-- become PostgreSQL-owned per the retained-data contract. Each log receives a
-- fresh identity keyed by its account-local date, and each membership references
-- the authoritative account through its external user identifier in
-- platform_identities. The denormalized users array is never stored; profile
-- fields are rebuilt from accounts at read time.
CREATE TABLE daily_user_logs (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    -- Start of the account-local month/day/year; exactly one row per date.
    log_date DATE NOT NULL UNIQUE,
    created_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

-- One membership per account per log. first_seen_position preserves insertion
-- order so reporting lists accounts in the order they first signed in on that
-- local day, matching the legacy userIds array order.
CREATE TABLE daily_user_log_members (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    daily_user_log_id UUID NOT NULL REFERENCES daily_user_logs(id) ON DELETE CASCADE,
    account_user_id TEXT NOT NULL CHECK (account_user_id <> ''),
    first_seen_position INTEGER NOT NULL CHECK (first_seen_position >= 0),
    created_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
    UNIQUE (daily_user_log_id, account_user_id)
);

-- Active-user reporting filters logs by date and joins members to accounts.
CREATE INDEX daily_user_log_members_account_user_id_idx ON daily_user_log_members (account_user_id);

-- +goose Down
DROP TABLE daily_user_log_members;
DROP TABLE daily_user_logs;
