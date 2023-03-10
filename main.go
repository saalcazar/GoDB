package main

import (
	"log"

	"github.com/saalcazar/GoDB/pkg/invoice"
	"github.com/saalcazar/GoDB/pkg/invoiceheader"
	"github.com/saalcazar/GoDB/pkg/invoiceitem"
	"github.com/saalcazar/GoDB/storage"
)

func main() {
	storage.NewPostgresDB()
	storageHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
	storageItems := storage.NewPsqlInvoiceItem(storage.Pool())

	storageInvoice := storage.NewPsqlInvoice(storage.Pool(), storageHeader, storageItems)

	m := &invoice.Model{
		Header: &invoiceheader.Model{
			Client: "Alejandro Cabrera",
		},
		Items: invoiceitem.Models{
			&invoiceitem.Model{ProductID: 7},
			&invoiceitem.Model{ProductID: 9},
			&invoiceitem.Model{ProductID: 11},
		},
	}

	serviceInvoice := invoice.NewService(storageInvoice)
	if err := serviceInvoice.Create(m); err != nil {
		log.Fatalf("Invoice.Create: %v", err)
	}
}
