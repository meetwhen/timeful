-- +goose Up
ALTER TABLE postgres_events DROP CONSTRAINT postgres_events_public_id_format;
ALTER TABLE postgres_events DROP CONSTRAINT postgres_events_short_id_format;

-- Existing PostgreSQL links use p_ plus the canonical short identifier.
UPDATE postgres_events SET short_id = substring(short_id FROM 3);

ALTER TABLE postgres_events DROP COLUMN public_id;
ALTER TABLE postgres_events ADD CONSTRAINT postgres_events_short_id_format CHECK (
    short_id ~ '^[0-9A-HJKMNPQRSTVWXYZ]{8}$'
);

-- +goose Down
ALTER TABLE postgres_events DROP CONSTRAINT postgres_events_short_id_format;
ALTER TABLE postgres_events ADD COLUMN public_id TEXT;
UPDATE postgres_events SET public_id = 'p_' || lpad(short_id, 26, '0'), short_id = 'p_' || short_id;
ALTER TABLE postgres_events ALTER COLUMN public_id SET NOT NULL;
ALTER TABLE postgres_events ADD CONSTRAINT postgres_events_public_id_key UNIQUE (public_id);
ALTER TABLE postgres_events ADD CONSTRAINT postgres_events_public_id_format CHECK (
    public_id ~ '^p_[0-9A-HJKMNPQRSTVWXYZ]{26}$'
);
ALTER TABLE postgres_events ADD CONSTRAINT postgres_events_short_id_format CHECK (
    short_id ~ '^p_[0-9A-HJKMNPQRSTVWXYZ]{8}$'
);
