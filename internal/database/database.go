package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func sendRequest(db *sql.DB, request string) {
	_, err := db.Exec(request)
	if err != nil {
		panic(err)
	}
	fmt.Println("request sent successfully")
}

func selectData(db *sql.DB, request string) *sql.Rows {
	rows, err := db.Query(request)
	if err != nil {
		panic(err)
	}
	fmt.Println("data get successfully")
	return rows
}

func Start() {
	connStr := "user=postgres password=12345 dbname=database port=5438 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	defer db.Close()

	fmt.Println("Connection!")

	rows := selectData(db, "SELECT id, name, email, phone, address FROM customer;")

	defer rows.Close()

	customers := []Customer{}
	for rows.Next() {
		customer := Customer{}
		err := rows.Scan(&customer.Id, &customer.Name, &customer.Email, &customer.Phone, &customer.Address)
		if err != nil {
			fmt.Println("reading error")
			continue
		}
		customers = append(customers, customer)
	}
	for _, customer := range customers {
		fmt.Println(customer.Id, customer.Name, customer.Email, customer.Phone, customer.Address)
	}

}
