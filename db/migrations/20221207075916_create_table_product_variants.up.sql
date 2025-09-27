CREATE TABLE IF NOT EXISTS product_variants(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    product_id BIGINT NOT NULL,
    unit_id BIGINT,
    unit_size FLOAT,
    type VARCHAR(100) NOT NULL,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    price NUMERIC,
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    updated_at BIGINT
);

ALTER TABLE product_variants
    ADD CONSTRAINT fk_products_product_variants
        FOREIGN KEY (product_id) REFERENCES products(id)
            ON DELETE CASCADE;

ALTER TABLE product_variants
    ADD CONSTRAINT fk_units_product_variants
        FOREIGN KEY (unit_id) REFERENCES units(id)
            ON DELETE SET NULL;