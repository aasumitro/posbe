CREATE TYPE order_status AS ENUM ('none', 'booking', 'order', 'bill', 'paid', 'cancel');

CREATE TYPE payment_method as ENUM ('cash', 'card', 'e-wallet', 'qris');

CREATE TABLE IF NOT EXISTS orders (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    cashier_id BIGINT,
    shift_id BIGINT,
    table_id BIGINT,
    time_open BIGINT,
    time_close BIGINT,
    customer VARCHAR(255),
    gross NUMERIC(12,2),
    discount NUMERIC(12,2),
    net NUMERIC(12,2),
    tax NUMERIC(12,2),
    total NUMERIC(12,2),
    payment_method payment_method,
    payment_amount NUMERIC(12,2),
    change NUMERIC(12,2),
    notes TEXT,
    cancel_reason TEXT,
    status order_status DEFAULT 'none',
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    updated_at BIGINT -- will be set on repo
);

ALTER TABLE orders
    ADD CONSTRAINT fk_order_cashier
        FOREIGN KEY (cashier_id) REFERENCES users(id)
            ON DELETE SET NULL;
ALTER TABLE orders
    ADD CONSTRAINT fk_order_shift
        FOREIGN KEY (shift_id) REFERENCES active_shifts(id)
            ON DELETE SET NULL;
ALTER TABLE orders
    ADD CONSTRAINT fk_order_table
        FOREIGN KEY (table_id) REFERENCES tables(id)
            ON DELETE SET NULL;

CREATE TABLE IF NOT EXISTS order_products (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    order_id BIGINT,
    product_id BIGINT,
    category_id BIGINT,
    subcategory_id BIGINT,
    name VARCHAR(255),
    quantity INT DEFAULT 1,
    price NUMERIC(12,2),
    net NUMERIC(12,2),
    notes TEXT,
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    updated_at BIGINT -- will be set on repo
);

ALTER TABLE order_products
    ADD CONSTRAINT fk_order_product_order
        FOREIGN KEY (order_id) REFERENCES orders(id)
            ON DELETE CASCADE;
ALTER TABLE order_products
    ADD CONSTRAINT fk_order_product_product
        FOREIGN KEY (product_id) REFERENCES products(id)
            ON DELETE CASCADE;
ALTER TABLE order_products
    ADD CONSTRAINT fk_order_product_category
        FOREIGN KEY (category_id) REFERENCES categories(id)
            ON DELETE SET NULL;
ALTER TABLE order_products
    ADD CONSTRAINT fk_order_product_subcategory
        FOREIGN KEY (subcategory_id) REFERENCES subcategories(id)
            ON DELETE SET NULL;

CREATE TABLE IF NOT EXISTS order_product_options (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    order_id BIGINT,
    order_product_id BIGINT,
    variant_id BIGINT,
    name VARCHAR(255),
    value VARCHAR(255),
    price NUMERIC(12,2),
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    updated_at BIGINT -- will be set on repo
);

ALTER TABLE order_product_options
    ADD CONSTRAINT fk_order_product_options_order
        FOREIGN KEY (order_id) REFERENCES orders(id)
            ON DELETE CASCADE;
ALTER TABLE order_product_options
    ADD CONSTRAINT fk_order_product_options_product
        FOREIGN KEY (order_product_id) REFERENCES order_products(id)
            ON DELETE CASCADE;
ALTER TABLE order_product_options
    ADD CONSTRAINT fk_order_product_options_variant
        FOREIGN KEY (variant_id) REFERENCES product_variants(id)
            ON DELETE CASCADE;

CREATE TABLE IF NOT EXISTS order_product_addons (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    order_id BIGINT,
    order_product_id BIGINT,
    addon_id BIGINT,
    name VARCHAR(255),
    quantity INT DEFAULT 1,
    price NUMERIC(12,2),
    net NUMERIC(12,2),
    notes TEXT,
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    updated_at BIGINT -- will be set on repo
);

ALTER TABLE order_product_addons
    ADD CONSTRAINT fk_order_product_addons_order
        FOREIGN KEY (order_id) REFERENCES orders(id)
            ON DELETE CASCADE;
ALTER TABLE order_product_addons
    ADD CONSTRAINT fk_order_product_addons_product
        FOREIGN KEY (order_product_id) REFERENCES order_products(id)
            ON DELETE CASCADE;
ALTER TABLE order_product_addons
    ADD CONSTRAINT fk_order_product_addons_addon
        FOREIGN KEY (addon_id) REFERENCES product_addons(id)
            ON DELETE CASCADE;

CREATE INDEX idx_orders_status ON orders(status);
CREATE INDEX idx_orders_shift_id ON orders(shift_id);
CREATE INDEX idx_order_products_order_id ON order_products(order_id);
CREATE INDEX idx_orders_created_at ON orders(created_at);
CREATE INDEX idx_order_products_product_id ON order_products(product_id);
CREATE INDEX idx_order_product_options_order_product_id ON order_product_options(order_product_id);
CREATE INDEX idx_order_product_addons_order_product_id ON order_product_addons(order_product_id);
CREATE UNIQUE INDEX uniq_order_product_option
    ON order_product_options (order_product_id, name, value);
CREATE UNIQUE INDEX uniq_order_product_addon
    ON order_product_addons (order_product_id, addon_id);