package handlers

import (
	"database/sql"
	"encoding/json"
	"net/http"

	"github.com/DenHax/LibreLib-back/internal/database"
)

var db *sql.DB = database.GetDB()

func GetAllBookByID(w http.ResponseWriter, r *http.Request) {
	bookList, err := database.GetBooksByID(db)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookList)
}
