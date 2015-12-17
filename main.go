package main

import (
	_ "github.com/joho/godotenv/autoload"
	"log"
	"net/http"
)

func main() {

	router := NewRouter()

	log.Fatal(http.ListenAndServe(":8080", router))
}
