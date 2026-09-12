-- +goose Up
-- Stored JSONB payloads can still embed the retired 24-character account
-- identifier and the legacy zero sentinel, which the strict canonical UUID
-- decoder rejects on read. The account identity now comes from the consolidated
-- columns, which every served path re-projects, so the duplicated payload keys
-- are removed. This is a separate migration version because
-- 20260912120000_account_identity_platform_uuid.sql had already been applied in
-- some environments before this rewrite existed, and goose keys on version, so
-- the rewrite could never reach them from there. Removing the keys is
-- idempotent: a rewritten payload no longer matches.
UPDATE postgres_events
SET payload = payload - '_id' - 'ownerId' - 'signUpResponses' - 'responses'
WHERE payload ?| array['_id', 'ownerId', 'signUpResponses', 'responses'];

UPDATE postgres_event_responses
SET payload = payload - 'userId' - 'user'
WHERE payload ?| array['userId', 'user'];

-- +goose Down
-- The rewrite only removes derived payload keys; the consolidated columns that
-- every served path re-projects already exist, so there is nothing to restore.
SELECT 1;
