package main

import (
	"fmt"
	"log"
	"net/http"

	"github.com/Bronsun/RESTAPI---GO/controller"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
)

//go:generate sqlboiler postgres
func XD(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("I think it is pretty simple but sqlboiler is pretty hard XD"))
}

func routes() {
	r := mux.NewRouter()
	r.HandleFunc("/", XD).Methods("GET")
	r.HandleFunc("/register", controller.RegisterHandler).Methods("POST")
	r.HandleFunc("/login", controller.LoginHandler).Methods("POST")
	fmt.Println("Server is working on port :8888")
	log.Fatal(http.ListenAndServe(":8888", r))
}

func main() {
	database.connectDB()
	routes()
}
