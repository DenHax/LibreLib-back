CREATE TABLE "salesman" (
	"id" INTEGER NOT NULL UNIQUE PRIMARY KEY,
	"shopname" VARCHAR,
	"email" VARCHAR,
	"phone" VARCHAR
);

CREATE TABLE "customer" (
	"id" INTEGER NOT NULL UNIQUE PRIMARY KEY,
	"name" VARCHAR,
	"email" VARCHAR,
	"phone" VARCHAR,
	"address" VARCHAR
);

CREATE TABLE "product" (
	"id" INTEGER NOT NULL UNIQUE PRIMARY KEY,
	"salesmanid" INTEGER,
	"price" INTEGER,
	"type" VARCHAR,
	CONSTRAINT fk_productSalesmanID 
		FOREIGN KEY("salesmanid")
		REFERENCES "salesman"("id")
		ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE "book" (
	"name" VARCHAR,
	"author" VARCHAR,
	"genre" VARCHAR,
	"year" SMALLINT
) INHERITS("product");



CREATE TABLE "order" (
	"id" INTEGER NOT NULL UNIQUE PRIMARY KEY,
	"date" DATE,
	"status" VARCHAR,
	"customerid" INTEGER,
	CONSTRAINT fk_Order_CustomerID
		FOREIGN KEY("customerid")
		REFERENCES "customer"("id")
		ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE "cart-product" (
	"productid" INTEGER,
	"cartid" INTEGER,
	CONSTRAINT fk_cart__Product_cartID
		FOREIGN KEY("cartid")
		REFERENCES "customer"("id")
		ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT fk_cart__Product_productID 
		FOREIGN KEY("productid")
		REFERENCES "product"("id")
		ON UPDATE NO ACTION ON DELETE NO ACTION
);

CREATE TABLE "order-product" (
	"orderid" INTEGER,
	"productid" INTEGER,
	CONSTRAINT fk_Product_ID 
		FOREIGN KEY("productid")
		REFERENCES "product"("id")
		ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT fk_Order_ID
		FOREIGN KEY("orderid")
		REFERENCES "order"("id")
		ON UPDATE NO ACTION ON DELETE NO ACTION
	
);

CREATE TABLE "favorites-product" (
	"productid" INTEGER,
	"favoritesid" INTEGER,
	CONSTRAINT fk_favorites__Product_favoritesID
		FOREIGN KEY("favoritesid")
		REFERENCES "customer"("id")
		ON UPDATE NO ACTION ON DELETE NO ACTION,
	CONSTRAINT fk_favorites__Product_productID
		FOREIGN KEY("productid")
		REFERENCES "product"("id")
		ON UPDATE NO ACTION ON DELETE NO ACTION
);
