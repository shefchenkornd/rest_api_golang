package utils

import "github.com/gorilla/mux"

func BuildBookResource(r *mux.Router, prefix string) {
	r.HandleFunc(prefix + "/{id}", handlers.GetBookById).Methods("GET")
	r.HandleFunc(prefix, handlers.CreateBook).Methods("POST")
	r.HandleFunc(prefix + "/{id}", handlers.UpdateBookById).Methods("PUT")
	r.HandleFunc(prefix + "/{id}", handlers.DeleteBookById).Methods("DELETE")
}

func BuildManyBooksResource(r *mux.Router, prefix string) {
	r.HandleFunc(prefix, handlers.GetAllBooks).Methods("GET")
}
