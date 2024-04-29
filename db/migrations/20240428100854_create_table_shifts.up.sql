CREATE TABLE IF NOT EXISTS shifts(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    name VARCHAR(255) NOT NULL,
    start_time BIGINT NOT NULL,
    end_time BIGINT NOT NULL,
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    updated_at BIGINT
);

CREATE TABLE IF NOT EXISTS store_shifts(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    shift_id BIGINT NOT NULL,
    open_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    open_by BIGINT NOT NULL,
    open_cash BIGINT NOT NULL,
    close_at BIGINT,
    close_by BIGINT,
    close_cash BIGINT,
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    updated_at BIGINT
);

ALTER TABLE store_shifts ADD CONSTRAINT fk_shift_store_shifts
    FOREIGN KEY (shift_id) REFERENCES shifts(id);
