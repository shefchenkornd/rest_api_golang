package handlers

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"rest_api/models"
	"strconv"
)

func GetBookById(w http.ResponseWriter, req *http.Request) {
	initHeaders(w)

	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := models.Message{Message: "Invalid Id"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	log.Println("Get info about book by id", id)
	book, found := models.FindBookById(id)
	if !found {
		msg := models.Message{Message: "Book not found"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func CreateBook(w http.ResponseWriter, req *http.Request) {
	initHeaders(w)
	log.Println("Creating new book")

	var book models.Book

	err := json.NewDecoder(req.Body).Decode(&book)
	if err != nil {
		msg := models.Message{Message: "Invalid data"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	newBookId := len(models.GetAllBooks()) + 1
	book.Id = newBookId
	models.CreateBook(book)

	w.WriteHeader(http.StatusCreated)
	json.NewEncoder(w).Encode(book)
}

func UpdateBookById(w http.ResponseWriter, req *http.Request) {
	initHeaders(w)
	log.Println("Updating book with ID")

	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := models.Message{Message: "Invalid Id"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	oldBook, found := models.FindBookById(id)
	if !found {
		msg := models.Message{Message: "Book not found"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(msg)
		return
	}

	var newBook models.Book
	err = json.NewDecoder(req.Body).Decode(&newBook)
	if err != nil {
		msg := models.Message{Message: "Invalid data"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	book, err := models.UpdateBook(newBook, oldBook)
	if err != nil {
		msg := models.Message{Message: "Invalid data"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(book)
}

func DeleteBookById(w http.ResponseWriter, req *http.Request) {
	initHeaders(w)
	log.Println("Deleting book with ID")

	vars := mux.Vars(req)
	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		msg := models.Message{Message: "Invalid Id"}
		w.WriteHeader(http.StatusBadRequest)
		json.NewEncoder(w).Encode(msg)
		return
	}

	book, found := models.FindBookById(id)
	if !found {
		msg := models.Message{Message: "Book not found"}
		w.WriteHeader(http.StatusNotFound)
		json.NewEncoder(w).Encode(msg)
		return
	}

	models.DeleteBook(book)

	msg := models.Message{Message: "Successfully deleted"}
	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(msg)
}
