package router

import (
	"github.com/DenHax/LibreLib-back/internal/handlers"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	router.HandleFunc("/getBooks", handlers.GetAllBooks).Methods("GET")
	router.HandleFunc("/getBooks/{id}", handlers.GetBooksByCustomerID).Methods("GET")
	return router
}
