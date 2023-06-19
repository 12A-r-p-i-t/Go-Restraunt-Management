package model

import (
	"fmt"
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
	fmt.Println(result)
	return result, nil
}

func GetFoodItemsByOrderID(orderID uint) ([]Food, error) {
	db := getDBInstance()

	var results []Food
	if err := db.Table("order_items").Where("order_items.order_id = ?", orderID).Joins("JOIN foods on foods.id = order_items.food_id").Find(&results).Error; err != nil {
		log.Fatal("Error in joining the table and getting output :", err)
		return nil, err
	}

	// if err := db.Table("orderItems").Joins("JOIN orders on orders.id = orderItems.order_id").Joins("JOIN foods on foods.id = orderItems.food_id").Find(&results).Error; err != nil {
	// 	log.Fatal("Error in joining the table and getting output :", err)
	// 	return nil, err
	// }
	fmt.Println(results[0].ID)
	fmt.Println(results[0].Menu_id)
	fmt.Println(results[0].Food_image)
	fmt.Println(results[0].Price)
	fmt.Println(results)
	return results, nil

	// var orderItems []OrderItem
	// err := db.Where("id = ?", orderID).Find(&orderItems)
	// if gorm.IsRecordNotFoundError(err.Error) {
	// 	log.Fatal("Error in getting all the order Items for given order ID :", err)
	// 	return nil, err.Error
	// }
	// return orderItems, nil
}

func GetTableItemsByOrderID(orderID uint) ([]Table, error) {
	db := getDBInstance()
	fmt.Println("oii")
	var results []Table
	if err := db.Table("orders").Where("orders.id = ?", orderID).Joins("JOIN tables on tables.id = orders.table_id").Find(&results).Error; err != nil {
		log.Fatal("Error in joining the table and getting output :", err)
		return nil, err
	}
	// if err := db.Table("orders").Joins("JOIN orderItems on orderItems.order_id = orders.id").Joins("JOIN tables on table.id = orders.table_id").Find(&results).Error; err != nil {
	// 	log.Fatal("Error in joining the table and getting output :", err)
	// 	return nil, err
	// }
	fmt.Println(results)
	return results, nil
}

func GetAllOrderItemsByID(orderID uint) ([]OrderItem, error) {
	db := getDBInstance()

	var orderItems []OrderItem
	if err := db.Table("order_items").Where("order_id = ?", orderID).Find(&orderItems).Error; err != nil {
		log.Fatal("Error in getting all orderItems from DB :", err)
		return nil, err
	}
	fmt.Println(orderItems)
	return orderItems, nil
}
