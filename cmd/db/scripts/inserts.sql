INSERT INTO bshop.user (user_id, firstname, surname, phone_number, email, birthdate, hashed_password) 
VALUES 
(1, 'John', 'Doe', '9234555890', 'johndoe@gmail.com', '1990-01-01', 'sMkLxJyKwRnFpE'),
(2, 'Jane', 'Doe', '7983654821', 'janedoe@gmail.com', '1992-05-15', 'qZzGtHcXvIbNfD'),
(3, 'Bob', 'Smith', '5654274869', 'bobsmith@gmail.com', '1985-11-30', 'rTjUoPwQkLmYnB'),
(4, 'Alice', 'Johnson', '5553575123', 'alicejohnson@gmail.com', '1998-03-22', 'fGhJkLpOzXcVbN'),
(5, 'David', 'Lee', '7975651801', 'davidlee@gmail.com', '1982-09-10', 'aSdFgHjKlMnBvC'),
(6, 'Sarah', 'Kim', '1415278080', 'sarahkim@gmail.com', '1995-06-18', 'eRtYhUjIiOpLkN'),
(7, 'Michael', 'Brown', '7094541363', 'michaelbrown@gmail.com', '1978-12-25', 'dFgHjKlMnBvCxZ'),
(8, 'Emily', 'Davis', '4518320073', 'emilydavis@gmail.com', '1991-08-07', 'qWeRtYuIoPzXcVb'),
(9, 'James', 'Wilson', '1528384959', 'jameswilson@gmail.com', '1980-04-12', 'bNmJkLoPqRsTfGh'),
(10, 'Jessica', 'Taylor', '9581637854', 'jessicataylor@gmail.com', '1993-02-28', 'yHjKlMnBvCxZqW');

INSERT INTO bshop.product (product_id, category, name, brand, price, available, rating) VALUES
(1, 'Молочные продукты', 'Молоко', 'Простоквашино', 92, true, 8.7),
(2, 'Молочные продукты', 'Творог', 'Домик в деревне', 55, true, 8.0),
(3, 'Бакалея', 'Макароны', 'Barilla', 120, true, 8.5),
(4, 'Бакалея', 'Рис', 'Мистраль', 166, true, 8.3),
(5, 'Фрукты и овощи', 'Яблоки', 'Голден', 170, true, 8.6),
(6, 'Фрукты и овощи', 'Огурцы', 'Fresh', 190, true, 7.3),
(7, 'Кондитерские изделия', 'Шоколад', 'Alpen Gold', 96, true, 9.5),
(8, 'Кондитерские изделия', 'Печенье', 'McVities', 93, false, 8.4),
(9, 'Напитки', 'Чай', 'Lipton', 70, true, 8.8),
(10, 'Напитки', 'Кофе', 'Nescafe', 650, true, 7.5);

INSERT INTO bshop.review (review_id, user_id, comment, product_id, date, mark) 
VALUES 
(1, 1, 'Качественный продукт', 5, '2022-01-01', 9),
(2, 3, NULL, 1, '2022-01-02', 8),
(3, 7, NULL, 2, '2022-01-03', 8),
(4, 5, NULL, 6, '2022-01-04', 6),
(5, 4, NULL, 1, '2022-01-05', 9),
(6, 2, 'Не стоит своих денег', 6, '2022-01-06', 5),
(7, 6, NULL, 8, '2022-01-07', 7),
(8, 8, NULL, 3, '2022-01-08', 8),
(9, 9, NULL, 7, '2022-01-09', 10),
(10, 10, NULL, 10, '2022-01-10', 6);

INSERT INTO bshop.purchase (purchase_id, client_id, purchase_price, date)
VALUES 
(1, 3, 1500, '2021-10-01'),
(2, 5, 700, '2014-11-13'),
(3, 2, 550, '2018-10-23'),
(4, 1, 2000, '2020-07-06'),
(5, 4, 1000, '2022-08-25'),
(6, 6, 650, '2019-12-26'),
(7, 8, 500, '2015-04-19'),
(8, 7, 800, '2016-01-13'),
(9, 9, 1300, '2019-03-18'),
(10, 10, 950, '2017-05-10');

INSERT INTO bshop.purchase_product (pp_id, purchase_id, product_id, amount) VALUES
(1, 1, 2, '2 шт'),
(2, 1, 3, '2 пачки'),
(3, 1, 4, '1 пачка'),
(4, 2, 1, '3 шт'),
(5, 2, 5, '2 кг'),
(6, 3, 7, '1 шт'),
(7, 3, 8, '4 пачки'),
(8, 4, 9, '3 шт'),
(9, 4, 6, '2 пачки'),
(10, 5, 10, '1 шт')
