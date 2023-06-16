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

func GetOrders(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allOrders, err := model.GetOrders()
	if err != nil {
		log.Fatal("Error in getting orders from database :", err)
		return
	}
	json.NewEncoder(w).Encode(allOrders)
}

func GetOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	order_id := vars["orderID"]
	orderID, err := strconv.Atoi(order_id)
	if err != nil {
		log.Fatal("Error in converting orderID from string to int :", err)
		return
	}
	order, err := model.GetOrderByID(uint(orderID))
	if err != nil {
		log.Fatal("Error in getting order for given ID :", err)
		return
	}
	json.NewEncoder(w).Encode(order)
}

func CreateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error in reading the request body while creating a order :", err)
		return
	}
	var order model.Order
	err = json.Unmarshal(bytes, &order)
	if err != nil {
		log.Fatal("Error in unmarshalling data to order struct :", err)
		return
	}
	if order.Table_ID != 0 {
		_, err := model.GetTableByID(order.Table_ID)
		if err != nil {
			log.Fatal("Error in fetching the table with given tableID :", err)
			return
		}
	}
	insertedOrder := order.InsertOrder()
	json.NewEncoder(w).Encode(insertedOrder)
}

func UpdateOrder(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	order_id := vars["orderID"]
	orderID, err := strconv.Atoi(order_id)
	if err != nil {
		log.Fatal("Error in converting orderID from string to int :", err)
		return
	}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error in reading the request body while updating the order :", err)
		return
	}
	var order model.Order
	err = json.Unmarshal(bytes, &order)
	if err != nil {
		log.Fatal("Error in unmarshalling data to order model :", err)
		return
	}
	order.ID = uint(orderID)
	if order.Table_ID != 0 {
		_, err = model.GetTableByID(uint(order.Table_ID))
		if err != nil {
			log.Fatal("No such table exists with given tableID :", err)
			return
		}
	}
	updatedOrder, err := order.UpdateOrder()
	if err != nil {
		log.Fatal("Error while updating the order :", err)
		return
	}
	json.NewEncoder(w).Encode(updatedOrder)
}

func OrderItemOrderCreator(order model.Order) uint {

	insertedOrder := order.InsertOrder()

	return insertedOrder.ID
}
