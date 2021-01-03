package controller

import (
	"database/sql"
	"net/http"
	//. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

var db *sql.DB

type ResponseResult struct {
	Error  string `json:"error"`
	Result string `json:"result"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented"))
}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented"))
}
func AddLinkHandler(w http.ResponseWriter, r *http.Request) {

}
func CheckLinkHandler(w http.ResponseWriter, r *http.Request) {

}
