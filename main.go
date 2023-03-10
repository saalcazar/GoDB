package main

import (
	"database/sql"
	"errors"
	"fmt"
	"log"

	"github.com/saalcazar/GoDB/pkg/product"
	"github.com/saalcazar/GoDB/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m, err := serviceProduct.GetById(17)
	switch {
	case errors.Is(err, sql.ErrNoRows):
		fmt.Println("No existe un producto con ese ID")
	case err != nil:
		log.Fatalf("product.GetByID: %v", err)
	default:
		fmt.Println(m)
	}
}
