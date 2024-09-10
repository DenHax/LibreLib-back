package database

import (
	"database/sql"
	"fmt"
)

func Set_Book_data(rows *sql.Rows) []Book {
	books := []Book{}
	for rows.Next() {
		book := Book{}
		err := rows.Scan(&book.Name, &book.Author, &book.Genre, &book.Year, &book.Price, &book.Product.ID, &book.Product.SalesmanID, &book.Product.Type)
		if err != nil {
			fmt.Println(err)
			fmt.Println("reading error")
			continue
		}
		books = append(books, book)
	}
	return books
}