CREATE TYPE table_types AS ENUM ('rectangle', 'circle');

CREATE TABLE IF NOT EXISTS tables (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    floor_id BIGINT,
    name VARCHAR(255),
    x_pos FLOAT,
    y_pos FLOAT,
    w_size FLOAT, -- this for rectangle
    h_size FLOAT, -- this for rectangle
    d_size FLOAT, -- this for circle d stand for diameter
    capacity INT,
    type TABLE_TYPES DEFAULT 'rectangle',
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    updated_at BIGINT
);

ALTER TABLE tables
    ADD CONSTRAINT fk_floors_tables
        FOREIGN KEY (floor_id) REFERENCES floors(id)
            ON DELETE SET NULL;