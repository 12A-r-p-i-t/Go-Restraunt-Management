package model

import (
	"time"

	"github.com/jinzhu/gorm"
)

type Order struct {
	gorm.Model
	Order_Date time.Time `json:"order_date"`
	Table_ID   uint      `json:"table_id"`
}
