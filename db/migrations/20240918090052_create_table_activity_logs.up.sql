CREATE TABLE IF NOT EXISTS activity_logs (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    user_id BIGINT,
    status_code INT,
    role VARCHAR(10),
    description TEXT,
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now())
);