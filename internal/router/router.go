package router

import (
	"github.com/DenHax/LibreLib-back/internal/handlers"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	// booksRouter := router.PathPrefix("/books").Subrouter()
	router.HandleFunc("/books", handlers.GetBooksHandler).Methods("GET")
	router.HandleFunc("/cart", handlers.GetCustomerCartHandler).Methods("GET")
	return router
}
