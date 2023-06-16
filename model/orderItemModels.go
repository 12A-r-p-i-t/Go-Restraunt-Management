package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

type OrderItem struct {
	gorm.Model
	Quantity   string  `json:"quantity"`
	Unit_price float64 `json:"unit_price"`
	Food_id    string  `json:"food_id"`
	Order_id   uint    `json:"order_id"`
}

func GetAllOrderItems() ([]OrderItem, error) {
	db := getDBInstance()

	var allOrderItems []OrderItem
	err := db.Find(&allOrderItems)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("Error in fetching all order items :", err.Error)
		return nil, err.Error
	}
	return allOrderItems, nil
}

func GetOrderItemByID(orderItemID uint) (OrderItem, error) {
	db := getDBInstance()

	var orderItem OrderItem
	err := db.Where("id = ?", orderItemID).First(&orderItem)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("No such order Item found with given orderItemID :", err.Error)
		return OrderItem{}, err.Error
	}
	return orderItem, nil
}

func (orderItem *OrderItem) InsertOrderItem() OrderItem {
	db := getDBInstance()

	db.NewRecord(orderItem)
	db.Create(orderItem)
	return *orderItem
}

func (orderItem *OrderItem) UpdateOrderItem() (OrderItem, error) {
	db := getDBInstance()

	err := db.Model(&orderItem).Update(orderItem)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("No such record found with given orderItemId to update :", err.Error)
		return OrderItem{}, err.Error
	}
	return GetOrderItemByID(orderItem.ID)
}
