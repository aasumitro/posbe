CREATE TABLE IF NOT EXISTS addons(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255),
    description VARCHAR(255),
    price FLOAT,
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    updated_at BIGINT
);
