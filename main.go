package main

import (
	"log"
	"net/http"
	"os"

	"github.com/12A-r-p-i-t/restraunt-management/controller"
	"github.com/12A-r-p-i-t/restraunt-management/database"
	"github.com/12A-r-p-i-t/restraunt-management/middleware"
	"github.com/12A-r-p-i-t/restraunt-management/model"
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
	db := database.GetDB()
	db.AutoMigrate(&model.User{})
	db.AutoMigrate(&model.Food{})
	db.AutoMigrate(&model.Menu{})
	db.AutoMigrate(&model.Table{})

	// Starting New router
	r := mux.NewRouter()

	// Public User routes (login, Signup, Get All Users, Get User By ID)
	r.HandleFunc("/users", controller.GetUsers).Methods("GET")
	r.HandleFunc("/users/{id}", controller.GetUser).Methods("GET")
	r.HandleFunc("/users/login", controller.Login).Methods("POST")
	r.HandleFunc("/users/signup", controller.SignUp).Methods("POST")

	//Adding middleware to route
	api := r.PathPrefix("/api").Subrouter()
	api.Use(middleware.Authorize)

	//Private Routes Related to foods
	api.HandleFunc("/foods", controller.GetFoods).Methods("GET")
	api.HandleFunc("/foods/{foodID}", controller.GetFood).Methods("GET")
	api.HandleFunc("/foods", controller.CreateFood).Methods("POST")
	api.HandleFunc("/foods/{foodID}", controller.UpdateFood).Methods("PUT")

	//Private Routes related to Menus
	api.HandleFunc("/menus", controller.GetMenus).Methods("GET")
	api.HandleFunc("/menus/{menuID}", controller.GetMenu).Methods("GET")
	api.HandleFunc("/menus", controller.CreateMenu).Methods("POST")
	api.HandleFunc("/menus/{menuID}", controller.UpdateMenu).Methods("PUT")

	//Private Routes related to Tables
	api.HandleFunc("/tables", controller.GetTables).Methods("GET")
	api.HandleFunc("/tables/{tableID}", controller.GetTable).Methods("GET")
	api.HandleFunc("/tables", controller.CreateTable).Methods("POST")
	api.HandleFunc("/tables/{tableID}", controller.UpdateTable).Methods("PUT")

	port := os.Getenv("PORT")
	log.Fatal(http.ListenAndServe(port, r))
}
