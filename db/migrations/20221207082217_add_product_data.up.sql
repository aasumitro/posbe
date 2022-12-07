INSERT INTO addons (name, description, price)
VALUES ('oat milk', 'replace', 1),
       ('raw milk', 'replace', 1),
       ('cheese', 'extra cheese', 1),
       ('chocolate', 'extra chocolate', 1);

INSERT INTO products (category_id, subcategory_id, sku, name, description, price)
VALUES (2, 4, 'JMGO100', 'mango juice', 'this sweet, tangy, and fruity tropical juice can be made using a blender, handheld blender, or a food processor in under 5 minutes.', 25),
       (1, 2, 'WA5S100', 'wagyu a5 steak', 'The highest yield grade and meat quality grade for Wagyu beef is A5, where A represents the yield grade, and 5 represents the meat quality grade. A5 Wagyu beef denotes meat with ideal firmness and texture, coloring, yield, and beef marbling score.', 100);

INSERT INTO product_variants (product_id, type, name, description, unit_id, unit_size, price)
VALUES (1, 'size', 's', 'small', 4, 250, 0),
       (1, 'size', 'm', 'medium', 4, 480, 2),
       (1, 'size', 'l', 'large', 4, 650, 3),
       (1, 'size', 'xl', 'extra large', 5, 1.5, 1),
       (2, 'size', 'half', 'half portion', 1, 250, 0),
       (2, 'size', 'normal', 'normal portion', 1, 500, 100);
