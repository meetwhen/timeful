-- +goose Up
-- PostgreSQL 18 supplies uuidv7(). Public identifiers remain storage-opaque.
CREATE TABLE postgres_events (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    public_id TEXT NOT NULL UNIQUE,
    short_id TEXT NOT NULL UNIQUE,
    owner_external_id TEXT NULL,
    name TEXT NOT NULL,
    type TEXT NOT NULL,
    is_archived BOOLEAN NOT NULL DEFAULT FALSE,
    is_deleted BOOLEAN NOT NULL DEFAULT FALSE,
    num_responses INTEGER NOT NULL DEFAULT 0 CHECK (num_responses >= 0),
    schedule_version INTEGER NOT NULL DEFAULT 1,
    creator_posthog_id TEXT NULL,
    payload JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
    CONSTRAINT postgres_events_public_id_format CHECK (
        public_id ~ '^p_[0-9A-HJKMNPQRSTVWXYZ]{26}$'
    ),
    CONSTRAINT postgres_events_short_id_format CHECK (
        short_id ~ '^p_[0-9A-HJKMNPQRSTVWXYZ]{8}$'
    ),
    CONSTRAINT postgres_events_type CHECK (type IN ('specific_dates', 'dow')),
    CONSTRAINT postgres_events_payload_object CHECK (jsonb_typeof(payload) = 'object')
);

CREATE INDEX postgres_events_active_creator_posthog_id_idx
    ON postgres_events (creator_posthog_id, created_at DESC)
    WHERE NOT is_deleted;

CREATE TABLE postgres_event_responses (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    event_id UUID NOT NULL REFERENCES postgres_events(id) ON DELETE CASCADE,
    respondent_kind TEXT NOT NULL,
    account_user_id TEXT NULL,
    guest_id TEXT NULL,
    canonical_guest_name TEXT NULL,
    guest_edit_policy TEXT NULL,
    guest_ownership_mode TEXT NULL,
    -- Retained temporarily so tokenless open mutations can return credentials.
    guest_edit_token TEXT NULL,
    payload JSONB NOT NULL DEFAULT '{}'::jsonb,
    created_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
    CONSTRAINT postgres_event_responses_kind CHECK (
        respondent_kind IN ('account', 'guest')
    ),
    CONSTRAINT postgres_event_responses_identity CHECK (
        (respondent_kind = 'account' AND account_user_id IS NOT NULL)
        OR
        (respondent_kind = 'guest' AND account_user_id IS NULL AND canonical_guest_name IS NOT NULL)
    ),
    CONSTRAINT postgres_event_responses_token_ownership CHECK (
        guest_ownership_mode IS DISTINCT FROM 'token'
        OR (guest_id IS NOT NULL AND guest_edit_token IS NOT NULL)
    ),
    CONSTRAINT postgres_event_responses_payload_object CHECK (jsonb_typeof(payload) = 'object')
);

CREATE INDEX postgres_event_responses_event_id_idx
    ON postgres_event_responses (event_id);

CREATE UNIQUE INDEX postgres_event_responses_account_unique_idx
    ON postgres_event_responses (event_id, account_user_id)
    WHERE respondent_kind = 'account';

CREATE UNIQUE INDEX postgres_event_responses_guest_id_unique_idx
    ON postgres_event_responses (event_id, guest_id)
    WHERE respondent_kind = 'guest' AND guest_id IS NOT NULL;

CREATE UNIQUE INDEX postgres_event_responses_guest_name_unique_idx
    ON postgres_event_responses (event_id, canonical_guest_name)
    WHERE respondent_kind = 'guest';

-- +goose Down
DROP TABLE postgres_event_responses;
DROP TABLE postgres_events;
