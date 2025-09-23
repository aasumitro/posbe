INSERT INTO floors (name)
VALUES ('1st');

INSERT INTO tables (floor_id, name, x_pos, y_pos, w_size, h_size, d_size, capacity, type)
VALUES (1, 'A1', 0, 0, 50 , 50, 0, 4, 'rectangle'),
        (1, 'A2', 100, 0, 0 , 0, 50, 4, 'circle');

-- pos_type : none, restaurant, coffee_shop, store, karaoke
INSERT INTO store_prefs (key, value)
VALUES
    ('name', 'Lorem Store'),
    ('phone', '+62872222'),
    ('email', 'lorem@store.id'),
    ('type', 'restaurant'), -- restaurant, bar, coffee, store, karaoke
    ('address', 'Jalan Suka Maju'),
    ('currency', 'IDR'), -- IDR/USD
    ('service_rate', '5'), -- in percentage
    ('service_category', 'standard'),
    ('tax_rate', '10'), -- in percentage
    ('tax_category', 'standard'),
    ('feature_floor', '1'),  -- true or false
