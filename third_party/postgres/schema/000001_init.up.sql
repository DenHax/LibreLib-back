CREATE TABLE "Customer" (
	"id" INTEGER NOT NULL UNIQUE,
	"name" VARCHAR,
	"email" VARCHAR,
	"phone" VARCHAR,
	"address" VARCHAR,
	PRIMARY KEY("id")
);


-- CREATE TABLE "Book" (
-- 	"id" INTEGER NOT NULL UNIQUE,
-- 	"name" VARCHAR,
-- 	"author" VARCHAR,
-- 	"genre" VARCHAR,
-- 	"yaer" SMALLINT,
-- 	"type" VARCHAR,
-- 	PRIMARY KEY("id")
-- );


CREATE TABLE "Salesman" (
	"id" INTEGER NOT NULL UNIQUE,
	"shopname" VARCHAR,
	"email" VARCHAR,
	"phone" VARCHAR,
	PRIMARY KEY("id")
);


CREATE TABLE "Product" (
	"id" INTEGER NOT NULL UNIQUE,
	"salesmanID" INTEGER,
	"price" INTEGER,
	"type" VARCHAR,
	PRIMARY KEY("id")
);


CREATE TABLE "Order" (
	"id" INTEGER NOT NULL UNIQUE,
	"date" DATE,
	"status" VARCHAR,
	"customerID" INTEGER,
	PRIMARY KEY("id")
);


CREATE TABLE "Cart-Product" (
	"productID" INTEGER,
	"cartID" INTEGER
);


CREATE TABLE "Order-Product" (
	"orderID" INTEGER,
	"poductID" INTEGER
);


CREATE TABLE "Favorites-Product" (
	"productID" INTEGER,
	"favoritesID" INTEGER
);

-- ON UPDATE NO ACTION ON DELETE NO ACTION;
-- ALTER TABLE "Product"
-- ADD FOREIGN KEY("salesmanID") REFERENCES "Salesman"("id")
-- ON UPDATE NO ACTION ON DELETE NO ACTION;
-- ALTER TABLE "Order-Product"
-- ADD FOREIGN KEY("poductID") REFERENCES "Product"("id")
-- ON UPDATE NO ACTION ON DELETE NO ACTION;
-- ALTER TABLE "Order-Product"
-- ADD FOREIGN KEY("orderID") REFERENCES "Order"("id")
-- ON UPDATE NO ACTION ON DELETE NO ACTION;
-- ALTER TABLE "Order"
-- ADD FOREIGN KEY("customerID") REFERENCES "Customer"("id")
-- ON UPDATE NO ACTION ON DELETE NO ACTION;
-- ALTER TABLE "Favorites-Product"
-- ADD FOREIGN KEY("favoritesID") REFERENCES "Customer"("id")
-- ON UPDATE NO ACTION ON DELETE NO ACTION;
-- ALTER TABLE "Favorites-Product"
-- ADD FOREIGN KEY("productID") REFERENCES "Product"("id")
-- ON UPDATE NO ACTION ON DELETE NO ACTION;
-- ALTER TABLE "Cart-Product"
-- ADD FOREIGN KEY("cartID") REFERENCES "Customer"("id")
-- ON UPDATE NO ACTION ON DELETE NO ACTION;
-- ALTER TABLE "Cart-Product"
-- ADD FOREIGN KEY("productID") REFERENCES "Product"("id")
-- ON UPDATE NO ACTION ON DELETE NO ACTION;
