CREATE TABLE IF NOT EXISTS customers
(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    email VARCHAR(255) UNIQUE,
    phone VARCHAR(255) UNIQUE,
    description VARCHAR(255),
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    updated_at BIGINT
);