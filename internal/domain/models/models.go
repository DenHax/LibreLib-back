package models

type Salesman struct {
	ID       int    `json:"id"`
	ShopName string `json:"shopname"`
	Email    string `json:"email"`
	Phone    string `json:"phone"`
}

type Customer struct {
	ID      int    `json:"id"`
	Name    string `json:"name"`
	Email   string `json:"email"`
	Phone   string `json:"phone"`
	Address string `json:"address"`
}

type Product struct {
	ID         int    `json:"id"`
	SalesmanID int    `json:"salesmanid"`
	Price      int    `json:"price"`
	Type       string `json:"type"`
}

type Book struct {
	Name   string `json:"name"`
	Author string `json:"author"`
	Genre  string `json:"genre"`
	Year   int16  `json:"year"`
	Product
}

type Order struct {
	ID         int    `json:"id"`
	Date       string `json:"date"`
	Status     string `json:"status"`
	CustomerID int    `json:"customerid"`
}

type CartProduct struct {
	ProductID int `json:"productid"`
	CartID    int `json:"cartid"`
}

type OrderProduct struct {
	OrderID   int `json:"orderid"`
	ProductID int `json:"productid"`
}

type FavoritesProduct struct {
	ProductID   int `json:"productid"`
	FavoritesID int `json:"favoritesid"`
}
