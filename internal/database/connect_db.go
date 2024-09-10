package database

import (
	"database/sql"
	"fmt"
	"os"
)

var (
	USER     string
	PASSWORD string
	HOST     string
	DBNAME   string
	PORT     string
)

func set_env_var() {
	USER = os.Getenv("DB_USER")
	PASSWORD = os.Getenv("POSTGRES_PASSWORD")
	HOST = os.Getenv("POSTGRES_HOST")
	DBNAME = os.Getenv("DB_NAME")
	PORT = os.Getenv("POSTGRES_PORT")
}

func Connect() *sql.DB {
	set_env_var()
	fmt.Println(USER, PASSWORD, HOST, DBNAME, PORT)
	connStr := "user=postgres password=12345 dbname=database port=5438 sslmode=disable"
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		panic(err)
	}
	fmt.Println("Connection!")
	return db
}
