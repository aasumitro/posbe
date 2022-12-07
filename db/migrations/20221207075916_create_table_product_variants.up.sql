CREATE TYPE variant_types AS ENUM ('none', 'size');

CREATE TABLE IF NOT EXISTS product_variants(
    id BIGSERIAL PRIMARY KEY NOT NULL,
    product_id BIGINT NOT NULL,
    unit_id BIGINT NOT NULL,
    unit_size FLOAT,
    type VARIANT_TYPES DEFAULT 'none',
    name VARCHAR(255),
    description VARCHAR(255),
    price FLOAT
);

ALTER TABLE product_variants ADD CONSTRAINT fk_products_product_variants
    FOREIGN KEY (product_id) REFERENCES products(id);

ALTER TABLE product_variants ADD CONSTRAINT fk_units_product_variants
    FOREIGN KEY (unit_id) REFERENCES units(id);