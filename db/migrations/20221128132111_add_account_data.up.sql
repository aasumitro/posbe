INSERT INTO roles (name, description)
VALUES
    ('admin', 'admin level can access all of the features/menus'),
    ('cashier', 'cashier level can access room, order & payment menu'),
    ('waiter', 'waiter level can access room & order menu');

-- password is secret
INSERT INTO users(role_id, name, username, email, password)
VALUES
    (1, 'Test Admin', 'admin', 'admin@posbe.test', '4c388700ed16de8681d5f5b4785c94f60af20dd765a5acade2b0f8e86357315a.dda885bab15f14eebdeabfde43173ee0de14db0a3a8e65438974678a4b3a5135'),
    (2, 'Test Cashier', 'cashier', 'cashier@posbe.test', '4c388700ed16de8681d5f5b4785c94f60af20dd765a5acade2b0f8e86357315a.dda885bab15f14eebdeabfde43173ee0de14db0a3a8e65438974678a4b3a5135'),
    (3, 'Test Waiter', 'waiter', 'waiter@posbe.test', '4c388700ed16de8681d5f5b4785c94f60af20dd765a5acade2b0f8e86357315a.dda885bab15f14eebdeabfde43173ee0de14db0a3a8e65438974678a4b3a5135'),
    (3, 'ToBe Removed', 'tbr', 'tbr@posbe.test', '4c388700ed16de8681d5f5b4785c94f60af20dd765a5acade2b0f8e86357315a.dda885bab15f14eebdeabfde43173ee0de14db0a3a8e65438974678a4b3a5135');

