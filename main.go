package main

import (
	"log"

	"github.com/saalcazar/GoDB/pkg/product"
	"github.com/saalcazar/GoDB/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)
	err := serviceProduct.Delete(3)
	if err != nil {
		log.Fatalf("product.Update: %v", err)
	}
}
