package controller

import (
	"github.com/Bronsun/RESTAPI---GO/models"
	"context"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"log"
	"net/http"
	"github.com/Bronsun/database"
	"github.com/Bronsun/models"
	jwt "github.com/dgrijalva/jwt-go"
	"golang.org/x/crypto/bcrypt"
)

boil.SetDB(db)

func RegisterHandler(w http.ResponseWriter, r *http.Request){
	w.Header().Set("Content-Type", "application/json")
	var user models.User

	body,_:=ioutil.ReadAll(r.Body)
	err:=json.Unmarshal(body, &user)
	var res models.ResponseResult


}