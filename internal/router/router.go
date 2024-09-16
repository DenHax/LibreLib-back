package router

import (
	"github.com/DenHax/LibreLib-back/internal/handlers"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/books", handlers.GetAllBooks).Methods("GET")
	router.HandleFunc("/cart", handlers.GetBooksByCustomerID).Methods("GET")
	return router
}
