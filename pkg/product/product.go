package product

import (
	"fmt"
	"time"
)

// Modelo de producto
type Model struct {
	ID          uint
	Name        string
	Observation string
	Price       int
	CreatedAt   time.Time
	UpdatedAt   time.Time
}

// Implementa la interface Stringer del paquete format - alpimprimir la información esta este tabulada y visualmente legible
func (m *Model) String() string {
	return fmt.Sprintf(
		"%02d | %-20s | %-20s | %5d | %10s | %10s",
		m.ID, m.Name, m.Observation, m.Price,
		m.CreatedAt.Format("2006-01-02"), m.UpdatedAt.Format("2006-01-02"))
}

// Slice de Modelo
type Models []*Model

type Storage interface {
	Migrate() error
	Create(*Model) error
	// Update(*Model) error
	// GetAll(Models, error)
	// GetByID(uint) (*Model, error)
	// Delete(uint) error
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

// Usado para crear un Producto
func (s *Service) Create(m *Model) error {
	m.CreatedAt = time.Now()
	return s.storage.Create(m)
}
