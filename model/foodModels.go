package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Food struct {
	gorm.Model
	Name       string  `json:"name"`
	Price      float64 `json:"price"`
	Food_image string  `json:"food_image"`
	Menu_id    uint    `json:"menu_id"`
}

func GetAllFoods() ([]Food, error) {
	db := getDBInstance()

	var Foods []Food
	err := db.Find(&Foods)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("No record found in Foods database")
		return nil, err.Error
	}
	return Foods, nil
}

func GetFoodByID(foodID uint) (Food, error) {
	db := getDBInstance()

	var food Food
	err := db.Where("id = ?", foodID).Find(&food)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("No such food found with given ID")
		return food, err.Error
	}
	return food, nil
}

func (food *Food) InsertFood() Food {
	db := getDBInstance()

	db.NewRecord(food)
	db.Create(food)
	return *food
}

func (food *Food) UpdateFood() (Food, error) {
	db := getDBInstance()

	err := db.Model(&food).Update(food)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("No such record found with give ID to update")
		return Food{}, err.Error
	}
	return GetFoodByID(food.ID)
}
