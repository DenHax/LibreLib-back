package handlers

import (
	"encoding/json"
	"net/http"
	"sync"
	// "github.com/DenHax/LibreLib-back/internal/database"
)

var (
	mutex sync.Mutex
)

func GetNameHandler(w http.ResponseWriter, r *http.Request) {
	if r.Method == http.MethodOptions {
		w.WriteHeader(http.StatusOK)
		return
	}

	mutex.Lock()
	defer mutex.Unlock()
	json.NewEncoder(w).Encode(map[string]string{"name": "hello"})
}

// var db *sql.DB = database.GetDB()

// func GetAllBooks(w http.ResponseWriter, r *http.Request) {
// 	bookList, err := database.GetBooks(db)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(bookList)
// }
//
// func GetBooksByCustomerID(w http.ResponseWriter, r *http.Request) {
// 	id_str := r.URL.Query().Get("id")
// 	id, err := strconv.Atoi(id_str)
// 	if err != nil {
// 		http.Error(w, "Invalid customer ID", http.StatusBadRequest)
// 		return
// 	}
// 	book, err := database.GetBooksByCustomerID(db, id)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(book)
// }
//
// func GetBooksByAuthor(w http.ResponseWriter, r *http.Request) {
// 	author := r.URL.Query().Get("author")
// 	book, err := database.GetBooksByAuthor(db, author)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(book)
// }
//
// func GetBooksByGenre(w http.ResponseWriter, r *http.Request) {
// 	genre := r.URL.Query().Get("genre")
// 	book, err := database.GetBooksByGenre(db, genre)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(book)
// }
//
// func GetBooksByYear(w http.ResponseWriter, r *http.Request) {
// 	f_year_str := r.URL.Query().Get("f_year")
// 	s_year_str := r.URL.Query().Get("s_year")
// 	f_year, err := strconv.Atoi(f_year_str)
// 	if err != nil {
// 		http.Error(w, "Invalid year", http.StatusBadRequest)
// 		return
// 	}
// 	s_year, err := strconv.Atoi(s_year_str)
// 	if err != nil {
// 		http.Error(w, "Invalid year", http.StatusBadRequest)
// 		return
// 	}
// 	book, err := database.GetBooksByYear(db, f_year, s_year)
// 	if err != nil {
// 		w.WriteHeader(http.StatusBadRequest)
// 	}
// 	w.WriteHeader(http.StatusOK)
// 	json.NewEncoder(w).Encode(book)
// }
