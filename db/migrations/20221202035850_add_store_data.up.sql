INSERT INTO floors (name)
VALUES ('1st'), ('2nd');

INSERT INTO tables (floor_id, name, x_pos, y_pos, w_size, h_size, capacity)
VALUES (1, 'A1', 0, 0, 4 , 4, 4);

INSERT INTO rooms (floor_id, name, x_pos, y_pos, w_size, h_size, capacity, price)
VALUES (2, 'R1', 0, 0, 4 , 4, 4, 100);

-- pos_type : none, restaurant, coffee_shop, store, karaoke
INSERT INTO store_prefs (key, value)
VALUES ('pos_type', 'restaurant'), -- restaurant, bar, coffee, store, karaoke
       ('feature_floor', '1'),  -- true or false
       ('feature_room', '0'),  -- true or false
       ('feature_table', '1'),  -- true or false
       -- ('feature_', ''),  -- true or false
       ('fe_theme', 'dark'),  -- dark or light
       ('fe_lang', 'id'),  -- en_US or id_ID
       ('fe_locale', 'Asia/Makassar'), -- Asia/Jayapura Asia/Makassar, Asia/Jakarta
       ('currency', 'IDR'), -- IDR/USD
       ('currency_rate', '15400'); -- TO USD