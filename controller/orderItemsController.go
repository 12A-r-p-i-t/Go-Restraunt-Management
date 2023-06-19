package controller

import (
	"encoding/json"
	"fmt"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/12A-r-p-i-t/restraunt-management/model"
	"github.com/gorilla/mux"
)

type OrderItemPack struct {
	Table_ID   uint              `json:"table_id"`
	OrderItems []model.OrderItem `json:"order_items"`
}

func GetOrderItems(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	orderItems, err := model.GetAllOrderItems()
	if err != nil {
		log.Fatal("Error in getting all order items :", err)
		return
	}
	json.NewEncoder(w).Encode(orderItems)
}

func GetOrderItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	order_item_id := vars["orderItemID"]
	orderItemId, err := strconv.Atoi(order_item_id)
	if err != nil {
		log.Fatal("Error in converting orderItemId from string to int :", err)
		return
	}
	orderItem, err := model.GetOrderItemByID(uint(orderItemId))
	if err != nil {
		log.Fatal("Error in fetching orderItem with given ID :", err)
		return
	}
	json.NewEncoder(w).Encode(&orderItem)
}

func GetOrderItemsByOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	order_id := vars["orderID"]
	orderID, err := strconv.Atoi(order_id)
	if err != nil {
		log.Fatal("Error in converting orderID from string to int :", err)
		return
	}
	allOrderItems, err := ItemsByOrder(uint(orderID))
	if err != nil {
		log.Fatal("Error in getting all order Items :", err)
		return
	}
	json.NewEncoder(w).Encode(allOrderItems)
}

func ItemsByOrder(orderId uint) (map[string]interface{}, error) {
	getOrderItems, err := model.GetAllOrderItemsByID(orderId)
	if err != nil {
		log.Fatal("Error in fetching order Items for given order ID :", err)
	}
	getFoodItems, err := model.GetFoodItemsByOrderID(orderId)
	if err != nil {
		log.Fatal("Error in fetching order Items for given orderID :", err)
		return nil, err
	}
	var payment_due float64
	payment_due = 0
	for _, orderItem := range getOrderItems {
		for _, foodItem := range getFoodItems {
			FoodID, err := strconv.Atoi(orderItem.Food_id)
			if err != nil {
				log.Fatal("Error in converting foodID from string to int :", err)
				return nil, err
			}
			if uint(FoodID) == foodItem.ID {
				Quantity, err := strconv.Atoi(orderItem.Quantity)
				if err != nil {
					log.Fatal("Error in converting Quantity to int :", err)
				}
				payment_due += float64(Quantity) * float64(foodItem.Price)
			}
		}
	}
	getTableItems, err := model.GetTableItemsByOrderID(orderId)
	if err != nil {
		log.Fatal("Error in fetching table items for given orderID :", err)
		return nil, err
	}
	tableNumber := getTableItems[0].TableNumber
	mp := make(map[string]interface{})
	mp["payment_due"] = payment_due
	mp["order_items"] = getOrderItems
	mp["table_number"] = tableNumber
	return mp, nil

}

func CreateOrderItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")

	var orderItemPack OrderItemPack
	var order model.Order
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error in reading request body while creating order Item :", err)
		return
	}
	err = json.Unmarshal(bytes, &orderItemPack)
	if err != nil {
		log.Fatal("Error in unmarshalling data to orderItemPack struct :", err)
		return
	}
	fmt.Println(orderItemPack)
	order.Order_Date, _ = time.Parse(time.RFC3339, time.Now().Format(time.RFC3339))
	orderItemsToBeInserted := []interface{}{}
	order.Table_ID = orderItemPack.Table_ID
	order.ID = OrderItemOrderCreator(order)
	fmt.Println("Oii Oii")
	for _, orderItem := range orderItemPack.OrderItems {
		orderItem.Order_id = order.ID
		orderItemInserted := orderItem.InsertOrderItem()
		orderItemsToBeInserted = append(orderItemsToBeInserted, orderItemInserted)
	}
	fmt.Println("Haha")

	json.NewEncoder(w).Encode(orderItemsToBeInserted)
}

func UpdateOrderItem(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	order_item_id := vars["orderItemID"]
	orderItemID, err := strconv.Atoi(order_item_id)
	if err != nil {
		log.Fatal("Error in converting orderItemID from string to int :", err)
		return
	}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error in reading request body while updating order Item", err)
		return
	}
	var orderItem model.OrderItem
	err = json.Unmarshal(bytes, &orderItem)
	if err != nil {
		log.Fatal("Error in unmarshalling to orderItem struct :", err)
		return
	}
	if orderItem.Food_id != "" {
		FoodId, err := strconv.Atoi(orderItem.Food_id)
		if err != nil {
			log.Fatal("Error in converting FoodId from string to int :", err)
			return
		}
		_, err = model.GetFoodByID(uint(FoodId))
		if err != nil {
			log.Fatal("Error in fetching food with given Food ID :", err)
			return
		}
	}
	if orderItem.Order_id != 0 {
		_, err = model.GetOrderByID(orderItem.Order_id)
		if err != nil {
			log.Fatal("Error in fetching order with given order ID :", err)
			return
		}
	}
	orderItem.ID = uint(orderItemID)
	updatedOrderItem, err := orderItem.UpdateOrderItem()
	if err != nil {
		log.Fatal("Error while updating the orderItem :", err)
		return
	}
	json.NewEncoder(w).Encode(updatedOrderItem)
}
