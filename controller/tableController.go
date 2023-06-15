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

func GetTables(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	tables, err := model.GetTables()
	if err != nil {
		log.Fatal("Error in fetching tables from database :", err)
		return
	}
	json.NewEncoder(w).Encode(tables)
}

func GetTable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	table_ID := vars["tableID"]
	tableID, err := strconv.Atoi(table_ID)
	if err != nil {
		log.Fatal("Error in converting tableID from string to int :", err)
		return
	}
	table, err := model.GetTableByID(uint(tableID))
	if err != nil {
		log.Fatal("Error in fetching the table for the given ID :", err)
		return
	}
	json.NewEncoder(w).Encode(table)
}

func CreateTable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error in reading from request body while creating table :", err)
		return
	}
	var table model.Table
	err = json.Unmarshal(bytes, &table)
	if err != nil {
		log.Fatal("Error in unmarshalling data to table struct :", err)
		return
	}
	insertedTable := table.InsertTable()
	json.NewEncoder(w).Encode(&insertedTable)
}

func UpdateTable(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error in reading data from request body while updating table :", err)
		return
	}
	vars := mux.Vars(r)
	table_id := vars["tableID"]
	tableID, err := strconv.Atoi(table_id)
	if err != nil {
		log.Fatal("Error while converting table_ID string to int :", err)
		return
	}
	var newTable model.Table
	err = json.Unmarshal(bytes, &newTable)
	if err != nil {
		log.Fatal("Error in unmarshalling data to newTable struct :", err)
		return
	}
	newTable.ID = uint(tableID)
	updatedTable, err := newTable.UpdateTable()
	if err != nil {
		log.Fatal("Error while updating the table :", err)
		return
	}
	json.NewEncoder(w).Encode(updatedTable)
}
