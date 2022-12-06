INSERT INTO roles (name, description)
VALUES
    ('admin', 'admin level can access all of the features/menus'),
    ('cashier', 'cashier level can access room, order & payment menu'),
    ('waiter', 'waiter level can access room & order menu');

-- password is secret
INSERT INTO users(role_id, name, username, email, phone, password)
VALUES
    (1, 'A. A. Sumitro', 'aasumitro', 'hello@aasumitro.id', 82271115593, '2ad1a22d5b3c9396d16243d2fe7f067976363715e322203a456278bb80b0b4a4.7ab4dcccfcd9d36efc68f1626d2fb80804a6508f9c3a7b44f430ba082b6870d2')

