INSERT INTO categories (name)
VALUES ('foods'), ('beverages');

INSERT INTO subcategories (category_id, name)
VALUES (1, 'meat'),
       (1, 'seafood'),
       (1, 'chicken'),
       (2, 'coffee'),
       (2, 'tea'),
       (2, 'juice');

INSERT INTO units (magnitude, name, symbol)
VALUES ('mass', 'gram', 'g'),
       ('mass', 'milligram', 'mg'),
       ('mass', 'kilogram', 'kg'),
       ('mass', 'milliliter', 'ml'),
       ('mass', 'liter', 'l'),
       ('mass', 'portion', 'prt');
