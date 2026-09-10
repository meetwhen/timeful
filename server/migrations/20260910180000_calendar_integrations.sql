-- +goose Up
-- Calendar integrations leave the retained MongoDB users document and become
-- PostgreSQL-owned per the retained-data contract. Each connection keeps its
-- legacy email_CALENDARTYPE map key as calendar_key so the runtime key semantics
-- survive the move, and is owned by exactly one platform identity resolved
-- through platform_identities.external_user_id.
CREATE TABLE calendar_accounts (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    platform_identity_id UUID NOT NULL REFERENCES platform_identities(id) ON DELETE CASCADE,
    calendar_key TEXT NOT NULL CHECK (calendar_key <> ''),
    calendar_type TEXT NOT NULL CHECK (calendar_type IN ('google', 'outlook', 'apple', 'ics')),
    email TEXT NOT NULL DEFAULT '',
    picture TEXT NOT NULL DEFAULT '',
    -- NULL preserves a legacy document that omitted enabled, distinct from an
    -- explicit false.
    enabled BOOLEAN NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
    UNIQUE (platform_identity_id, calendar_key)
);

CREATE INDEX calendar_accounts_platform_identity_id_idx
    ON calendar_accounts (platform_identity_id);

-- A sub-calendar receives a fresh identity scoped to its connection and keeps
-- the provider calendar id as its runtime key. enabled is NULL when the legacy
-- document omitted it, distinct from an explicit false.
CREATE TABLE calendar_sub_calendars (
    id UUID PRIMARY KEY DEFAULT uuidv7(),
    calendar_account_id UUID NOT NULL REFERENCES calendar_accounts(id) ON DELETE CASCADE,
    sub_calendar_id TEXT NOT NULL CHECK (sub_calendar_id <> ''),
    name TEXT NOT NULL DEFAULT '',
    enabled BOOLEAN NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
    UNIQUE (calendar_account_id, sub_calendar_id)
);

-- Exactly one credential row per connection. Provider secrets are encrypted at
-- rest with the versioned AES-256-GCM envelope; expiry and scope are not secrets
-- and stay plaintext so token refresh never needs to decrypt. A NULL ciphertext
-- means the credential was absent, never an empty secret.
CREATE TABLE calendar_account_credentials (
    calendar_account_id UUID PRIMARY KEY REFERENCES calendar_accounts(id) ON DELETE CASCADE,
    oauth_access_token_ciphertext TEXT NULL,
    oauth_refresh_token_ciphertext TEXT NULL,
    oauth_access_token_expires_at TIMESTAMPTZ NULL,
    oauth_scope TEXT NULL,
    apple_password_ciphertext TEXT NULL,
    ics_feed_url_ciphertext TEXT NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

-- Exactly one preference row per platform identity. primary_account_key NULL
-- preserves the legacy first-Google-account fallback; token_origin NULL means
-- undefined and is never exposed through the API; calendar_options NULL
-- preserves an absent preference distinct from a present empty object.
CREATE TABLE calendar_preferences (
    platform_identity_id UUID PRIMARY KEY REFERENCES platform_identities(id) ON DELETE CASCADE,
    primary_account_key TEXT NULL CHECK (primary_account_key IS NULL OR primary_account_key <> ''),
    token_origin TEXT NULL CHECK (token_origin IS NULL OR token_origin IN ('ios', 'android', 'web')),
    calendar_options JSONB NULL,
    created_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp(),
    updated_at TIMESTAMPTZ NOT NULL DEFAULT clock_timestamp()
);

-- +goose Down
DROP TABLE calendar_preferences;
DROP TABLE calendar_account_credentials;
DROP TABLE calendar_sub_calendars;
DROP TABLE calendar_accounts;
