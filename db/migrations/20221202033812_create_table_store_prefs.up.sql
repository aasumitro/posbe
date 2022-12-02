CREATE TABLE IF NOT EXISTS store_prefs (
    key VARCHAR(255) UNIQUE NOT NULL,
    value VARCHAR(255),
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz
);
