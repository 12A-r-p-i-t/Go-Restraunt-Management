package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/12A-r-p-i-t/restraunt-management/helper"
	"github.com/12A-r-p-i-t/restraunt-management/model"
	"github.com/gorilla/mux"
	"golang.org/x/crypto/bcrypt"
)

var User *model.User

func GetUsers(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	AllUsers, err := model.FindAllUsers()
	if err != nil {
		log.Fatal("Error in fetching all data from User table")
	}
	json.NewEncoder(w).Encode(&AllUsers)
}

func GetUser(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	userID := vars["id"]
	ID, err := strconv.Atoi(userID)
	if err != nil {
		log.Fatal("Error while converting from string to int in GetUser: ", err)
	}
	foundUser, err := model.FindUserByID(ID)
	if err != nil {
		log.Fatal("Unable to fetch User with given ID: ", err)
	}
	json.NewEncoder(w).Encode(foundUser)
}

func Login(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	val, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error in reading from request body during login")
		return
	}
	var user model.User
	err = json.Unmarshal(val, &user)
	if err != nil {
		log.Fatal("Error in unmarshalling data to model struct during login")
		return
	}
	foundUser, err := user.FindUser()
	if err != nil {
		log.Fatal("No such user exists in database with given emailID")
		return
	}
	err = bcrypt.CompareHashAndPassword([]byte(foundUser.Password), []byte(user.Password))
	if err != nil {
		log.Fatal("Password does not match: ", err)
		return
	}
	token, expirationTime, err := helper.GenerateToken(*foundUser)
	if err != nil {
		log.Fatal("Error in generating the token: ", err)
		return
	}
	http.SetCookie(w, &http.Cookie{
		Name:    "token",
		Value:   token,
		Expires: expirationTime,
	})
	json.NewEncoder(w).Encode(foundUser)
}

func SignUp(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	val, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error in reading from request body during signup")
		return
	}
	var user model.User
	err = json.Unmarshal(val, &user)
	if err != nil {
		log.Fatal("Error in Unmarshaling data to model struct during signup")
		return
	}
	password := []byte(user.Password)
	hashedPassword, err := bcrypt.GenerateFromPassword(password, bcrypt.DefaultCost)
	if err != nil {
		log.Fatal("Error in hashing the password during signup")
		return
	}
	user.Password = string(hashedPassword)
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
