package controller

import (
	"encoding/json"
	"io"
	"log"
	"net/http"
	"strconv"
	"time"

	"github.com/12A-r-p-i-t/restraunt-management/model"
	"github.com/gorilla/mux"
)

type InvoiceViewFormat struct {
	Invoice_id       uint
	Payment_method   string
	Order_id         string
	Payment_status   string
	Payment_due      interface{}
	Table_number     interface{}
	Payment_due_date time.Time
	Order_details    interface{}
}

func GetInvoices(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	allInvoices, err := model.GetInvoices()
	if err != nil {
		log.Fatal("Error in fetching all invoices from database :", err)
		return
	}
	json.NewEncoder(w).Encode(allInvoices)
}

func GetInvoice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	invoice_ID := vars["invoiceID"]
	invoiceID, err := strconv.Atoi(invoice_ID)
	if err != nil {
		log.Fatal("Error in converting invoiceID from string to int :", err)
		return
	}
	invoice, err := model.GetInvoiceByID(uint(invoiceID))
	if err != nil {
		log.Fatal("Error in getting invoice by ID :", err)
		return
	}
	var invoiceView InvoiceViewFormat

	allOrderItems, err := ItemsByOrder(invoice.Order_id)
	invoiceView.Order_id = invoice.Order_id
	invoiceView.Payment_due_date = invoice.Payment_due_date
	invoiceView.Payment_method = ""
	if invoice.Payment_method != "" {
		invoiceView.Payment_method = invoice.Payment_method
	}

	invoiceView.Invoice_id = invoice.ID
	invoiceView.Payment_status = invoice.Payment_status
	invoiceView.Payment_due = allOrderItems[0]["payment_due"]
	invoiceView.Table_number = allOrderItems[0]["table_number"]
	invoiceView.Order_details = allOrderItems[0]["order_items"]

	json.NewEncoder(w).Encode(&invoiceView)
}

func CreateInvoice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error in reading from request body while creating invoice :", err)
		return
	}
	var invoice model.Invoice
	err = json.Unmarshal(bytes, &invoice)
	if err != nil {
		log.Fatal("Error in unmarshalling data to invoice struct :", err)
		return
	}
	order_ID := invoice.Order_id
	orderID, err := strconv.Atoi(order_ID)
	if err != nil {
		log.Fatal("Error in converting order_ID from string to int :", err)
		return
	}
	_, err = model.GetOrderByID(uint(orderID))
	if err != nil {
		log.Fatal("No such order found with given orderID :", err)
		return
	}
	if invoice.Payment_status == "" {
		invoice.Payment_status = "Pending"
	}
	insertedInvoice := invoice.InsertInvoice()
	json.NewEncoder(w).Encode(insertedInvoice)
}

func UpdateInvoice(w http.ResponseWriter, r *http.Request) {
	w.Header().Set("Content-Type", "application/json")
	vars := mux.Vars(r)
	invoice_id := vars["invoiceID"]
	invoiceID, err := strconv.Atoi(invoice_id)
	if err != nil {
		log.Fatal("Error in converting invoiceID from string to int :", err)
		return
	}
	bytes, err := io.ReadAll(r.Body)
	if err != nil {
		log.Fatal("Error in reading from request body while updating invoice :", err)
		return
	}
	var invoice model.Invoice
	err = json.Unmarshal(bytes, &invoice)
	if err != nil {
		log.Fatal("Error in unmarshalling data to invoice struct :", err)
		return
	}
	if invoice.Order_id != "" {
		order_id := invoice.Order_id
		orderID, err := strconv.Atoi(order_id)
		if err != nil {
			log.Fatal("Error in converting order id from string to int in update invoice :", err)
			return
		}
		_, err = model.GetOrderByID(uint(orderID))
		if err != nil {
			log.Fatal("Error in getting order with given OrderID :", err)
			return
		}
	}
	invoice.ID = uint(invoiceID)
	updatedInvoice, err := invoice.UpdateInvoice()
	if err != nil {
		log.Fatal("Erro in updating invoice :", err)
		return
	}
	json.NewEncoder(w).Encode(updatedInvoice)
}
