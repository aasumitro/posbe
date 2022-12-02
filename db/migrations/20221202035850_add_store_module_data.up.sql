INSERT INTO floors (name)
VALUES ('1st'), ('2nd');

INSERT INTO tables (floor_id, name, x_pos, y_pos, w_size, h_size, capacity)
VALUES (1, 'A1', 0, 0, 4 , 4, 4);

INSERT INTO rooms (floor_id, name, x_pos, y_pos, w_size, h_size, capacity, price)
VALUES (2, 'R1', 0, 0, 4 , 4, 4, 100);

-- pos_type : none, restaurant, coffee_shop, store, karaoke
INSERT INTO store_prefs (key, value)
VALUES ('type', ''),
       ('feature_floor', ''),  -- true or false
       ('feature_room', ''),  -- true or false
       ('feature_table', ''),  -- true or false
       ('feature_', ''),  -- true or false
       ('theme', ''),  -- dark or light
       ('lang', ''),  -- en_US or id_ID
       ('locale', ''); -- Asia/Jayapura Asia/Makassar, Asia/Jakarta