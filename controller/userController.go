package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"

	"github.com/12A-r-p-i-t/restraunt-management/helper"
	"github.com/12A-r-p-i-t/restraunt-management/model"
)

var User *model.User

func GetUsers(w http.ResponseWriter, r *http.Request) {

}

func GetUser(w http.ResponseWriter, r *http.Request) {

}

func Login(w http.ResponseWriter, r *http.Request) {

}

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	val, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error in reading from request body during signup")
	}
	var user model.User
	err = json.Unmarshal(val, &user)
	if err != nil {
		log.Fatal("Error in Unmarshaling data to model struct during signup")
	}
	isPresent := model.CheckEmail(user.Email)
	if isPresent {
		json.NewEncoder(w).Encode("Email already exist")
		return
	}
	isPresent = model.CheckPhone(user.Phone)
	if isPresent {
		json.NewEncoder(w).Encode("Phone Number already used")
		return
	}
	token, expirationTime, err := helper.GenerateToken(user)
	if err != nil {
		log.Fatal("Error in generating the token: ", err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})
	insertedUser := user.InsertUser()
	json.NewEncoder(w).Encode(insertedUser)
}
