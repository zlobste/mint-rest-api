DROP TABLE users_organizations;
DROP TABLE payment_details;
DROP TABLE orders;
DROP TABLE dishes;
DROP TABLE menu;
DROP TABLE users;
DROP TABLE organizations;

CREATE TABLE users (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    name VARCHAR NOT NULL,
    email VARCHAR NOT NULL UNIQUE,
    password VARCHAR NOT NULL,
    role  VARCHAR NOT NULL,
    balance MONEY DEFAULT 0,
    blocked BOOLEAN DEFAULT false
);

CREATE TABLE dishes (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    title VARCHAR NOT NULL,
    description VARCHAR NOT NULL,
    cost MONEY NOT NULL,
    disabled BOOLEAN DEFAULT false
);

CREATE TABLE orders (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    cost MONEY NOT NULL,
    datetime TIMESTAMP WITH TIME ZONE NOT NULL,
    canceled BOOLEAN NOT NULL DEFAULT FALSE,
    dish_id BIGSERIAL NOT NULL,
    user_id BIGSERIAL NOT NULL,
    in_progress BOOLEAN NOT NULL DEFAULT FALSE,
    FOREIGN KEY (dish_id) REFERENCES dishes (id) ON DELETE CASCADE,
    FOREIGN KEY (user_id) REFERENCES users (id) ON DELETE CASCADE
);

CREATE TABLE institutions (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    title VARCHAR NOT NULL,
    address VARCHAR NOT NULL,
    disabled BOOLEAN DEFAULT false
);


CREATE TABLE payment_details (
    id BIGSERIAL NOT NULL PRIMARY KEY,
    bank VARCHAR NOT NULL,
    account VARCHAR NOT NULL UNIQUE,
    institution_id BIGSERIAL NOT NULL,
    FOREIGN KEY (institution_id) REFERENCES institutions (id) ON DELETE CASCADE
);