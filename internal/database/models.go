package database

type Cart_Product struct {
	ProductID int
	CartID    int
}

type Customer struct {
	Id      int
	Name    string
	Email   string
	Phone   string
	Address string
}

type Salesman struct {
	Id       int
	Shopname string
	Email    string
	Phone    string
}

type Product struct {
	Id          int
	SalesmanID  int
	Price       int
	ProductType string
}

type Order struct {
	Id         int
	Date       string
	Status     string
	CustomerID int
}

type Order_Product struct {
	OrderID   int
	ProductID int
}

type Favorites_Product struct {
	ProductID   int
	FavoritesID int
}
