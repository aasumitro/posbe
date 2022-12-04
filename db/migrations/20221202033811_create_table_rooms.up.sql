CREATE TABLE IF NOT EXISTS rooms (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    floor_id BIGINT,
    name VARCHAR(255),
    x_pos FLOAT,
    y_pos FLOAT,
    w_size FLOAT,
    h_size FLOAT,
    capacity INT,
    price FLOAT,
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    updated_at BIGINT
);

ALTER TABLE rooms ADD CONSTRAINT fk_floors_rooms
    FOREIGN KEY (floor_id) REFERENCES floors(id);