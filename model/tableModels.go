package model

import (
	"log"

	"github.com/jinzhu/gorm"
)

type Table struct {
	gorm.Model
	NumberOfGuest int `json:"number_of_guest"`
	TableNumber   int `json:"table_number"`
}

func GetTables() ([]Table, error) {
	db := getDBInstance()

	var tables []Table
	err := db.Find(&tables)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("No tables found in Tables database :", err)
		return nil, err.Error
	}
	return tables, nil
}

func GetTableByID(tableID uint) (Table, error) {
	db := getDBInstance()

	var table Table
	err := db.Where("id = ?", tableID).First(&table)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("No such table found with given ID :", err)
		return table, err.Error
	}
	return table, nil
}

func (table *Table) InsertTable() Table {
	db := getDBInstance()

	db.NewRecord(table)
	db.Create(table)
	return *table
}

func (table *Table) UpdateTable() (Table, error) {
	db := getDBInstance()

	err := db.Model(&table).Update(table)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("No table found with given ID while updating table :", err.Error)
		return Table{}, err.Error
	}
	return GetTableByID(table.ID)
}
