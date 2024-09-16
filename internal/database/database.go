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

func GetBooks(db *sql.DB) ([]Book, error) {
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

func GetBooksByCustomerID(db *sql.DB, ID int) ([]Book, error) {
	fmt.Println(ID)
	query :=
		`SELECT 
		p.id AS product_id,
		p.salesmanid,
		p.price,
		p.type,
		b.name AS book_name,
		b.author,
		b.genre,
		b.year
	FROM 
		"cart-product" cp
	JOIN 
		"product" p ON cp.productid = p.id
	JOIN 
		"book" b ON p.id = b.id
	JOIN 
		"customer" c ON cp.cartid = c.id
	WHERE 
		c.id = $1
	AND 
		p.type = 'book';`

	rows, err := db.Query(query, ID)

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

func GetBooksByAuthor(db *sql.DB, author string) ([]Book, error) {
	query :=
		`SELECT
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
		b.author = $1;`

	rows, err := db.Query(query, author)

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

func GetBooksByGenre(db *sql.DB, genre string) ([]Book, error) {
	query :=
		`SELECT
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
		b.genre = $1;`

	rows, err := db.Query(query, genre)

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

func GetBooksByYear(db *sql.DB, year1, year2 int) ([]Book, error) {
	query :=
		`SELECT
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
		b.year >= $1 AND b.year <= $2;`

	rows, err := db.Query(query, year1, year2)

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

func GetBooksByPrice(db *sql.DB, price1, price2 float64) ([]Book, error) {
	query :=
		`SELECT
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
		p.price >= $1 AND p.price <= $2;`

	rows, err := db.Query(query, price1, price2)

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
