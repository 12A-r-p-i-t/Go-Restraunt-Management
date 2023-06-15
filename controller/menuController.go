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

func GetMenus(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	menus, err := model.GetMenus()
	if err != nil {
		log.Fatal("Error in fetching all menus from database :", err)
		return
	}
	json.NewEncoder(w).Encode(menus)
}

func GetMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	menu_ID := vars["menuID"]
	menuID, err := strconv.Atoi(menu_ID)
	if err != nil {
		log.Fatal("Error in converting menu_id from string to int :", err)
		return
	}
	menu, err := model.GetMenuByID(uint(menuID))
	if err != nil {
		log.Fatal("Error in fetching given menuID from Database :", err)
	}
	json.NewEncoder(w).Encode(menu)
}

func CreateMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error in reading data from request body while creating a new menu :", err)
		return
	}
	var newMenu model.Menu
	err = json.Unmarshal(bytes, &newMenu)
	if err != nil {
		log.Fatal("Error in unmarshalling data to newMenu struct :", err)
		return
	}
	insertedMenu := newMenu.InsertMenu()
	json.NewEncoder(w).Encode(insertedMenu)
}

func UpdateMenu(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	menu_id := vars["menuID"]
	menuID, err := strconv.Atoi(menu_id)
	if err != nil {
		log.Fatal("Error in converting menuID from string to int :", err)
		return
	}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error in reading request body while updating menu :", err)
		return
	}
	var updatedMenu model.Menu
	err = json.Unmarshal(bytes, &updatedMenu)
	if err != nil {
		log.Fatal("Error in unmarshalling data to updated Menu :", err)
		return
	}
	updatedMenu.ID = uint(menuID)
	newMenu, err := updatedMenu.UpdateMenu()
	if err != nil {
		log.Fatal("Error in updating Menu :", err)
		return
	}
	json.NewEncoder(w).Encode(newMenu)
}
