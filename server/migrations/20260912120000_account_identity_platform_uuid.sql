-- +goose Up
-- Consolidate account identity onto platform_identities.id.
--
-- Every account reference is rewritten from the retired 24-character
-- hexadecimal external identifier to the native uuid platform identity, each
-- rewrite is verified to have zero unmapped rows, and only then are the legacy
-- TEXT columns and platform_identities.external_user_id dropped. Goose wraps
-- this migration in one transaction, so a failure rolls back the whole unit and
-- leaves the legacy mapping intact. Every legacy-column read is guarded by a
-- column-existence check, so re-running the migration is a no-op.
--
-- account_deletion_tombstones has no backfill from a live identity: the deletion
-- transaction removes the platform identity before it records the tombstone, so
-- every existing tombstone names an identity that is already gone. Those rows
-- are removed with the legacy column because the 24-character value is
-- unreachable after the cutover and no platform identity exists to protect. The
-- table keeps no foreign key so future tombstones survive identity deletion.

-- postgres_events.owner_platform_identity_id already exists and becomes the
-- canonical owner reference; only nulls from the legacy column are filled.
-- +goose StatementBegin
DO $$
DECLARE
    target oid;
BEGIN
    target := to_regclass('postgres_events');
    IF target IS NOT NULL AND EXISTS (
        SELECT 1 FROM pg_attribute
        WHERE attrelid = target AND attname = 'owner_external_id' AND attnum > 0 AND NOT attisdropped
    ) THEN
        UPDATE postgres_events e
        SET owner_platform_identity_id = p.id
        FROM platform_identities p
        WHERE e.owner_platform_identity_id IS NULL
          AND e.owner_external_id IS NOT NULL
          AND p.external_user_id = e.owner_external_id;

        IF EXISTS (
            SELECT 1
            FROM postgres_events e
            WHERE e.owner_external_id IS NOT NULL
              AND e.owner_platform_identity_id IS NULL
        ) THEN
            RAISE EXCEPTION 'postgres_events rows reference an unmapped owner external id';
        END IF;
    END IF;
END $$;
-- +goose StatementEnd

ALTER TABLE postgres_events DROP COLUMN IF EXISTS owner_external_id;

-- Event responses: replace the TEXT account reference with the platform
-- identity uuid.
ALTER TABLE postgres_event_responses
    ADD COLUMN IF NOT EXISTS platform_identity_id UUID REFERENCES platform_identities(id);

-- +goose StatementBegin
DO $$
DECLARE
    target oid;
BEGIN
    target := to_regclass('postgres_event_responses');
    IF target IS NOT NULL AND EXISTS (
        SELECT 1 FROM pg_attribute
        WHERE attrelid = target AND attname = 'account_user_id' AND attnum > 0 AND NOT attisdropped
    ) THEN
        UPDATE postgres_event_responses r
        SET platform_identity_id = p.id
        FROM platform_identities p
        WHERE r.platform_identity_id IS NULL
          AND r.account_user_id IS NOT NULL
          AND p.external_user_id = r.account_user_id;

        IF EXISTS (
            SELECT 1
            FROM postgres_event_responses r
            WHERE r.account_user_id IS NOT NULL
              AND r.platform_identity_id IS NULL
        ) THEN
            RAISE EXCEPTION 'postgres_event_responses rows reference an unmapped account user id';
        END IF;
    END IF;
END $$;
-- +goose StatementEnd

ALTER TABLE postgres_event_responses DROP COLUMN IF EXISTS account_user_id;
CREATE INDEX IF NOT EXISTS postgres_response_platform_identity_idx
    ON postgres_event_responses (platform_identity_id);

-- Signup responses: the identity check and the one-account-per-event unique
-- index move to the platform identity uuid.
ALTER TABLE event_signup_responses
    ADD COLUMN IF NOT EXISTS platform_identity_id UUID REFERENCES platform_identities(id);

-- +goose StatementBegin
DO $$
DECLARE
    target oid;
