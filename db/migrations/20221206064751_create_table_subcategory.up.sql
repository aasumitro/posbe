CREATE TABLE IF NOT EXISTS subcategories (
    id BIGSERIAL PRIMARY KEY NOT NULL,
    category_id BIGINT,
    name VARCHAR(50)
);

ALTER TABLE subcategories ADD CONSTRAINT fk_categories_subcategories
    FOREIGN KEY (category_id) REFERENCES categories(id);