package database

import (
	"database/sql"

	_ "github.com/lib/pq"
	"gopkg.in/inconshreveable/log15.v2"
)

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
