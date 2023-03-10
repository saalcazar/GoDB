package invoice

import (
	"github.com/saalcazar/GoDB/pkg/invoiceheader"
	"github.com/saalcazar/GoDB/pkg/invoiceitem"
)

type Model struct {
	Header *invoiceheader.Model
	Items  invoiceitem.Models
}

// Implementa el almacenamiento de base de datos
type Storage interface {
	Create(*Model) error
}

// Servicio de Invoice
type Service struct {
	storage Storage
}

// retorna un puntero de servicio
func NewService(s Storage) *Service {
	return &Service{s}
}

// Crea una nueva factura
func (s *Service) Create(m *Model) error {
	return s.storage.Create(m)
}
