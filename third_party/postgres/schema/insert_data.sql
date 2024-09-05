INSERT INTO "salesman" (id, shopname, email, phone) VALUES
(1, 'Shop A', 'shopa@example.com', '123-456-7890'),
(2, 'Shop B', 'shopb@example.com', '098-765-4321'),
(3, 'Shop C', 'shopc@example.com', '555-555-5555'),
(4, 'Shop D', 'shopd@example.com', '666-777-8888');

INSERT INTO "customer" (id, name, email, phone, address) VALUES
(1, 'Alice', 'alice@example.com', '111-222-3333', '123 Main St'),
(2, 'Bob', 'bob@example.com', '444-555-6666', '456 Oak Ave'),
(3, 'Charlie', 'charlie@example.com', '777-888-9999', '789 Pine Rd'),
(4, 'Diana', 'diana@example.com', '999-000-1111', '321 Maple St');

INSERT INTO "product" (id, salesmanid, price, type) VALUES
(1, 1, 100, 'Books'),
(2, 2, 150, 'Books'),
(3, 3, 200, 'Books'),
(4, 4, 50, 'Books'),
(5, 1, 75, 'Books');

INSERT INTO "book" (id, salesmanid, price, type, name, author, genre, year) VALUES
(6, 1, 50, 'Book', 'The Great Gatsby', 'F. Scott Fitzgerald', 'Fiction', 1925),
(7, 2, 60, 'Book', 'To Kill a Mockingbird', 'Harper Lee', 'Fiction', 1960),
(8, 3, 70, 'Book', '1984', 'George Orwell', 'Dystopian', 1949),
(9, 4, 80, 'Book', 'Brave New World', 'Aldous Huxley', 'Dystopian', 1932);

INSERT INTO "order" (id, date, status, customerid) VALUES
(1, '2024-01-01', 'Processing', 1),
(2, '2024-01-02', 'Shipped', 2),
(3, '2024-01-03', 'Delivered', 3),
(4, '2024-01-04', 'Processing', 4);

INSERT INTO "cart-product" (productid, cartid) VALUES
(1, 1),
(2, 1),
(3, 2),
(4, 3),
(5, 4);

INSERT INTO "order-product" (orderid, productid) VALUES
(1, 1),
(1, 2),
(2, 3),
(3, 4),
(4, 5);

INSERT INTO "favorites-product" (productid, favoritesid) VALUES
(1, 1),
(2, 1),
(3, 2),
(4, 3),
(5, 4);
