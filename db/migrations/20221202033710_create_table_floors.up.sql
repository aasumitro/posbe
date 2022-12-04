CREATE TABLE IF NOT EXISTS floors (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255),
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    updated_at BIGINT
);