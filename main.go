package main

import (
	"fmt"
	"log"

	"github.com/saalcazar/GoDB/pkg/product"
	"github.com/saalcazar/GoDB/storage"
)

func main() {
	storage.NewPostgresDB()

	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		Name:        "HTML desde cero",
		Observation: "iniciate en este mundo del diseño",
		Price:       90,
	}

	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Create: %v", err)
	}

	fmt.Printf("%+v\n", m)
}
