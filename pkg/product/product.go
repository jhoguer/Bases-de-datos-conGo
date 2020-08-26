package product

import (
	"time"
)

// Model of product
type Model struct {
	ID           uint
	Name         string
	Observations string
	Price        int
	CreatedAt    time.Time
	UpdatedAt    time.Time
}

// Model slice of Model
type Models []*Model

type Storage interface {
	Migrate() error
	Create(*Model) error
	// Update(*Model) error
	// GetAll() (Models, error)
	// GetByID(uint) (*Model, error)
	// Delete(uint) error
}

// Service of Product
type Service struct {
	storage Storage
}

// NewService return a pointer of Service
func NewService(s Storage) *Service {
	return &Service{s}
}

// Migrate is used for migrate product
func (s *Service) Migrated() error {
	return s.storage.Migrate()
}

// Migrate is used for create a product
func (s *Service) Create(m *Model) error {
	m.CreatedAt = time.Now()
	err := s.storage.Create(m)
	if err != nil {
		return err
	}

	return nil
}