BEGIN
    target := to_regclass('event_signup_responses');
    IF target IS NOT NULL AND EXISTS (
        SELECT 1 FROM pg_attribute
        WHERE attrelid = target AND attname = 'account_user_id' AND attnum > 0 AND NOT attisdropped
    ) THEN
        UPDATE event_signup_responses r
        SET platform_identity_id = p.id
        FROM platform_identities p
        WHERE r.platform_identity_id IS NULL
          AND r.account_user_id IS NOT NULL
          AND p.external_user_id = r.account_user_id;

        IF EXISTS (
            SELECT 1
            FROM event_signup_responses r
            WHERE r.account_user_id IS NOT NULL
              AND r.platform_identity_id IS NULL
        ) THEN
            RAISE EXCEPTION 'event_signup_responses rows reference an unmapped account user id';
        END IF;
    END IF;
END $$;
-- +goose StatementEnd

ALTER TABLE event_signup_responses DROP CONSTRAINT IF EXISTS event_signup_responses_identity;
DROP INDEX IF EXISTS event_signup_responses_account_unique_idx;
ALTER TABLE event_signup_responses DROP COLUMN IF EXISTS account_user_id;
ALTER TABLE event_signup_responses ADD CONSTRAINT event_signup_responses_identity CHECK (
    (respondent_kind = 'account' AND platform_identity_id IS NOT NULL)
    OR
    (respondent_kind = 'guest' AND platform_identity_id IS NULL AND canonical_guest_name IS NOT NULL AND canonical_guest_name <> '')
);
CREATE UNIQUE INDEX IF NOT EXISTS event_signup_responses_account_unique_idx
    ON event_signup_responses (event_id, platform_identity_id)
    WHERE respondent_kind = 'account';

-- Attendees: the release-on-deletion relation stores the platform identity uuid.
ALTER TABLE event_attendees
    ADD COLUMN IF NOT EXISTS platform_identity_id UUID REFERENCES platform_identities(id);

-- +goose StatementBegin
DO $$
DECLARE
    target oid;
BEGIN
    target := to_regclass('event_attendees');
    IF target IS NOT NULL AND EXISTS (
        SELECT 1 FROM pg_attribute
        WHERE attrelid = target AND attname = 'account_user_id' AND attnum > 0 AND NOT attisdropped
    ) THEN
        UPDATE event_attendees a
        SET platform_identity_id = p.id
        FROM platform_identities p
        WHERE a.platform_identity_id IS NULL
          AND a.account_user_id IS NOT NULL
          AND p.external_user_id = a.account_user_id;

        IF EXISTS (
            SELECT 1
            FROM event_attendees a
            WHERE a.account_user_id IS NOT NULL
              AND a.platform_identity_id IS NULL
        ) THEN
            RAISE EXCEPTION 'event_attendees rows reference an unmapped account user id';
        END IF;
    END IF;
END $$;
-- +goose StatementEnd

DROP INDEX IF EXISTS event_attendees_account_user_id_idx;
ALTER TABLE event_attendees DROP COLUMN IF EXISTS account_user_id;
CREATE INDEX IF NOT EXISTS event_attendees_platform_identity_id_idx
    ON event_attendees (platform_identity_id);

-- Folders are account-scoped, so their owner reference is required.
ALTER TABLE folders
    ADD COLUMN IF NOT EXISTS platform_identity_id UUID REFERENCES platform_identities(id);

-- +goose StatementBegin
DO $$
DECLARE
    target oid;
BEGIN
    target := to_regclass('folders');
    IF target IS NOT NULL AND EXISTS (
        SELECT 1 FROM pg_attribute
        WHERE attrelid = target AND attname = 'account_user_id' AND attnum > 0 AND NOT attisdropped
    ) THEN
        UPDATE folders f
        SET platform_identity_id = p.id
        FROM platform_identities p
        WHERE f.platform_identity_id IS NULL
          AND f.account_user_id IS NOT NULL
          AND p.external_user_id = f.account_user_id;

        IF EXISTS (
            SELECT 1
            FROM folders f
            WHERE f.account_user_id IS NOT NULL
              AND f.platform_identity_id IS NULL
        ) THEN
            RAISE EXCEPTION 'folders rows reference an unmapped account user id';
        END IF;
    END IF;
