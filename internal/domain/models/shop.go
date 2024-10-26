package models

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
