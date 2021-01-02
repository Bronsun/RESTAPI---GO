package main

import (
	"log"
	"net/http"

	"github.com/Bronsun/controller"
	"github.com/gorilla/mux"
)

//go:generate sqlboiler postgres

func main() {
	r := mux.NewRouter()
	r.HandleFunc("/register", controller.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", controller.LoginHandler).Methods("POST")

	log.Fatal(http.ListenAndServe(":5050", r))
}
