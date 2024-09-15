package database

import (
	"database/sql"
	"fmt"

	_ "github.com/lib/pq"
)

func GetDB() *sql.DB {
	connStr := "user=postgres password=pwd3052 dbname=postgres port=5438 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection!")
	return db
}

func GetBooksByID(db *sql.DB) ([]Book, error) {
	query := `
	SELECT 
		p.id AS product_id,
		p.salesmanid,
		p.price,
		p.type,
		b.name AS book_name,
		b.author,
		b.genre,
		b.year
	FROM 
		"product" p
	JOIN 
		"book" b ON p.id = b.id
	WHERE 
		p.type = 'book';`

	rows, err := db.Query(query)
	if err != nil {
		panic(err)
	}
	defer rows.Close()

	books := []Book{}

	for rows.Next() {
		p := Book{}
		err := rows.Scan(&p.ID, &p.SalesmanID, &p.Price, &p.Type, &p.Name, &p.Author, &p.Genre, &p.Year)
		if err != nil {
			fmt.Println(err)
			continue
		}
		books = append(books, p)
	}
	return books, nil
}
