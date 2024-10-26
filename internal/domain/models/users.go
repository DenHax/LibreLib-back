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
