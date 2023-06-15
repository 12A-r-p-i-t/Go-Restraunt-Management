package model

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	Order_Date time.Time `json:"order_date"`
	Table_ID   uint      `json:"table_id"`
}

func GetOrders() ([]Order, error) {
	db := getDBInstance()

	var orders []Order
	err := db.Find(&orders)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("No records found in orders table :", err.Error)
		return nil, err.Error
	}
	return orders, nil
}

func GetOrderByID(orderID uint) (Order, error) {
	db := getDBInstance()

	var order Order
	err := db.Where("id = ?", orderID).First(&order)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("Error in finding the order with given ID :", err)
		return Order{}, err.Error
	}
	return order, nil
}

func (order *Order) InsertOrder() Order {
	db := getDBInstance()

	db.NewRecord(order)
	db.Create(order)
	return *order
}

func (order *Order) UpdateOrder() (Order, error) {
	db := getDBInstance()

	err := db.Model(&order).Update(order)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("Error while updating the order for given ID :", err)
		return Order{}, err.Error
	}
	return GetOrderByID(order.ID)
}
