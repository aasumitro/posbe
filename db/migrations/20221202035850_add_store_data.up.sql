INSERT INTO floors (name)
VALUES ('1st');

INSERT INTO tables (floor_id, name, x_pos, y_pos, w_size, h_size, d_size, capacity, type)
VALUES (1, 'A1', 0,0,50 , 50, 0, 4, 'rectangle'),
        (1, 'A2',0,300,0 , 0, 50, 4, 'circle'),
        (1, 'A3',14,602,0 , 0, 50, 3, 'circle'),
        (1, 'A4',400,0,0 , 0, 50, 5, 'circle'),
        (1, 'A5',403,322,0 , 0, 50, 6, 'circle'),
        (1, 'A6',425,695,0 , 0, 50, 2, 'circle');

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
    ('feature_floor', '1'),  -- true or false
