package DetalleFactura

import "time"

type Model struct {
	Id        uint
	Factura   uint
	Producto  uint
	CreatedAt time.Time
	UpdatedAt time.Time
}
