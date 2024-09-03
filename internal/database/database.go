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
	sendRequest(db, `TRUNCATE TABLE "Customer"`)
	sendRequest(db, `INSERT INTO "Customer" ("id", "name", "email", "phone", "address") 
		VALUES (1, 'John Doe', 'johndoe@example.com', '123-456-7890', '123 Elm St, Springfield, IL');`)
	sendRequest(db, `INSERT INTO "Customer" ("id", "name", "email", "phone", "address") 
		VALUES (2, 'Jane Smith', 'janesmith@example.com', '098-765-4321', '456 Oak St, Metropolis, IL');`)
	sendRequest(db, `INSERT INTO "Customer" ("id", "name", "email", "phone", "address") 
		VALUES (3, 'Alice Johnson', 'alice.johnson@example.com', '555-123-4567', '789 Maple St, Gotham, NY');`)
	sendRequest(db, `INSERT INTO "Customer" ("id", "name", "email", "phone", "address") 
		VALUES (4, 'Bob White', 'bob.white@example.com', '444-555-6666', '321 Pine St, Smallville, KS');`)

	
	rows := selectData(db, "SELECT id, name, email, phone, address FROM \"Customer\"")
	defer rows.Close()

	users := []Customer{}
	for rows.Next() {
		p := Customer{}
		err := rows.Scan(&p.Id, &p.Name, &p.Email, &p.Phone, &p.Address)
		if err != nil {
			fmt.Println(err)
			continue
		}
		users = append(users, p)
	}
	for _, p := range users {
		fmt.Println(p.Id, p.Name, p.Email, p.Phone, p.Address)
	}
}