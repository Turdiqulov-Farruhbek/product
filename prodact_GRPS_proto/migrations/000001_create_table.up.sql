-- product, category, user, cart, transaction
create table category
(
    id   uuid primary key,
    name varchar(36),
    created_at  timestamp default now(),
    updated_at  timestamp default now()
);

create table product
(
    id          uuid primary key,
    name        varchar(36),
    price       int,
    category_id uuid references category (id),
    expired_at  timestamp,
    created_at  timestamp default now(),
    updated_at  timestamp default now(),
    deleted_at  bigint    default 0
);

create table users
(
    id         uuid primary key,
    name       varchar(36),
    balance    int       default 0,
    email      varchar,
    phone      varchar,
    created_at timestamp default now(),
    updated_at timestamp default now(),
    deleted_at bigint    default 0
);

create table basket
(
    id      uuid primary key,
    user_id uuid references users (id)
);

create table item
(
    id         uuid primary key,
    basket_id  uuid references basket (id),
    product_id uuid references product (id),
    quantity   int
);

create table transaction
(
    id          uuid primary key,
    basket_id   uuid references basket (id),
    total_price int default 0
);