END $$;
-- +goose StatementEnd

DROP INDEX IF EXISTS folders_account_user_id_idx;
ALTER TABLE folders DROP COLUMN IF EXISTS account_user_id;
ALTER TABLE folders ALTER COLUMN platform_identity_id SET NOT NULL;
CREATE INDEX IF NOT EXISTS folders_platform_identity_id_idx
    ON folders (platform_identity_id);

-- Folder memberships are account-scoped, so their owner reference is required.
ALTER TABLE folder_events
    ADD COLUMN IF NOT EXISTS platform_identity_id UUID REFERENCES platform_identities(id);

-- +goose StatementBegin
DO $$
DECLARE
    target oid;
BEGIN
    target := to_regclass('folder_events');
    IF target IS NOT NULL AND EXISTS (
        SELECT 1 FROM pg_attribute
        WHERE attrelid = target AND attname = 'account_user_id' AND attnum > 0 AND NOT attisdropped
    ) THEN
        UPDATE folder_events fe
        SET platform_identity_id = p.id
        FROM platform_identities p
        WHERE fe.platform_identity_id IS NULL
          AND fe.account_user_id IS NOT NULL
          AND p.external_user_id = fe.account_user_id;

        IF EXISTS (
            SELECT 1
            FROM folder_events fe
            WHERE fe.account_user_id IS NOT NULL
              AND fe.platform_identity_id IS NULL
        ) THEN
            RAISE EXCEPTION 'folder_events rows reference an unmapped account user id';
        END IF;
    END IF;
END $$;
-- +goose StatementEnd

DROP INDEX IF EXISTS folder_events_event_unique_idx;
ALTER TABLE folder_events DROP COLUMN IF EXISTS account_user_id;
ALTER TABLE folder_events ALTER COLUMN platform_identity_id SET NOT NULL;
CREATE UNIQUE INDEX IF NOT EXISTS folder_events_event_unique_idx
    ON folder_events (platform_identity_id, event_id)
    WHERE event_id IS NOT NULL;

-- Access transfers: the source account reference becomes a platform identity
-- uuid while the credential-based path keeps source_credential_id.
ALTER TABLE access_transfers
    ADD COLUMN IF NOT EXISTS platform_identity_id UUID REFERENCES platform_identities(id);

-- +goose StatementBegin
DO $$
DECLARE
    target oid;
BEGIN
    target := to_regclass('access_transfers');
    IF target IS NOT NULL AND EXISTS (
        SELECT 1 FROM pg_attribute
        WHERE attrelid = target AND attname = 'external_user_id' AND attnum > 0 AND NOT attisdropped
    ) THEN
        UPDATE access_transfers t
        SET platform_identity_id = p.id
        FROM platform_identities p
        WHERE t.platform_identity_id IS NULL
          AND t.external_user_id IS NOT NULL
          AND p.external_user_id = t.external_user_id;

        IF EXISTS (
            SELECT 1
            FROM access_transfers t
            WHERE t.external_user_id IS NOT NULL
              AND t.platform_identity_id IS NULL
        ) THEN
            RAISE EXCEPTION 'access_transfers rows reference an unmapped external user id';
        END IF;
    END IF;
END $$;
-- +goose StatementEnd

ALTER TABLE access_transfers DROP CONSTRAINT IF EXISTS access_transfers_check;
ALTER TABLE access_transfers DROP COLUMN IF EXISTS external_user_id;
ALTER TABLE access_transfers DROP CONSTRAINT IF EXISTS access_transfers_source_xor_platform_identity;
ALTER TABLE access_transfers ADD CONSTRAINT access_transfers_source_xor_platform_identity
    CHECK ((source_credential_id IS NULL) <> (platform_identity_id IS NULL));

-- Daily log memberships: the account reference becomes a platform identity uuid.
ALTER TABLE daily_user_log_members
    ADD COLUMN IF NOT EXISTS platform_identity_id UUID REFERENCES platform_identities(id);

