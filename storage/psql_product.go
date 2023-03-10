package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/GoDB/pkg/product"
)

const (
	psqlMigrateProduct = `CREATE TABLE IF NOT EXISTS products(
		id SERIAL NOT NULL,
		name VARCHAR(25) NOT NULL,
		observation VARCHAR(100),
		price INT NOT NULL,
		created_at TIMESTAMP NOT NULL DEFAULT now(),
		updated_at TIMESTAMP,
		CONSTRAINT pk_products PRIMARY KEY (id)
	)`
	psqlCreateProduct = `INSERT INTO products (name, observation, price, created_at)
	VALUES($1, $2, $3, $4) RETURNING id`
)

// usado para trabajar con postgrs y el paquete
type PsqlProduct struct {
	db *sql.DB
}

// Retorna un nuevo puntero de PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}
}

// Implementa la interface Product.store
func (p *PsqlProduct) Migrate() error {
	stmt, err := p.db.Prepare(psqlMigrateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()
	if err != nil {
		return err
	}

	fmt.Println("Migración de product ejecutada correctamente")
	return nil
}

// Insertar datos
// Implementa la interface de prodcut.Storage
func (p *PsqlProduct) Create(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlCreateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()
	err = stmt.QueryRow(
		m.Name,
		stringToNull(m.Observation),
		m.Price,
		m.CreatedAt,
	).Scan(&m.ID)
	if err != nil {
		return err
	}

	fmt.Println("Se creo el producto correctamente")
	return nil
}
