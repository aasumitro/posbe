INSERT INTO floors (name)
VALUES ('1st');

INSERT INTO tables (floor_id, name, x_pos, y_pos, w_size, h_size, d_size, capacity, type)
VALUES (1, 'A1', 176,58,50 , 50, 0, 4, 'rectangle'),
        (1, 'A2',174,396,0 , 0, 50, 4, 'circle'),
        (1, 'A3',189,692,0 , 0, 50, 3, 'circle'),
        (1, 'A4',600,58,0 , 0, 50, 5, 'circle'),
        (1, 'A5',608,409,0 , 0, 50, 6, 'circle'),
        (1, 'A6',618,785,0 , 0, 50, 2, 'circle');

-- pos_type : none, restaurant, coffee_shop, store, karaoke
INSERT INTO store_prefs (key, value)
VALUES
    ('name', 'Lorem Store'),
    ('phone', '+6282266668810'),
    ('email', 'lorem@store.posbe'),
    ('type', 'restaurant'), -- restaurant, bar, coffee, store, karaoke
    ('address', 'Jalan Merdeka 12 <> Ruko A, Megamas. <> kota manado <> sulawesi utara <> indonesia <> 95711'),
    ('currency', 'IDR'), -- IDR/USD
    ('service_rate', '5'), -- in percentage
    ('service_category', 'standard'),
    ('tax_rate', '10'), -- in percentage
    ('tax_category', 'standard'),
    ('feature_floor', '1');  -- true or false
