**USER**  (таблица пользователей) 

| Название         | Описание           | Тип данных     | Ограничение   |
|------------------| -------------------| -------------- |---------------|
| `user_id`        | Идентификатор      | `INTEGER`      | `PRIMARY KEY` |
| `firstname`      | Имя                | `VARCHAR(100)` | `NOT NULL`    |
| `surname`        | Фамилия            | `VARCHAR(100)` | `NOT NULL`    |
| `phone_number`   | Номер телефона     | `VARCHAR(20)`  | `NOT NULL`    |
| `email`          | Почта              | `VARCHAR(200)` | `NOT NULL`    |
| `birthdate`      | Дата рождения      | `DATE`         | `NOT NULL`    |
| `hashed_password`| Хешированный пароль| `VARCHAR(200)` | `NOT NULL`    |


**PRODUCT**  (таблица с описанием товаров) 

| Название       | Описание             | Тип данных     | Ограничение   |
|----------------| ---------------------| -------------- |---------------|
| `product_id`   | Идентификатор        | `INTEGER`      | `PRIMARY KEY` |
| `category`     | Категория товара     | `VARCHAR(100)` | `NOT NULL`    |
| `name`         | Название товара      | `VARCHAR(100)` | `NOT NULL`    |
| `brand`        | Фирма/марка          | `VARCHAR(50)`  |               |
| `price`        | Цена за кг (в рублях)| `INTEGER`      | `NOT NULL`    |
| `available`    | Есть ли в наличии    | `BOOLEAN`      | `NOT NULL`    |
| `rating`       | Рейтинг товара       | `NUMERIC`      | `NOT NULL`    |


**REVIEW**   (таблица с отзывами о товарах)

| Название       | Описание                                       | Тип данных     | Ограничение           |
|----------------| -----------------------------------------------| -------------- |-----------------------|
| `review_id`    | Идентификатор                                  | `INTEGER`      | `PRIMARY KEY`         |
| `user_id`      | Идентификатор пользователя, писавшего отзыв    | `INTEGER`      | `FOREIGN KEY NOT NULL`|
| `comment`      | Отзыв                                          | `TEXT`         |                       |
| `product_id`   | Идентификатор продукта, о котором написан отзыв| `INTEGER`      | `FOREIGN KEY NOT NULL`|
| `date`         | Дата отзыва                                    | `DATE`         | `NOT NULL`            |
| `mark`         | Оценка продукта (от 1 до 10)                   | `INTEGER`      | `NOT NULL`            |


**PURCHASE**   (таблица с информацией о покупках)

| Название        | Описание                   | Тип данных     | Ограничение           |
|-----------------| ---------------------------| -------------- |-----------------------|
| `purchase_id`   | Идентификатор              | `INTEGER`      | `PRIMARY KEY`         |
| `client_id`     | Идентификатор покупателя   | `INTEGER`      | `FOREIGN KEY NOT NULL`|
| `purchase_price`| Цена покупки               | `INTEGER`      | `NOT NULL`            |
| `date`          | Дата отзыва                | `DATE`         | `NOT NULL`            |


**PURCHASE_PRODUCT**   (таблица с содержанием покупок)

| Название       | Описание                                           | Тип данных     | Ограничение           |
|----------------| ---------------------------------------------------| -------------- |-----------------------|
| `pp_id`        | Идентификатор                                      | `INTEGER`      | `PRIMARY KEY`         |
| `purchase_id`  | Идентификатор покупки                              | `INTEGER`      | `FOREIGN KEY NOT NULL`|
| `product_id`   | Идентификатор продукта                             | `INTEGER`      | `FOREIGN KEY NOT NULL`|
| `amount`       | Текстовое описание количества купленного продукта  | `TEXT`         | `NOT NULL`            |


**USER_HISTORY** (версионная таблица для пользователей)  

| Название             | Описание                        | Тип данных     | Ограничение           |
|----------------------| --------------------------------| -------------- |-----------------------|
| `update_id`          | Идентификатор                   | `INTEGER`      | `PRIMARY KEY`         |
| `user_id`            | Идентификатор объекта изменения | `INTEGER`      | `FOREIGN KEY NOT NULL`|
| `phone_number`       | Новый номер телефона            | `VARCHAR(20)`  | `NOT NULL`            |
| `old_phone_number`   | Старый номер телефона           | `VARCHAR(20)`  | `NOT NULL`            |
| `email`              | Новый почтовый адрес            | `VARCHAR(200)` | `NOT NULL`            |
| `old_email`          | Старый почтовый адрес           | `VARCHAR(200)` | `NOT NULL`            |
| `hashed_password`    | Новый хеш пароля                | `VARCHAR(200)` | `NOT NULL`            |
| `old_hashed_password`| Старый хеш пароля               | `VARCHAR(200)` | `NOT NULL`            |


**PRODUCT_HISTORY** (версионная таблица для продуктов) 

| Название          | Описание                        | Тип данных     | Ограничение           |
|-------------------| --------------------------------| -------------- |-----------------------|
| `update_id`       | Идентификатор                   | `INTEGER`      | `PRIMARY KEY`         |
| `product_id`      | Идентификатор объекта изменения | `INTEGER`      | `FOREIGN KEY NOT NULL`|
| `price`           | Новая цена                      | `INTEGER`      | `NOT NULL`            |
| `old_price`       | Старая цена                     | `INTEGER`      | `NOT NULL`            |
| `rating`          | Новый рейтинг                   | `NUMERIC`      | `NOT NULL`            |
| `old_rating`      | Старый рейтинг                  | `NUMERIC`      |                       |
