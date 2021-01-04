package controller

import (
	"encoding/json"
	"io/ioutil"
	"net/http"

	"github.com/Bronsun/RESTAPI---GO/models"
	//. "github.com/volatiletech/sqlboiler/v4/queries/qm"
)

type ResponseResult struct {
	Error string `json:"error"`
}

func RegisterHandler(w http.ResponseWriter, r *http.Request) {
	body, err := ioutil.ReadAll(r.Body)
	if err != nil {
		resp := ResponseResult{Error: "Nothing Posted"}
		json.NewEncoder(w).Encode(resp)
	}
	var formattedBody models.User
	err = json.Unmarshal(body, &formattedBody)
	db := database.connectDB()

}

func LoginHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("Not implemented"))
}
func AddLinkHandler(w http.ResponseWriter, r *http.Request) {

}
func CheckLinkHandler(w http.ResponseWriter, r *http.Request) {

}
