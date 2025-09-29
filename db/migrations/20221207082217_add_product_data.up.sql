INSERT INTO product_addons (name, description, price)
VALUES ('oat milk', 'replace', 1),
       ('raw milk', 'replace', 1),
       ('cheese', 'extra cheese', 1),
       ('chocolate', 'extra chocolate', 1);

INSERT INTO products (category_id, subcategory_id, sku, name, description)
VALUES (2, 4, 'JMGO100', 'mango juice', 'this sweet, tangy, and fruity tropical juice can be made using a blender, handheld blender, or a food processor in under 5 minutes.'),
       (1, 1, 'WA5S100', 'wagyu a5 steak', 'The highest yield grade and meat quality grade for Wagyu beef is A5, where A represents the yield grade, and 5 represents the meat quality grade. A5 Wagyu beef denotes meat with ideal firmness and texture, coloring, yield, and beef marbling score.'),
       (1, 3, 'FCL100', 'fried chicken lalapan', 'Chicken Lalapan is fried/grilled chicken + rice + fresh veggies + sambal → served with lalapan (raw or lightly blanched vegetables) and sambal (spicy chili paste).');

INSERT INTO product_variants (product_id, type, name, description, unit_id, unit_size, price)
VALUES (1, 'sizes', 's', 'small', 4, 250, 12),
       (1, 'sizes', 'm', 'medium', 4, 480, 14),
       (1, 'sizes', 'l', 'large', 4, 650,16),
       (1, 'sizes', 'xl', 'extra large', 5, 1.5, 18),
       (2, 'sizes', 'h', 'half portion', 1, 250, 25),
       (2, 'sizes', 'n', 'normal portion', 1, 500, 50),
       (2, 'doneness', 'wd', 'well done', null, null, 0),
       (2, 'doneness', 'md', 'medium done', null, null, 0),
       (2, 'doneness', 'mr', 'medium rare', null, null, 0),
       (2, 'doneness', 'm', 'medium', null, null, 0),
       (2, 'doneness', 'r', 'rare', null, null, 0),
       (3, 'base', 'base', 'base', 6, 1, 25000);
