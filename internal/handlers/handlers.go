package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DenHax/LibreLib-back/internal/database"
)

var db *sql.DB = database.GetDB()

func GetAllBooks(w http.ResponseWriter, r *http.Request) {
	bookList, err := database.GetBooks(db)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookList)
}

func GetBooksByCustomerID(w http.ResponseWriter, r *http.Request) {
	id_str := r.URL.Query().Get("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}
	book, err := database.GetBooksByCustomerID(db, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
