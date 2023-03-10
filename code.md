## Migrar tabla de productos
```go
storageProduct := storage.NewPsqlProduct(storage.Pool())
serviceProduct := product.NewService(storageProduct)
if err := serviceProduct.Migrate(); err != nil {
log.Fatalf("product.Migrate: %v", err)
}
```
## Migrar la tabla InvoiceHeader
```go
storageInvoiceHeader := storage.NewPsqlInvoiceHeader(storage.Pool())
serviceInvoiceHeader := invoiceheader.NewService(storageInvoiceHeader)
if err := serviceInvoiceHeader.Migrate(); err != nil {
log.Fatalf("invoiceHeader.Migrate: %v", err)
}
```
## Migrar la tabla InvoiceItem
```go
storageInvoiceItem := storage.NewPsqlInvoiceItem(storage.Pool())
serviceInvoiceItem := invoiceitem.NewService(storageInvoiceItem)
if err := serviceInvoiceItem.Migrate(); err != nil {
	log.Fatalf("invoiceItem.Migrate; %v", err)
}
```

## Ingresar datos a una tabla
```go
	storageProduct := storage.NewPsqlProduct(storage.Pool())
	serviceProduct := product.NewService(storageProduct)

	m := &product.Model{
		Name:        "HTML desde cero",
		Observation: "iniciate en este mundo del diseño web",
		Price:       90,
	}

	if err := serviceProduct.Create(m); err != nil {
		log.Fatalf("product.Create: %v", err)
	}

	fmt.Printf("%+v\n", m)
```