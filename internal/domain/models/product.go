package models

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
