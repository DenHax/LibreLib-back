package router

import (
	"github.com/DenHax/LibreLib-back/internal/handlers"
	"github.com/gorilla/mux"
)

func GetRouter() *mux.Router {
	router := mux.NewRouter()
	booksRouter := router.PathPrefix("/books").Subrouter()
	booksRouter.HandleFunc("/author", handlers.GetNameHandler)

	// booksRouter.HandleFunc("", handlers.GetAllBooks).Methods("GET")
	// booksRouter.HandleFunc("/author", handlers.GetBooksByAuthor).Methods("GET")
	// router.HandleFunc("/books/genre", handlers.GetBooksByGenre).Methods("GET")
	// router.HandleFunc("/books/year", handlers.GetBooksByYear).Methods("GET")
	// router.HandleFunc("/cart", handlers.GetBooksByCustomerID).Methods("GET")
	return router
}
