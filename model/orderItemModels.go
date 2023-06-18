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

func GetFoodItems() ([]Food, error) {
	db := getDBInstance()

	var result []Food
	err := db.Model(&OrderItem{}).Joins("left join orderItems on orderItems.Food_id = foods.ID").Find(&result)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("Error in getting food Items by joining order items and food items table", err.Error)
		return nil, err.Error
	}
	return result, nil
}

func GetOrderItemsByOrderID(orderID uint) ([]Food, error) {
	db := getDBInstance()

	var results []Food
	if err := db.Table("orderItems").Joins("JOIN orders on order.id = orderItems.order_id").Joins("JOIN foods on food.id = orderItems.food_id").Find(&results).Error; err != nil {
		log.Fatal("Error in joining the table and getting output :", err)
		return nil, err
	}
	return results, nil

	// var orderItems []OrderItem
	// err := db.Where("id = ?", orderID).Find(&orderItems)
	// if gorm.IsRecordNotFoundError(err.Error) {
	// 	log.Fatal("Error in getting all the order Items for given order ID :", err)
	// 	return nil, err.Error
	// }
	// return orderItems, nil
}
