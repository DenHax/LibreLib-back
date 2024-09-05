SELECT customer.name AS customer_name, COUNT("cart-product".productid) AS product_count
FROM customer
JOIN "cart-product" ON customer.id = "cart-product".cartid
GROUP BY customer.name;
