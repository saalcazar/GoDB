package invoiceitem

import "time"

//Model of invoiceitem
type Model struct {
	ID               uint
	InvoiceHeader_ID uint
	ProductID        uint
	CreatedAt        time.Time
	UpdatedAt        time.Time
}

type Storage interface {
	Migrate() error
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
