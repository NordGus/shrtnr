CREATE TABLE IF NOT EXISTS urls (
    id TEXT PRIMARY KEY NOT NULL,
    uuid TEXT UNIQUE NOT NULL,
    target TEXT UNIQUE NOT NULL,
    created_at INTEGER NOT NULL,
    deleted_at INTEGER,
    CONSTRAINT urls_uuid_target_uniqueness_index UNIQUE (uuid, target) ON CONFLICT ROLLBACK
) STRICT;
