package api

import (
	"github.com/DenHax/LibreLib-back/internal/middleware"
	"github.com/gorilla/mux"
)

func (h Handler) NewApi(r *mux.Router) *mux.Router {
	api := r.PathPrefix("/api/").Subrouter()
	api.Use(middleware.HeaderMiddleware)
	v1 := api.PathPrefix("/v1/").Subrouter()

	booksRouter := v1.PathPrefix("/books").Subrouter()
	// booksRouter.HandleFunc("/", handlers.GetAllBooks(s)).Methods("GET")
	//
	// booksRouter.HandleFunc("/author", handlers.GetNameHandler).Methods("GET")

	// booksRouter.HandleFunc("", handlers.GetAllBooks).Methods("GET")
	// booksRouter.HandleFunc("/author", handlers.GetBooksByAuthor).Methods("GET")
	// router.HandleFunc("/books/genre", handlers.GetBooksByGenre).Methods("GET")
	// router.HandleFunc("/books/year", handlers.GetBooksByYear).Methods("GET")
	// router.HandleFunc("/cart", handlers.GetBooksByCustomerID).Methods("GET")
	return api
}
