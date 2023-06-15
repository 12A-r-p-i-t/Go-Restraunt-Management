package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"

	"github.com/12A-r-p-i-t/restraunt-management/model"
	"github.com/gorilla/mux"
)

func GetFoods(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allFoods, err := model.GetAllFoods()
	if err != nil {
		log.Fatal("Error in fetching foods from database :", err)
	}
	json.NewEncoder(w).Encode(allFoods)
}

func GetFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	food_id := vars["foodID"]
	ID, err := strconv.Atoi(food_id)
	if err != nil {
		log.Fatal("Error in converting string to int in GetFood Route")
		return
	}
	food, err := model.GetFoodByID(uint(ID))
	if err != nil {
		log.Fatal("Error in fetching food with given ID from database")
		return
	}
	json.NewEncoder(w).Encode(food)
}

func CreateFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error in reading request body :", err)
		return
	}
	var food model.Food
	err = json.Unmarshal(bytes, &food)
	if err != nil {
		log.Fatal("Error in unmarshalling food data to struct :", err)
		return
	}
	menuID := food.Menu_id
	_, err = model.GetMenuByID(uint(menuID))
	if err != nil {
		log.Fatal("No such menu find in menu database :", err)
	}
	insertedFood := food.InsertFood()
	json.NewEncoder(w).Encode(insertedFood)
}

func UpdateFood(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	food_ID := vars["foodID"]
	foodID, err := strconv.Atoi(food_ID)
	if err != nil {
		log.Fatal("Error in converting foodID from string to int :", err)
		return
	}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error in reading from request body during updateFood")
		return
	}
	var food model.Food
	err = json.Unmarshal(bytes, &food)
	if err != nil {
		log.Fatal("Error in unmarshalling data to food struct in update food")
		return
	}
	if food.Menu_id != 0 {
		menuID := food.Menu_id
		_, err = model.GetMenuByID(uint(menuID))
		if err != nil {
			log.Fatal("Error in getting menu with give menu_id")
			return
		}
	}
	food.ID = uint(foodID)
	updateFood, err := food.UpdateFood()
	if err != nil {
		log.Fatal("Error in updating Food :", err)
		return
	}
	json.NewEncoder(w).Encode(updateFood)
}
