package main

import (
	"log"
	"net/http"
	"superhero-search/db"
	"superhero-search/handlers"

	"github.com/gorilla/mux"
)

func main() {
	// Connect to MongoDB
	db.Connect()

	// Set up routes
	router := mux.NewRouter()
	router.HandleFunc("/search", handlers.SearchHandler).Methods("GET")

	// Start the server
	log.Println("Server is running on http://localhost:8082")
	log.Fatal(http.ListenAndServe(":8082", router))
}
