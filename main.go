package main

import (
	"log"
	"net/http"
)

const port = "8080"

func main() {
	log.Println("Trying to start REST API pizza!")


	// Инициализируем маршрутизатор
	// router := mux.NewRouter()


	log.Fatalln(http.ListenAndServe(":"+port, nil))
}
