CREATE TABLE IF NOT EXISTS floors (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255),
    created_at timestamptz NOT NULL DEFAULT (now()),
    updated_at timestamptz
);