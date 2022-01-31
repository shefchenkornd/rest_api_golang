package main

import (
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
	"log"
	"net/http"
	"os"
	"rest_api_pizza/utils"
)

const apiPrefix = "/api/v1"

var (
	port                    string
	bookResourcePrefix      = apiPrefix + "/book"  // /api/v1/book
	manyBooksResourcePrefix = apiPrefix + "/books" // /api/v1/books
)

func init() {
	log.Println("Starting init function...")

	err := godotenv.Load()
	if err != nil {
		log.Fatal("couldn't find .env file: ", err)
	}

	port = os.Getenv("app_port")

	log.Println("init function completed successfully!")
}

func main() {
	log.Println("Starting REST API on port:", port)

	router := mux.NewRouter()
	utils.BuildBookResource(router, bookResourcePrefix)
	utils.BuildManyBooksResource(router, manyBooksResourcePrefix)

	log.Println("Router initializing successfully. Ready to go!")
	log.Fatalln(http.ListenAndServe(":"+port, router))
}
