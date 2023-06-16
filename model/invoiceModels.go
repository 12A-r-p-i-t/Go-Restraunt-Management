package model

import (
	"log"
	"time"

	"github.com/jinzhu/gorm"
)

type Invoice struct {
	gorm.Model
	Order_id         string    `json:"order_id"`
	Payment_method   string    `json:"payment_method"`
	Payment_status   string    `json:"payment_status"`
	Payment_due_date time.Time `json:"payment_due_date"`
}

func GetInvoices() ([]Invoice, error) {
	db := getDBInstance()

	var invoice []Invoice
	err := db.Find(&invoice)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("Error in gettting all invoices from DB :", err)
		return nil, err.Error
	}
	return invoice, nil
}

func GetInvoiceByID(invoiceID uint) (Invoice, error) {
	db := getDBInstance()

	var invoice Invoice
	err := db.Where("id = ?", invoiceID).First(&invoice)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("Error in getting invoice for given ID :", err.Error)
		return Invoice{}, err.Error
	}
	return invoice, nil
}

func (invoice *Invoice) InsertInvoice() Invoice {
	db := getDBInstance()

	db.NewRecord(invoice)
	db.Create(invoice)
	return *invoice
}

func (invoice *Invoice) UpdateInvoice() (Invoice, error) {
	db := getDBInstance()

	err := db.Model(&invoice).Update(invoice)
	if gorm.IsRecordNotFoundError(err.Error) {
		log.Fatal("Error in updating invoice :", err)
		return Invoice{}, err.Error
	}
	return *invoice, nil
}
