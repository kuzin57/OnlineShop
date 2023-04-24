CREATE SCHEMA IF NOT EXISTS bshop;

<<<<<<< HEAD
CREATE TABLE IF NOT EXISTS bshop.user (
  user_id SERIAL UNIQUE,
  firstname VARCHAR(100) NOT NULL,
  surname VARCHAR(100) NOT NULL,
  phone_number VARCHAR(20) NOT NULL,
  email VARCHAR(200) UNIQUE NOT NULL,
=======
CREATE TABLE bshop.user (
  user_id INTEGER PRIMARY KEY,
  firstname VARCHAR(100) NOT NULL,
  surname VARCHAR(100) NOT NULL,
  phone_number VARCHAR(20) NOT NULL,
  email VARCHAR(200) NOT NULL,
>>>>>>> 793bb79 (ddl added)
  birthdate DATE NOT NULL,
  hashed_password VARCHAR(200) NOT NULL
);

<<<<<<< HEAD
CREATE TABLE IF NOT EXISTS bshop.product (
  product_id SERIAL UNIQUE,
=======
CREATE TABLE bshop.product (
  product_id INTEGER PRIMARY KEY,
>>>>>>> 793bb79 (ddl added)
  category VARCHAR(100) NOT NULL,
  name VARCHAR(100) NOT NULL,
  brand VARCHAR(50),
  price INTEGER NOT NULL,
  available BOOLEAN NOT NULL,
<<<<<<< HEAD
  rating NUMERIC NOT NULL,
  CHECK (rating <= 10)
);

CREATE TABLE IF NOT EXISTS bshop.review (
  review_id SERIAL UNIQUE,
=======
  rating NUMERIC NOT NULL
);

CREATE TABLE bshop.review (
  review_id INTEGER PRIMARY KEY,
>>>>>>> 793bb79 (ddl added)
  user_id INTEGER NOT NULL,
  comment TEXT,
  product_id INTEGER NOT NULL,
  date DATE NOT NULL,
  mark INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES bshop.user(user_id) ON DELETE CASCADE,
<<<<<<< HEAD
  FOREIGN KEY (product_id) REFERENCES bshop.product(product_id) ON DELETE CASCADE,
  CHECK (mark >= 1 and mark <= 10)
);

CREATE TABLE IF NOT EXISTS bshop.purchase (
  purchase_id SERIAL UNIQUE,
  client_id INTEGER NOT NULL,
  purchase_price INTEGER NOT NULL,
=======
  FOREIGN KEY (product_id) REFERENCES bshop.product(product_id) ON DELETE CASCADE
);

CREATE TABLE bshop.purchase (
  purchase_id INTEGER PRIMARY KEY,
  client_id INTEGER NOT NULL,
  purchas_eprice INTEGER NOT NULL,
>>>>>>> 793bb79 (ddl added)
  date DATE NOT NULL,
  FOREIGN KEY (client_id) REFERENCES bshop.user(user_id) ON DELETE CASCADE
);

<<<<<<< HEAD
CREATE TABLE IF NOT EXISTS bshop.purchase_product (
  pp_id SERIAL UNIQUE,
=======
CREATE TABLE bshop.purchase_product (
  pp_id INTEGER PRIMARY KEY,
>>>>>>> 793bb79 (ddl added)
  purchase_id INTEGER NOT NULL,
  product_id INTEGER NOT NULL,
  amount TEXT NOT NULL,
  FOREIGN KEY (purchase_id) REFERENCES bshop.purchase(purchase_id) ON DELETE CASCADE,
  FOREIGN KEY (product_id) REFERENCES bshop.product(product_id) ON DELETE CASCADE
);

<<<<<<< HEAD
CREATE TABLE IF NOT EXISTS bshop.user_history (
=======
CREATE TABLE bshop.user_history (
>>>>>>> 793bb79 (ddl added)
  update_id INTEGER PRIMARY KEY,
  user_id INTEGER NOT NULL,
  phone_number VARCHAR(20) NOT NULL,
  old_phone_number VARCHAR(20) NOT NULL,
  email VARCHAR(200) NOT NULL,
  old_email VARCHAR(200) NOT NULL,
  hashed_password VARCHAR(200) NOT NULL,
  old_hashed_password VARCHAR(200) NOT NULL,
  FOREIGN KEY (user_id) REFERENCES bshop.user(user_id) ON DELETE CASCADE
);

CREATE TABLE bshop.product_history (
  update_id INTEGER PRIMARY KEY,
  product_id INTEGER NOT NULL,
  price INTEGER NOT NULL,
  old_price INTEGER NOT NULL,
  rating NUMERIC NOT NULL,
  old_rating NUMERIC,
  FOREIGN KEY (product_id) REFERENCES bshop.product(product_id) ON DELETE CASCADE
);