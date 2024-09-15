package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"
	"strconv"

	"github.com/DenHax/LibreLib-back/internal/database"
	"github.com/gorilla/mux"
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

func GetBooksByID(w http.ResponseWriter, r *http.Request) {
	vars := mux.Vars(r)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}
	book, err := database.GetBooksByID(db, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
