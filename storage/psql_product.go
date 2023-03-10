package storage

import (
	"database/sql"
	"fmt"

	"github.com/saalcazar/GoDB/pkg/product"
)

type scanner interface {
	Scan(dest ...any) error
}

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

	psqlGetAllProduct = `SELECT id, name, observation, price, created_at, updated_at FROM products`

	psqlGetProductById = psqlGetAllProduct + " WHERE id = $1"

	psqlUpdateProduct = `UPDATE products SET name = $1, observation = $2, price = $3, updated_at = $4 WHERE id = $5`

	psqlDeleteProduct = `DELETE FROM products WHERE id = $1`
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

// Implementa la interface product storage
func (p *PsqlProduct) GetAll() (product.Models, error) {
	stmt, err := p.db.Prepare(psqlGetAllProduct)
	if err != nil {
		return nil, err
	}
	defer stmt.Close()

	rows, err := stmt.Query()
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	ms := make(product.Models, 0)
	for rows.Next() {
		m, err := scanRowProduct(rows)
		if err != nil {
			return nil, err
		}
		ms = append(ms, m)
	}
	if err := rows.Err(); err != nil {
		return nil, err
	}

	return ms, nil
}

// GetById implementa la interface product.storre
func (p *PsqlProduct) GetByID(id uint) (*product.Model, error) {
	stmt, err := p.db.Prepare(psqlGetProductById)
	if err != nil {
		return &product.Model{}, err
	}
	defer stmt.Close()
	return scanRowProduct(stmt.QueryRow(id))
}

// Implementa la interface
func (p *PsqlProduct) Update(m *product.Model) error {
	stmt, err := p.db.Prepare(psqlUpdateProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	res, err := stmt.Exec(
		&m.Name,
		stringToNull(m.Observation),
		&m.Price,
		timeToNull(m.UpdatedAt),
		&m.ID,
	)
	if err != nil {
		return err
	}

	rowsAfected, err := res.RowsAffected()
	if err != nil {
		return err
	}

	if rowsAfected == 0 {
		return fmt.Errorf("no existe el producto con id: %d", m.ID)
	}

	fmt.Println("Se actualizo el producto correctamente")
	return nil

}

func (p *PsqlProduct) Delete(id uint) error {
	stmt, err := p.db.Prepare(psqlDeleteProduct)
	if err != nil {
		return err
	}
	defer stmt.Close()

	_, err = stmt.Exec(id)
	if err != nil {
		return err
	}

	fmt.Println("Se elimino el producto correctamente")
	return nil
}

func scanRowProduct(s scanner) (*product.Model, error) {
	m := &product.Model{}
	observationNull := sql.NullString{}
	updatedatNull := sql.NullTime{}
	err := s.Scan(
		&m.ID,
		&m.Name,
		&observationNull,
		&m.Price,
		&m.CreatedAt,
		&updatedatNull,
	)
	if err != nil {
		return &product.Model{}, err
	}
	m.Observation = observationNull.String
	m.UpdatedAt = updatedatNull.Time
	return m, nil
}
