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
