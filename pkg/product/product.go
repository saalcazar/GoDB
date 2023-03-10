package product

import (
	"fmt"
	"strings"
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

// Implementa la interface Stringer del paquete format - al imprimir la información esta este tabulada y visualmente legible
func (m *Model) String() string {
	return fmt.Sprintf(
		"%02d | %-20s | %-20s | %5d | %10s | %10s",
		m.ID, m.Name, m.Observation, m.Price,
		m.CreatedAt.Format("2006-01-02"), m.UpdatedAt.Format("2006-01-02"))
}

// Slice de Modelo
type Models []*Model

func (m Models) String() string {
	builder := strings.Builder{}
	builder.WriteString(fmt.Sprintf("%02s | %-20s | %-60s | %5s | %10s | %10s\n", "id", "name", "observations", "price", "created_at", "updated_at"))
	for _, model := range m {
		builder.WriteString(model.String() + "\n")
	}
	return builder.String()
}

type Storage interface {
	Migrate() error
	Create(*Model) error
	// Update(*Model) error
	GetAll() (Models, error)
	GetByID(uint) (*Model, error)
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

// GetAll se usa para leer todos los datos de una tabla
func (s *Service) GetAll() (Models, error) {
	return s.storage.GetAll()
}

func (s *Service) GetById(id uint) (*Model, error) {
	return s.storage.GetByID(id)
}
