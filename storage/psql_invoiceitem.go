package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/GoDB/pkg/invoiceitem"
)

const (
	psqlMigrateInvoiceItem = `CREATE TABLE IF NOT EXISTS invoice_items(
		id SERIAL NOT NULL,
		invoice_header_id INT NOT NULL,
		product_id INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT pk_invoice_items PRIMARY KEY (id),
		CONSTRAINT fk_invoice_header_id FOREIGN KEY (invoice_header_id) REFERENCES invoice_headers (id) ON UPDATE RESTRICT ON DELETE RESTRICT,
		CONSTRAINT fk_product_id FOREIGN KEY (product_id) REFERENCES products (id) ON UPDATE RESTRICT ON DELETE RESTRICT
	)`

	psqlCreateInvoiceItem = `INSERT INTO invoice_items(invoice_header_id, product_id) VALUES($1, $2) RETURNING id, created_at`
)

// usado para trabajar con postgrs y el paquete
type PsqlInvoiceItem struct {
	db *sql.DB
}

// Retorna un nuevo puntero de PsqlProduct
func NewPsqlInvoiceItem(db *sql.DB) *PsqlInvoiceItem {
	return &PsqlInvoiceItem{db}
}

// Implementa la interface Product.store
func (p *PsqlInvoiceItem) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migración de invoice item ejecutada correctamente")
	return nil
}

func (p *PsqlInvoiceItem) CreateTx(tx *sql.Tx, headerID uint, ms invoiceitem.Models) error {
	stmt, err := tx.Prepare(psqlCreateInvoiceItem)
	if err != nil {
		return err
	}
	defer stmt.Close()

	for _, item := range ms {
		err = stmt.QueryRow(headerID, item.ProductID).Scan(
			&item.ID,
			&item.CreatedAt,
		)
		if err != nil {
			return err
		}
	}
	return nil
}
