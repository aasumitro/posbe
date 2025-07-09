CREATE TABLE IF NOT EXISTS products (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    category_id BIGINT NOT NULL,
    subcategory_id BIGINT NOT NULL,
    sku VARCHAR(255) UNIQUE NOT NULL,
    image VARCHAR(255),
    gallery TEXT,
    name VARCHAR(255) NOT NULL,
    description VARCHAR(255),
    -- we don't need it (price), let's put this item into variant
    -- by default when user create new item we will add new base variant
    -- price NUMERIC NOT NULL,
    created_at BIGINT NOT NULL DEFAULT extract(epoch from now()),
    updated_at BIGINT
);

ALTER TABLE products ADD CONSTRAINT fk_products_categories
    FOREIGN KEY (category_id) REFERENCES categories(id);

ALTER TABLE products ADD CONSTRAINT fk_products_subcategories
    FOREIGN KEY (subcategory_id) REFERENCES subcategories(id);
