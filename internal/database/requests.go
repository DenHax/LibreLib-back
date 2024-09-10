package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func Select_data(db *sql.DB, request string) *sql.Rows {
	rows, err := db.Query(request)
	if err != nil {
		panic(err)
	}
	fmt.Println("data get successfully")
	return rows
}

func Send_request(db *sql.DB, request string) {
	_, err := db.Exec(request)
	if err != nil {
		panic(err)
	}
	fmt.Println("request sent successfully")
}
