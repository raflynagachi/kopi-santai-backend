create database kopi_santai_db;

drop table if exists game_leaderboards_tab;
drop table if exists games_tab;
drop table if exists order_items_tab;
drop table if exists orders_tab;
drop table if exists deliveries_tab;
drop table if exists reviews_tab;
drop table if exists promotions_tab;
drop table if exists menu_options_categories_tab;
drop table if exists menu_options_tab;
drop table if exists menus_tab;
drop table if exists categories_tab;
drop table if exists users_coupons_tab;
drop table if exists coupons_tab;
drop table if exists payment_options_tab;
drop table if exists users_tab;

create table users_tab(
    id bigserial primary key ,
    full_name varchar not null ,
    phone varchar not null ,
    email varchar not null unique ,
    role varchar default 'USER' ,
    address varchar not null ,
    password varchar not null ,
    profile_picture bytea not null ,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp
);

create table payment_options_tab(
    id bigserial primary key ,
    name varchar not null ,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp
);

create table coupons_tab(
    id bigserial primary key ,
    name varchar not null ,
    amount decimal not null ,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp
);

create table users_coupons_tab(
    id bigserial primary key ,
    user_id bigint not null ,
    coupon_id bigint not null ,
    foreign key (user_id) references users_tab(id) ,
    foreign key (coupon_id) references coupons_tab(id) on delete cascade
);

create table categories_tab(
    id bigserial primary key ,
    name varchar not null ,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp
);

create table menus_tab(
    id bigserial primary key ,
    category_id bigint not null ,
    name varchar not null ,
    price decimal not null ,
    image bytea not null ,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp,
    foreign key (category_id) references categories_tab(id)
);

create table menu_options_tab(
    id bigserial primary key ,
    name varchar not null ,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp
);

create table menu_options_categories_tab(
    id bigserial primary key ,
    category_id bigint not null ,
    menu_option_id bigint not null ,
    name varchar not null ,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp,
    foreign key (category_id) references categories_tab(id),
    foreign key (menu_option_id) references menu_options_tab(id)
);

create table reviews_tab(
    id bigserial not null ,
    user_id bigint not null ,
    menu_id bigint not null ,
    description varchar ,
    rating decimal not null ,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp,
    foreign key (user_id) references users_tab(id),
    foreign key (menu_id) references menus_tab(id),
    primary key (user_id, menu_id)
);

create table deliveries_tab(
    id bigserial primary key ,
    delivery_date timestamp not null ,
    status varchar default 'IN PROCESS',
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp
);

create table promotions_tab(
    id bigserial primary key ,
    coupon_id bigint not null ,
    name varchar not null ,
    description varchar,
    image bytea not null ,
    min_spent decimal not null ,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp,
    foreign key (coupon_id) references coupons_tab(id) on delete cascade
);

create table orders_tab(
    id bigserial primary key ,
    user_id bigint not null ,
    coupon_id bigint ,
    delivery_id bigint not null ,
    payment_option_id bigint not null ,
    ordered_date timestamp not null ,
    total_price decimal not null ,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp,
    foreign key (user_id) references users_tab(id),
    foreign key (coupon_id) references coupons_tab(id) on delete set null,
    foreign key (delivery_id) references deliveries_tab(id),
    foreign key (payment_option_id) references payment_options_tab(id)
);

create table order_items_tab(
    id bigserial primary key ,
    user_id bigint not null ,
    menu_id bigint not null ,
    order_id bigint ,
    quantity int not null ,
    description varchar ,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp,
    foreign key (user_id) references users_tab(id),
    foreign key (menu_id) references menus_tab(id),
    foreign key (order_id) references orders_tab(id)
);

create table games_tab(
    id bigserial primary key ,
    coupon_id bigint not null ,
    target_score int not null ,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp,
    foreign key (coupon_id) references coupons_tab(id) on delete set null
);

create table game_leaderboards_tab(
    id bigserial primary key ,
    user_id bigint not null ,
    score bigint default 0,
    tried int default 0,
    created_at timestamp default current_timestamp,
    updated_at timestamp default current_timestamp,
    deleted_at timestamp,
    foreign key (user_id) references users_tab(id)
);
