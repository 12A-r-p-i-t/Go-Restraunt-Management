package main

import (
	"log"
	"net/http"
	"os"

	"github.com/12A-r-p-i-t/restraunt-management/controller"
	"github.com/12A-r-p-i-t/restraunt-management/database"
	"github.com/gorilla/mux"
	"github.com/joho/godotenv"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatal("Error in loading the .env file")
	}

	// Database Connection
	database.Connect()
	// Starting New router
	r := mux.NewRouter()

	// User routes (login, Signup, Get All Users, Get User By ID)
	r.HandleFunc("/users", controller.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controller.GetUser).Methods("GET")
	r.HandleFunc("/users/login", controller.Login).Methods("POST")
	r.HandleFunc("/users/signup", controller.SignUp).Methods("POST")
	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(port, r))
}
