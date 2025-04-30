package main

import (
	srv "crud/server"
	"log"
	"net/http"

	"github.com/gorilla/mux"
)

func main() {
	router := mux.NewRouter()
	router.HandleFunc("/usuario", srv.GetUsuario).Methods("GET")
	router.HandleFunc("/createUsuario", srv.CreateUsuario).Methods("POST")

	log.Fatal(http.ListenAndServe(":5000", router))
}
