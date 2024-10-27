package api

import (
	"github.com/gorilla/mux"
)

func (h Handler) NewApi(r *mux.Router) *mux.Router {
	booksRouter := r.PathPrefix("/books").Subrouter()
	booksRouter.HandleFunc("/", GetNameHandler).Methods("GET")
	//
	// booksRouter.HandleFunc("/author", handlers.GetNameHandler).Methods("GET")

	// booksRouter.HandleFunc("", handlers.GetAllBooks).Methods("GET")
	// booksRouter.HandleFunc("/author", handlers.GetBooksByAuthor).Methods("GET")
	// router.HandleFunc("/books/genre", handlers.GetBooksByGenre).Methods("GET")
	// router.HandleFunc("/books/year", handlers.GetBooksByYear).Methods("GET")
	// router.HandleFunc("/cart", handlers.GetBooksByCustomerID).Methods("GET")
	return r
}
