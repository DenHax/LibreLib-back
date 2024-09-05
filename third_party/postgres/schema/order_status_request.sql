SELECT "order".id AS order_id, "order".status, product.id AS product_id, product.type, product.price
FROM "order"
JOIN "order-product" ON "order".id = "order-product".orderid
JOIN product ON "order-product".productid = product.id;
