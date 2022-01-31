package handlers

import (
	"encoding/json"
	"log"
	"net/http"
	"rest_api/models"
)

func initHeaders(w http.ResponseWriter) {
	w.Header().Set("Content-Type", "application/json")
}

func GetAllBooks(w http.ResponseWriter, req *http.Request) {
	initHeaders(w)
	log.Println("Get info about all books")

	books := models.GetAllBooks()
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(books)
}
