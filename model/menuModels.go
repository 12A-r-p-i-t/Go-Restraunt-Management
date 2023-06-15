package model

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type Menu struct {
	gorm.Model
	Name       string    `json:"name"`
	Category   string    `json:"category"`
	Start_Date time.Time `json:"start_date"`
	End_Date   time.Time `json:"end_date"`
}

func GetMenuByID(foodID uint) (Menu, error) {
	db := getDBInstance()

	var menu Menu
	err := db.Where("id = ?", foodID).First(&menu)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("No such record found")
		return menu, err.Error
	}
	return menu, nil
}

func GetMenus() ([]Menu, error) {
	db := getDBInstance()

	var menus []Menu
	err := db.Find(&menus)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("No records present in menus table")
		return nil, err.Error
	}
	return menus, nil
}

func (menu *Menu) InsertMenu() Menu {
	db := getDBInstance()

	db.NewRecord(menu)
	db.Create(menu)
	return *menu
}

func (menu *Menu) UpdateMenu() (Menu, error) {
	db := getDBInstance()

	err := db.Model(&menu).Update(menu)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("No such record found with give menuID to update :", err.Error)
		return Menu{}, err.Error
	}
	return GetMenuByID(menu.ID)
}
