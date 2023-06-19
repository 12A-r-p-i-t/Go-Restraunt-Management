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
	err := db.Model(&OrderItem{}).Joins("left join order_items on order_items.food_id = foods.id").Find(&result)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("Error in getting food Items by joining order items and food items table", err.Error)
		return nil, err.Error
	}
	return result, nil
}

func GetFoodItemsByOrderID(orderID uint) ([]Food, error) {
	db := getDBInstance()

	var results []Food
	var orderItems []OrderItem
	if err := db.Table("order_items").Where("order_id = ?", orderID).Find(&orderItems).Error; err != nil {
		log.Fatal("Error in fetching order Items from db :", err)
		return nil, err
	}
	for _, orderItem := range orderItems {
		var food Food
		if err := db.Table("foods").Where("id = ?", orderItem.Food_id).Find(&food).Error; err != nil {
			log.Fatal("Erro in fetching food items from db :", err)
			return nil, err
		}
		results = append(results, food)
	}
	return results, nil
}

func GetTableItemsByOrderID(orderID uint) ([]Table, error) {
	db := getDBInstance()
	var results []Table
	var orders []Order
	if err := db.Table("orders").Where("id = ?", orderID).Find(&orders).Error; err != nil {
		log.Fatal("Error in fetching orders from DB :", err)
		return nil, err
	}
	for _, order := range orders {
		var table Table
		if err := db.Table("tables").Where("id = ?", order.Table_ID).Find(&table).Error; err != nil {
			log.Fatal("Error in fetching table from DB :", err)
			return nil, err
		}
		results = append(results, table)
	}
	return results, nil
}

func GetAllOrderItemsByID(orderID uint) ([]OrderItem, error) {
	db := getDBInstance()

	var orderItems []OrderItem
	if err := db.Table("order_items").Where("order_id = ?", orderID).Find(&orderItems).Error; err != nil {
		log.Fatal("Error in getting all orderItems from DB :", err)
		return nil, err
	}
	return orderItems, nil
}
