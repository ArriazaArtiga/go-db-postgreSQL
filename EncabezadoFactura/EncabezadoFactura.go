package EncabezadoFactura

import "time"

type Model struct {
	Id        uint
	Cliente   string
	CreatedAt time.Time
	UpdatedAt time.Time
}
