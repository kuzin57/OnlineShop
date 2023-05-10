INSERT INTO bshop.user (firstname, surname, phone_number, email, birthdate, hashed_password) 
VALUES 
('John', 'Doe', '9234555890', 'johndoe@gmail.com', '1990-01-01', 'sMkLxJyKwRnFpE'),
('Jane', 'Doe', '7983654821', 'janedoe@gmail.com', '1992-05-15', 'qZzGtHcXvIbNfD'),
('Bob', 'Smith', '5654274869', 'bobsmith@gmail.com', '1985-11-30', 'rTjUoPwQkLmYnB'),
('Alice', 'Johnson', '5553575123', 'alicejohnson@gmail.com', '1998-03-22', 'fGhJkLpOzXcVbN'),
('David', 'Lee', '7975651801', 'davidlee@gmail.com', '1982-09-10', 'aSdFgHjKlMnBvC'),
('Sarah', 'Kim', '1415278080', 'sarahkim@gmail.com', '1995-06-18', 'eRtYhUjIiOpLkN'),
('Michael', 'Brown', '7094541363', 'michaelbrown@gmail.com', '1978-12-25', 'dFgHjKlMnBvCxZ'),
('Emily', 'Davis', '4518320073', 'emilydavis@gmail.com', '1991-08-07', 'qWeRtYuIoPzXcVb'),
('James', 'Wilson', '1528384959', 'jameswilson@gmail.com', '1980-04-12', 'bNmJkLoPqRsTfGh'),
('Jessica', 'Taylor', '9581637854', 'jessicataylor@gmail.com', '1993-02-28', 'yHjKlMnBvCxZqW');

INSERT INTO bshop.product (category, name, brand, price, available, rating, rating_amount) VALUES
('Молочные продукты', 'Молоко', 'Простоквашино', 92, true, 8.7, 1),
('Молочные продукты', 'Творог', 'Домик в деревне', 55, true, 8.0, 1),
('Бакалея', 'Макароны', 'Barilla', 120, true, 8.5, 1),
('Бакалея', 'Рис', 'Мистраль', 166, true, 8.3, 1),
('Фрукты и овощи', 'Яблоки', 'Голден', 170, true, 8.6, 1),
('Фрукты и овощи', 'Огурцы', 'Fresh', 190, true, 7.3, 1),
('Кондитерские изделия', 'Шоколад', 'Alpen Gold', 96, true, 9.5, 1),
('Кондитерские изделия', 'Печенье', 'McVities', 93, false, 8.4, 1),
('Напитки', 'Чай', 'Lipton', 70, true, 8.8, 1),
('Напитки', 'Кофе', 'Nescafe', 650, true, 7.5, 1);

INSERT INTO bshop.review (user_id, comment, product_id, date, mark) 
VALUES 
(1, 'Качественный продукт', 5, '2022-01-01', 9),
(3, NULL, 1, '2022-01-02', 8),
(7, NULL, 2, '2022-01-03', 8),
(5, NULL, 6, '2022-01-04', 6),
(4, NULL, 1, '2022-01-05', 9),
(2, 'Не стоит своих денег', 6, '2022-01-06', 5),
(6, NULL, 8, '2022-01-07', 7),
(8, NULL, 3, '2022-01-08', 8),
(9, NULL, 7, '2022-01-09', 10),
(10, NULL, 10, '2022-01-10', 6);

INSERT INTO bshop.purchase (client_id, purchase_price, date)
VALUES 
(3, 1500, '2021-10-01'),
(5, 700, '2014-11-13'),
(2, 550, '2018-10-23'),
(1, 2000, '2020-07-06'),
(4, 1000, '2022-08-25'),
(6, 650, '2019-12-26'),
(8, 500, '2015-04-19'),
(7, 800, '2016-01-13'),
(9, 1300, '2019-03-18'),
(10, 950, '2017-05-10');

INSERT INTO bshop.purchase_product (purchase_id, product_id, amount) VALUES
(1, 2, '2 шт'),
(1, 3, '2 пачки'),
(1, 4, '1 пачка'),
(2, 1, '3 шт'),
(2, 5, '2 кг'),
(3, 7, '1 шт'),
(3, 8, '4 пачки'),
(4, 9, '3 шт'),
(4, 6, '2 пачки'),
(5, 10, '1 шт');