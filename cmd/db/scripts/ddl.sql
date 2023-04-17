CREATE SCHEMA IF NOT EXISTS bshop;

CREATE TABLE IF NOT EXISTS bshop.Users (
    id INTEGER PRIMARY KEY,
    name VARCHAR(200),
    email VARCHAR(200) NOT NULL,
    password VARCHAR(100) NOT NULL
);

CREATE TABLE IF NOT EXISTS bshop.Products (
    id INTEGER PRIMARY KEY,
    name VARCHAR(200),
    category VARCHAR(200),
    price NUMERIC NOT NULL,
    stock INTEGER,
    country VARCHAR(100)
);