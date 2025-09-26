CREATE TYPE order_status AS ENUM ('none', 'booking', 'order', 'bill', 'paid', 'cancel');

CREATE TABLE IF NOT EXISTS orders (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    cashier_id BIGINT,
    shift_id BIGINT,
    table_id BIGINT,
    time_open BIGINT,
    time_close BIGINT,
    customer VARCHAR(255),
    gross NUMERIC,
    discount NUMERIC,
    net NUMERIC,
    tax NUMERIC,
    total NUMERIC,
    type VARCHAR(255),
    payment NUMERIC,
    change NUMERIC,
    notes TEXT,
    cancel_reason TEXT,
    status order_status DEFAULT 'none',
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    updated_at BIGINT
);

ALTER TABLE orders ADD CONSTRAINT fk_order_cashier
    FOREIGN KEY (cashier_id) REFERENCES users(id);
ALTER TABLE orders ADD CONSTRAINT fk_order_shift
    FOREIGN KEY (shift_id) REFERENCES active_shifts(id);
ALTER TABLE orders ADD CONSTRAINT fk_order_table
    FOREIGN KEY (table_id) REFERENCES tables(id);

CREATE TABLE IF NOT EXISTS order_products (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    order_id BIGINT,
    product_id BIGINT,
    category_id BIGINT,
    subcategory_id BIGINT,
    variant_id BIGINT,
    name VARCHAR(255),
    quantity INT DEFAULT 1,
    price NUMERIC,
    netto NUMERIC,
    notes TEXT,
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    updated_at BIGINT
);

ALTER TABLE order_products ADD CONSTRAINT fk_order_product_order
    FOREIGN KEY (order_id) REFERENCES orders(id);
ALTER TABLE order_products ADD CONSTRAINT fk_order_product_product
    FOREIGN KEY (product_id) REFERENCES products(id);
ALTER TABLE order_products ADD CONSTRAINT fk_order_product_category
    FOREIGN KEY (category_id) REFERENCES categories(id);
ALTER TABLE order_products ADD CONSTRAINT fk_order_product_subcategory
    FOREIGN KEY (subcategory_id) REFERENCES subcategories(id);
ALTER TABLE order_products ADD CONSTRAINT fk_order_product_variant
    FOREIGN KEY (variant_id) REFERENCES product_variants(id);

CREATE TABLE IF NOT EXISTS order_product_addons (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    order_id BIGINT,
    order_product_id BIGINT,
    addon_id BIGINT,
    name VARCHAR(255),
    quantity INT DEFAULT 1,
    price NUMERIC,
    netto NUMERIC,
    notes TEXT,
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    updated_at BIGINT
);

ALTER TABLE order_product_addons ADD CONSTRAINT fk_order_product_addons_order
    FOREIGN KEY (order_id) REFERENCES orders(id);
ALTER TABLE order_product_addons ADD CONSTRAINT fk_order_product_addons_product
    FOREIGN KEY (order_product_id) REFERENCES order_products(id);
ALTER TABLE order_product_addons ADD CONSTRAINT fk_order_product_addons_addon
    FOREIGN KEY (addon_id) REFERENCES product_addons(id);
