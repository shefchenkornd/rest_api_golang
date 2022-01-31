package main

import (
	"encoding/json"
	"github.com/gorilla/mux"
	"log"
	"net/http"
	"strconv"
)

const port = "8080"

func init() {
	log.Println("Starting init function...")

	log.Println("init function completed successfully!")
}

func FindPizzaById(id int) (Pizza, bool) {
	var pizza Pizza
	var found bool

	for _, p := range db {
		if p.ID == id {
			pizza = p
			found = true
		}
	}

	return pizza, found
}

type ErrorMessage struct {
	Message string `json:"message"`
}

func GetAllPizzas(w http.ResponseWriter, req *http.Request) {
	log.Println("Get infos about all pizzas in database")

	w.Header().Set("Content-Type", "application/json")
	w.WriteHeader(http.StatusOK)
	err := json.NewEncoder(w).Encode(db) // сериализация + запись во http.ResponseWriter
	if err != nil {
		log.Println(err)
	}
}

func GetPizzaById(w http.ResponseWriter, req *http.Request) {
	log.Println("Get infos about pizza by ID in database")
	w.Header().Set("Content-Type", "application/json")

	vars := mux.Vars(req) // {"id": "12"}

	id, err := strconv.Atoi(vars["id"])
	if err != nil {
		log.Printf("Can't convert %v to int\n", vars["id"])

		w.WriteHeader(http.StatusBadRequest)
		msg := ErrorMessage{Message: "Invalid pizza ID"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	p, found := FindPizzaById(id)
	if !found {
		w.WriteHeader(http.StatusNotFound)
		msg := ErrorMessage{Message: "pizza with that ID doesn't exists in DB"}
		json.NewEncoder(w).Encode(msg)
		return
	}

	w.WriteHeader(http.StatusOK)
	json.NewEncoder(w).Encode(p)
}

func main() {
	log.Println("Trying to start REST API pizza!")

	// Инициализируем маршрутизатор
	router := mux.NewRouter()
	router.HandleFunc("/pizzas", GetAllPizzas).Methods("GET")
	router.HandleFunc("/pizza/{id}", GetPizzaById).Methods("GET")


	log.Fatalln(http.ListenAndServe(":"+port, router))
}