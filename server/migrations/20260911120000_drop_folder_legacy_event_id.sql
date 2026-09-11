-- +goose Up
-- Folder membership stores exactly one PostgreSQL event reference. Refuse the
-- change while any row still carries a value in the retired reference column,
-- then drop that column and require the PostgreSQL reference.
-- +goose StatementBegin
DO $$
BEGIN
    IF EXISTS (SELECT 1 FROM folder_events WHERE legacy_event_id IS NOT NULL) THEN
        RAISE EXCEPTION 'folder_events still holds values in legacy_event_id; refusing to drop the column';
    END IF;
END
$$;
-- +goose StatementEnd

ALTER TABLE folder_events DROP COLUMN legacy_event_id;
ALTER TABLE folder_events ALTER COLUMN event_id SET NOT NULL;

-- +goose Down
ALTER TABLE folder_events ALTER COLUMN event_id DROP NOT NULL;
ALTER TABLE folder_events ADD COLUMN legacy_event_id TEXT NULL;
ALTER TABLE folder_events ADD CONSTRAINT folder_events_event_reference CHECK (
    (event_id IS NOT NULL AND legacy_event_id IS NULL)
    OR
    (event_id IS NULL AND legacy_event_id IS NOT NULL)
);
CREATE UNIQUE INDEX folder_events_legacy_event_unique_idx
    ON folder_events (account_user_id, legacy_event_id)
    WHERE legacy_event_id IS NOT NULL;
