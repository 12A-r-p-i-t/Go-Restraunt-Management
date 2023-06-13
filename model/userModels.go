package model

import (
	"fmt"

	"github.com/12A-r-p-i-t/restraunt-management/database"
	"github.com/jinzhu/gorm"
)

type User struct {
	gorm.Model
	First_name string `json:"first_name"`
	Last_name  string `json:"last_name"`
	Password   string `json:"password"`
	Email      string `json:"email"`
	Avatar     string `json:"avatar"`
	Phone      string `json:"phone"`
}

func getDBInstance() *gorm.DB {
	db := database.DB
	return db
}

func CheckEmail(email string) bool {
	db := getDBInstance()
	var user User
	err := db.Where("email = ?", email).First(&user)
	if gorm.IsRecordNotFoundError(err.Error) {
		fmt.Println("No such email exists in the database")
		return false
	}
	return true
}

func CheckPhone(phone string) bool {
	db := getDBInstance()
	var user User
	err := db.Where("phone = ?", phone).First(&user)
	if gorm.IsRecordNotFoundError(err.Error) {
		fmt.Println("No such phone number exists in the database")
		return false
	}
	return true
}

func (user *User) InsertUser() *User {
	db := getDBInstance()

	db.NewRecord(user)
	db.Create(user)
	return user
}

func (user *User) FindUser() (*User, error) {
	db := getDBInstance()

	var foundUser User
	err := db.Where("email = ?", user.Email).First(&foundUser)
	if gorm.IsRecordNotFoundError(err.Error) {
		fmt.Println("No such account with given emailID exists in the database")
		return nil, err.Error
	}
	return &foundUser, nil
}
