package invoiceheader

import (
	"database/sql"
	"time"
)

// Modelo de invoiceheader
type Model struct {
	ID        uint
	Client    string
	CreatedAt time.Time
	UpdatedAt time.Time
}

type Storage interface {
	Migrate() error
	CreateTx(*sql.Tx, *Model) error
}

// Servicio del producto
type Service struct {
	storage Storage
}

// New service retorna un puntero de servicio
func NewService(s Storage) *Service {
	return &Service{s}
}

// Es usado para migrar el produto
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
