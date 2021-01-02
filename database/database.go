package CRUD

import (
	"database/sql"

	"github.com/Bronsun/RESTAPI---GO/models"
	_ "github.com/lib/pq"
	"gopkg.in/inconshreveable/log15.v2"
)

var log = log15.New()

func CRUD() {
	db, err := sql.Open("postgres", "dbname=restapi host=localhost user=bronson password=test1234 sslmode = disable")
	if err != nil {
		log.Error("failed to open database", "err", err)
		return
	}
	log.Info("Connected to db")
	p := &models.User{
		Login:    "Bronsun",
		Password: "test1234",
	}
	if err := p.Insert(db); err != nil {
		log.Error("failed to insert user", "err", err)
		return
	}
	log.Info("Insert user into db", "id", p.ID)
}