-- +goose StatementBegin
DO $$
DECLARE
    target oid;
BEGIN
    target := to_regclass('daily_user_log_members');
    IF target IS NOT NULL AND EXISTS (
        SELECT 1 FROM pg_attribute
        WHERE attrelid = target AND attname = 'account_user_id' AND attnum > 0 AND NOT attisdropped
    ) THEN
        UPDATE daily_user_log_members m
        SET platform_identity_id = p.id
        FROM platform_identities p
        WHERE m.platform_identity_id IS NULL
          AND m.account_user_id IS NOT NULL
          AND p.external_user_id = m.account_user_id;

        IF EXISTS (
            SELECT 1
            FROM daily_user_log_members m
            WHERE m.account_user_id IS NOT NULL
              AND m.platform_identity_id IS NULL
        ) THEN
            RAISE EXCEPTION 'daily_user_log_members rows reference an unmapped account user id';
        END IF;
    END IF;
END $$;
-- +goose StatementEnd

DROP INDEX IF EXISTS daily_user_log_members_account_user_id_idx;
ALTER TABLE daily_user_log_members
    DROP CONSTRAINT IF EXISTS daily_user_log_members_daily_user_log_id_account_user_id_key;
ALTER TABLE daily_user_log_members DROP COLUMN IF EXISTS account_user_id;
ALTER TABLE daily_user_log_members ALTER COLUMN platform_identity_id SET NOT NULL;
ALTER TABLE daily_user_log_members
    DROP CONSTRAINT IF EXISTS daily_user_log_members_log_platform_identity_key;
ALTER TABLE daily_user_log_members
    ADD CONSTRAINT daily_user_log_members_log_platform_identity_key UNIQUE (daily_user_log_id, platform_identity_id);
CREATE INDEX IF NOT EXISTS daily_user_log_members_platform_identity_id_idx
    ON daily_user_log_members (platform_identity_id);

-- Deletion tombstones carry the deleted platform identity uuid with no foreign
-- key. Legacy rows name identities that were already removed, so they are
-- cleared with the legacy column.
ALTER TABLE account_deletion_tombstones
    ADD COLUMN IF NOT EXISTS platform_identity_id UUID;

-- +goose StatementBegin
DO $$
DECLARE
    target oid;
BEGIN
    target := to_regclass('account_deletion_tombstones');
    IF target IS NOT NULL AND EXISTS (
        SELECT 1 FROM pg_attribute
        WHERE attrelid = target AND attname = 'external_user_id' AND attnum > 0 AND NOT attisdropped
    ) THEN
        UPDATE account_deletion_tombstones t
        SET platform_identity_id = p.id
        FROM platform_identities p
        WHERE t.platform_identity_id IS NULL
          AND t.external_user_id IS NOT NULL
          AND p.external_user_id = t.external_user_id;

        DELETE FROM account_deletion_tombstones WHERE platform_identity_id IS NULL;
    END IF;
END $$;
-- +goose StatementEnd

ALTER TABLE account_deletion_tombstones
    DROP CONSTRAINT IF EXISTS account_deletion_tombstones_pkey;
ALTER TABLE account_deletion_tombstones DROP COLUMN IF EXISTS external_user_id;
ALTER TABLE account_deletion_tombstones ALTER COLUMN platform_identity_id SET NOT NULL;
ALTER TABLE account_deletion_tombstones ADD PRIMARY KEY (platform_identity_id);

-- The legacy account identifier is gone after every dependent reference is
-- consolidated.
ALTER TABLE platform_identities DROP COLUMN IF EXISTS external_user_id;

-- +goose Down
-- Reversing the consolidation would destroy the platform identity mapping and
-- recreate the account-deletion tombstone rows that were already unreachable at
-- cutover. Refuse instead of pretending to roll back.
-- +goose StatementBegin
DO $$ BEGIN
    RAISE EXCEPTION 'The account identity consolidation cannot be reversed';
END $$;
-- +goose StatementEnd
