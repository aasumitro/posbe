CREATE TABLE IF NOT EXISTS store_prefs (
    key VARCHAR(255) UNIQUE NOT NULL,
    value VARCHAR(255),
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    updated_at BIGINT
);
