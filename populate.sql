-- USER generate
insert into users_tab
(full_name, phone, email, address, password, profile_picture)
VALUES
('John Doe', '081122334455', 'john.doe@mail.com', 'Jl. Kemerdekaan, Kec. Tanah Merah, Kab. Tanah Tua, Prov. Sumatera Utara', '$2a$10$X9iRj9AhH9KpDQHX9sK8N.x37Bfif8Y6ZCglNOuh48MUaJwwk4Pwy', 'https://i.pravatar.cc/300'),
('Alice Doe', '081122334466', 'alice.doe@mail.com', 'Jl. Kebangsaan, Kec. Tanah Merah, Kab. Tanah Tua, Prov. Sumatera Utara', '$2a$10$X9iRj9AhH9KpDQHX9sK8N.x37Bfif8Y6ZCglNOuh48MUaJwwk4Pwy', 'https://i.pravatar.cc/300'),
('Ben Doe', '081122334477', 'ben.doe@mail.com', 'Jl. Pahlawan, Kec. Biru Muda, Kab. Laut Biru, Prov. Sulawesi Tenggara', '$2a$10$X9iRj9AhH9KpDQHX9sK8N.x37Bfif8Y6ZCglNOuh48MUaJwwk4Pwy', 'https://i.pravatar.cc/300');

-- ADMIN generate
insert into users_tab
(full_name, phone, email, role, address, password, profile_picture)
VALUES
('Charlie Doe', '081122334488', 'charlie.doe@mail.com', 'ADMIN', 'Jl. Kemudahan, Kec. Biru Muda, Kab. Laut Biru, Prov. Sulawesi Tenggara', '$2a$10$X9iRj9AhH9KpDQHX9sK8N.x37Bfif8Y6ZCglNOuh48MUaJwwk4Pwy', 'https://i.pravatar.cc/300');

--Insert categories
insert into categories_tab (name)
values ('Coffee'), ('Non-Coffee'), ('Bread');

--Insert Menus
insert into menus_tab
(category_id, name, price, image)
VALUES
(1, 'Coffee Milk', 20000, 'https://i.pravatar.cc/300'),
(1, 'Cappuccino', 24000, 'https://i.pravatar.cc/300'),
(1, 'Coffee Milk With Brown Sugar', 25000, 'https://i.pravatar.cc/300'),
(1, 'Caramel Macchiato', 20000, 'https://i.pravatar.cc/300'),
(2, 'Iced Tea', 10000, 'https://i.pravatar.cc/300'),
(2, 'Orange Juice', 15000, 'https://i.pravatar.cc/300'),
(2, 'Watermelon Juice', 18000, 'https://i.pravatar.cc/300'),
(3, 'Choco Pan De', 24000, 'https://i.pravatar.cc/300'),
(3, 'Filipino Choco', 23000, 'https://i.pravatar.cc/300'),
(3, 'Garlic Cheese', 15000, 'https://i.pravatar.cc/300');

insert into menu_option_categories_tab
(name)
values ('size'), ('toppings'), ('available in');

insert into menu_options_tab
(category_id, menu_option_category_id, name, price)
values
(1, 1, 'Regular', 0),
(1, 1, 'Large', 6000),
(1, 2, 'Boba', 5000),
(1, 2, 'Jelly', 5000),
(1, 2, 'Ice Cream', 6000),
(1, 3, 'Ice', 0),
(1, 3, 'Hot', 0);

-- Insert promotions
insert into promotions_tab
(menu_id, name, description, image, discount, expired_date)
VALUES
(1, 'Coffee Milk Merdeka', 'Attractive offer for Coffee Milk menu on August 17th to get discount up to 20%', 'https://localhost/imageurl.jpg', 20, '2022-08-17 23:59:59Z+07:00'),
(10, 'Garlic Cheese Merdeka', 'Attractive offer for Garlic Cheese menu on August 17th to get discount up to 10%', 'https://localhost/imageurl.jpg', 10, '2022-08-17 23:59:59Z+07:00');

insert into coupons_tab (name, amount, is_available, min_spent, expired_date)
values
('Discount 10k with spent 30k', 10000, true, 30000, '2022-08-31 23:59:59Z+07:00'),
('Discount 5k with spent 20k', 5000, true, 20000, '2022-08-31 23:59:59Z+07:00');

insert into payment_options_tab
(name)
values ('Shopeepay'), ('Gopay'), ('OVO');

-- Insert games specific score and prize
insert into games_tab (coupon_id, target_score)
values (1, 100),(2, 50);

-- Insert favorite menu of users
insert into users_menus_favorite_tab
(user_id, menu_id)
values
(1, 1),
(1, 2),
(2, 8),
(3, 10),
(3, 1),
(3, 4);

-- Insert owned coupons of user
insert into users_coupons_tab
(user_id, coupon_id, redeemed_date)
values
(1, 1, null),
(1, 2, null),
(2, 2, null);
