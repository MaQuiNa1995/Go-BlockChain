package main

import (
	"MaQuina1995/blockchain/model"
	"log"
	"net/http"
	"os"

	"github.com/gorilla/handlers"
)

func main() {
	port := os.Args[1]
	router := model.Lottery.NewRouter(port)

	allowedOrigins := handlers.AllowedOrigins([]string{"*"})
	allowedMethods := handlers.AllowedMethods([]string{"GET", "POST"})

	rules := handlers.CORS(allowedOrigins, allowedMethods)

	// launch server
	log.Fatal(http.ListenAndServe(":"+port, rules, router))
}
