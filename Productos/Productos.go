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
	Create(*Model) error
	Update(*Model) error
	GetAll() (Models, error)
	GetById(uint) (*Model, error)
	Delete(uint) error
}
