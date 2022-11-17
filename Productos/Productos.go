package Productos

import "time"

type Model struct {
	Id            uint
	Nombre        string
	Observaciones string
	Precio        float64
	CreatedAt     time.Time
	UpdatedAt     time.Time
}

type Models []*Model

type Storage interface {
	Migrate() error

	/*
		Create(*Model) error
		Update(*Model) error
		GetAll() (Models, error)
		GetById(uint) (*Model, error)
		Delete(uint) error
	*/
}

// servicio del producto
type Service struct {
	storage Storage
}

// retorna un servicio de puntero servicio
func NewService(s Storage) *Service {
	return &Service{s}
}

// es usado para migrar producto
func (s *Service) Migrate() error {
	return s.storage.Migrate()
}
