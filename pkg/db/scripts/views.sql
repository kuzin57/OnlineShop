CREATE OR REPLACE VIEW bshop.user_view AS
SELECT
REPLACE(firstname, substring(firstname, 2, length(firstname) - 2), repeat('*', 3)) as hidden_firstname,
REPLACE(surname, substring(surname, 2, length(surname) - 2), repeat('*', 3)) as hidden_surname,
REPLACE(cast(birthdate as TEXT), substring(cast(birthdate as TEXT), 2, length(cast(birthdate as TEXT)) - 2), repeat('*', 3)) as hidden_birthdate
FROM bshop.user;

CREATE OR REPLACE VIEW bshop.user_purchase_view AS
SELECT
bshop.user.firstname, bshop.user.surname,
REPLACE(cast(bshop.purchase.purchase_price as TEXT), substring(cast(bshop.purchase.purchase_price as TEXT), 1, length(cast(bshop.purchase.purchase_price as TEXT)) - 1), repeat('*', 3)) as purchase_cost
FROM bshop.purchase
JOIN bshop.user ON bshop.user.user_id = bshop.purchase.client_id;

CREATE OR REPLACE VIEW bshop.user_product_view AS
SELECT
bshop.user.firstname, bshop.user.surname, bshop.product.name, bshop.product.brand FROM
(select bshop.purchase_product.product_id, bshop.purchase.client_id as user_id
from bshop.purchase_product
join bshop.purchase on bshop.purchase.purchase_id = bshop.purchase_product.purchase_id) as tmp
JOIN bshop.user ON tmp.user_id = bshop.user.user_id
JOIN bshop.product ON bshop.product.product_id = tmp.product_id;

CREATE OR REPLACE VIEW bshop.review_info_view AS
SELECT
bshop.user.firstname, bshop.user.surname, bshop.product.name, bshop.product.brand, bshop.review.mark
FROM bshop.review
JOIN bshop.user ON bshop.user.user_id = bshop.review.user_id
JOIN bshop.product ON bshop.product.product_id = bshop.review.product_id;