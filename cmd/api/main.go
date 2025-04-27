package main

import (
	"encoding/json"
	"gogroq/internal/handler"
	"log"
	"net/http"

	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error loading .env file")
	}

	router := mux.NewRouter()

	router.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		json.NewEncoder(w).Encode("Running Groq API")
	}).Methods("GET")

	router.HandleFunc("/api/chat", handler.ChatPost).Methods("POST")

	log.Println("Server is running on port 9000")
	http.ListenAndServe(":9000", router)
}
