CREATE SCHEMA IF NOT EXISTS bshop;

CREATE TABLE bshop.user (
  user_id INTEGER PRIMARY KEY,
  firstname VARCHAR(100) NOT NULL,
  surname VARCHAR(100) NOT NULL,
  phone_number VARCHAR(20) NOT NULL,
  email VARCHAR(200) NOT NULL,
  birthdate DATE NOT NULL,
  hashed_password VARCHAR(200) NOT NULL
);

CREATE TABLE bshop.product (
  product_id INTEGER PRIMARY KEY,
  category VARCHAR(100) NOT NULL,
  name VARCHAR(100) NOT NULL,
  brand VARCHAR(50),
  price INTEGER NOT NULL,
  available BOOLEAN NOT NULL,
  rating NUMERIC NOT NULL,
  CHECK (rating <= 10)
);

CREATE TABLE bshop.review (
  review_id INTEGER PRIMARY KEY,
  user_id INTEGER NOT NULL,
  comment TEXT,
  product_id INTEGER NOT NULL,
  date DATE NOT NULL,
  mark INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES bshop.user(user_id) ON DELETE CASCADE,
  FOREIGN KEY (product_id) REFERENCES bshop.product(product_id) ON DELETE CASCADE,
  CHECK (mark >= 1 and mark <= 10)
);

CREATE TABLE bshop.purchase (
  purchase_id INTEGER PRIMARY KEY,
  client_id INTEGER NOT NULL,
  purchase_price INTEGER NOT NULL,
  date DATE NOT NULL,
  FOREIGN KEY (client_id) REFERENCES bshop.user(user_id) ON DELETE CASCADE
);

CREATE TABLE bshop.purchase_product (
  pp_id INTEGER PRIMARY KEY,
  purchase_id INTEGER NOT NULL,
  product_id INTEGER NOT NULL,
  amount TEXT NOT NULL,
  FOREIGN KEY (purchase_id) REFERENCES bshop.purchase(purchase_id) ON DELETE CASCADE,
  FOREIGN KEY (product_id) REFERENCES bshop.product(product_id) ON DELETE CASCADE
);

CREATE TABLE bshop.user_history (
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