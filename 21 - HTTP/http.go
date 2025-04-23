package main

import (
	"log"
	"net/http"
)

func home(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the home page!"))
}

func users(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Welcome to the users page!"))
}

func main() {
	// Funcao para criar rotas
	http.HandleFunc("/home", home)

	http.HandleFunc("/users", users)

	log.Fatal(http.ListenAndServe(":5000", nil))
}
