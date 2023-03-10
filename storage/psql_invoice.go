package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/GoDB/pkg/invoice"
	"github.com/saalcazar/GoDB/pkg/invoiceheader"
	"github.com/saalcazar/GoDB/pkg/invoiceitem"
)

// PsqlInvoice usado para trabajar con postgres y factura
type PsqlInvoice struct {
	db            *sql.DB
	storageHeader invoiceheader.Storage
	storageItems  invoiceitem.Storage
}

// Retorna un nuevo puntero de PSQLINVOICE
func NewPsqlInvoice(db *sql.DB, h invoiceheader.Storage, i invoiceitem.Storage) *PsqlInvoice {
	return &PsqlInvoice{
		db:            db,
		storageHeader: h,
		storageItems:  i,
	}
}

// Implementa la interface InvoiceStorage
func (p *PsqlInvoice) Create(m *invoice.Model) error {
	tx, err := p.db.Begin()
	if err != nil {
		return err
	}

	if err := p.storageHeader.CreateTx(tx, m.Header); err != nil {
		tx.Rollback()
		return fmt.Errorf("header: %w", err)
	}

	fmt.Printf("Factura creada con id; %d \n", m.Header.ID)

	if err := p.storageItems.CreateTx(tx, m.Header.ID, m.Items); err != nil {
		tx.Rollback()
		return err
	}
	fmt.Printf("Items creados; %d \n", len(m.Items))

	return tx.Commit()

}
