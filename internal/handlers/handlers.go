package handlers

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"net/http"
	"strconv"
	"strings"

	"github.com/DenHax/LibreLib-back/internal/database"
)

var db *sql.DB = database.GetDB()

func GetBooksHandler(w http.ResponseWriter, r *http.Request) {
	queryParams := r.URL.Query()
	var conditions []string
	var args []interface{}
	argID := 1
	for key, values := range queryParams {
		if len(values) > 0 {
			value := values[0]
			switch key {
			case "id":
				conditions = append(conditions, fmt.Sprintf("c.id = $%d", argID))
				args = append(args, value)
				argID++
			case "author":
				conditions = append(conditions, fmt.Sprintf("b.author = $%d", argID))
				args = append(args, value)
				argID++
			case "genre":
				conditions = append(conditions, fmt.Sprintf("b.genre = $%d", argID))
				args = append(args, value)
				argID++
			case "years":
				f_year := r.URL.Query().Get("f_year")
				s_year := r.URL.Query().Get("s_year")
				if f_year != "" && s_year != "" {
					conditions = append(conditions, fmt.Sprintf("b.year BETWEEN $%d AND $%d", argID, argID+1))
					args = append(args, f_year, s_year)
					argID += 2
				}
			case "min_price":
				conditions = append(conditions, fmt.Sprintf("p.price >= $%d", argID))
				args = append(args, value)
				argID++
			
			case "max_price":
				conditions = append(conditions, fmt.Sprintf("p.price <= $%d", argID))
				args = append(args, value)
				argID++
			}
		}
	}

	var wherePart string
	if len(conditions) > 0 {
		wherePart = "WHERE " + strings.Join(conditions, " AND ")
	}
	bookList, err := database.GetBooks(db, wherePart, args)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(bookList)
}

func GetCustomerCartHandler(w http.ResponseWriter, r *http.Request) {
	id_str := r.URL.Query().Get("id")
	id, err := strconv.Atoi(id_str)
	if err != nil {
		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
		return
	}
	book, err := database.GetCustomerCart(db, id)
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
	}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}
