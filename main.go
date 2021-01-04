package main

import (
	"database/sql"
	"fmt"
	"log"
	"net/http"

	"github.com/Bronsun/RESTAPI---GO/controller"
	"github.com/gorilla/mux"
	_ "github.com/lib/pq"
	"gopkg.in/inconshreveable/log15.v2"
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
func connectDB() *sql.DB {
	var log = log15.New()
	db, err := sql.Open("postgres", "dbname=restapi host=localhost user=bronson password=test1234 sslmode = disable")
	err = db.Ping()
	if err != nil {
		log.Error("failed to open database", "err", err)
	}
	log.Info("Connected to db")
	return db

}

func main() {
	connectDB()
	routes()
}
