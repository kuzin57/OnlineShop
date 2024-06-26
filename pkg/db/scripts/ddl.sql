CREATE SCHEMA IF NOT EXISTS bshop;

CREATE TABLE IF NOT EXISTS bshop.user (
  user_id SERIAL UNIQUE,
  firstname VARCHAR(100) NOT NULL,
  surname VARCHAR(100) NOT NULL,
  phone_number VARCHAR(20) NOT NULL,
  email VARCHAR(200) UNIQUE NOT NULL,
  birthdate DATE NOT NULL,
  hashed_password VARCHAR(200) NOT NULL
);

CREATE TABLE IF NOT EXISTS bshop.product (
  product_id SERIAL UNIQUE,
  category VARCHAR(100) NOT NULL,
  name VARCHAR(100) NOT NULL,
  brand VARCHAR(50),
  price INTEGER NOT NULL,
  available BOOLEAN NOT NULL,
  rating NUMERIC NOT NULL,
  rating_amount BIGINT NOT NULL,
  image_path VARCHAR(600),
  CHECK (rating <= 10)
);

CREATE TABLE IF NOT EXISTS bshop.review (
  review_id SERIAL UNIQUE,
  user_id INTEGER NOT NULL,
  comment TEXT,
  product_id INTEGER NOT NULL,
  date DATE NOT NULL,
  mark INTEGER NOT NULL,
  FOREIGN KEY (user_id) REFERENCES bshop.user(user_id) ON DELETE CASCADE,
  FOREIGN KEY (product_id) REFERENCES bshop.product(product_id) ON DELETE CASCADE,
  CHECK (mark >= 1 and mark <= 10)
);

CREATE TABLE IF NOT EXISTS bshop.purchase (
  purchase_id SERIAL UNIQUE,
  client_id INTEGER NOT NULL,
  purchase_price INTEGER NOT NULL,
  date DATE NOT NULL,
  delivery_date DATE NOT NULL,
  city VARCHAR(100) NOT NULL,
  street VARCHAR(300) NOT NULL,
  house_number INTEGER NOT NULL,
  flat_number INTEGER NOT NULL
  -- FOREIGN KEY (client_id) REFERENCES bshop.user(user_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS bshop.purchase_product (
  pp_id SERIAL UNIQUE,
  purchase_id INTEGER NOT NULL,
  product_id INTEGER NOT NULL,
  amount INTEGER NOT NULL
  -- FOREIGN KEY (purchase_id) REFERENCES bshop.purchase(purchase_id) ON DELETE CASCADE,
  -- FOREIGN KEY (product_id) REFERENCES bshop.product(product_id) ON DELETE CASCADE
);

CREATE TABLE IF NOT EXISTS bshop.user_history (
  user_id INTEGER NOT NULL,
  firstname VARCHAR(100) NOT NULL,
  surname VARCHAR(100) NOT NULL,
  phone_number VARCHAR(20) NOT NULL,
  email VARCHAR(200) UNIQUE NOT NULL,
  birthdate DATE NOT NULL,
  hashed_password VARCHAR(200) NOT NULL,
  change_time TIMESTAMP
);

CREATE TABLE IF NOT EXISTS bshop.product_history (
  product_id INTEGER NOT NULL,
  category VARCHAR(100) NOT NULL,
  name VARCHAR(100) NOT NULL,
  brand VARCHAR(50),
  price INTEGER NOT NULL,
  available BOOLEAN NOT NULL,
  rating NUMERIC NOT NULL,
  rating_amount BIGINT NOT NULL,
  image_path VARCHAR(600),
  change_time TIMESTAMP
);

-- SELECT bshop.product.product_id, name, brand, price, amount
-- FROM bshop.purchase_product
-- JOIN bshop.product
-- ON bshop.product.product_id = bshop.purchase_product.product_id
-- WHERE purchase_id = 1;

-- SELECT * FROM bshop.purchase;
-- SELECT * FROM bshop.purchase_product;

CREATE TABLE IF NOT EXISTS bshop.product_char (
  product_id SERIAL UNIQUE,
  country VARCHAR(100),
  net_weight NUMERIC NOT NULL,
  kcal INTEGER NOT NULL,
  proteins NUMERIC NOT NULL,
  fats NUMERIC NOT NULL,
  carbohydrates NUMERIC NOT NULL,
  expire_date INTEGER NOT NULL
);
