package storage

import (
	"database/sql"
	"fmt"
)

const (
	psqlMigrateInvoiceHeader = `CREATE TABLE IF NOT EXISTS invoice_headers(
		id SERIAL NOT NULL,
		client VARCHAR(100) NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT pk_Invoice_headers PRIMARY KEY (id)
	)`
)

// usado para trabajar con postgrs y el paquete InvoiceHeader
type PsqlInvoiceHeader struct {
	db *sql.DB
}

// Retorna un nuevo puntero de PsqlProduct
func NewPsqlInvoiceHeader(db *sql.DB) *PsqlInvoiceHeader {
	return &PsqlInvoiceHeader{db}
}

// Implementa la interface invoiceheader.store
func (p *PsqlInvoiceHeader) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceHeader)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migración de Invoice Header ejecutada correctamente")
	return nil
}
