package Storage

import (
	"database/sql"
	"fmt"
	"time"
)

const (
	psqlCreateProduct = `CREATE TABLE IF NOT EXISTS products(
		Id SERIAL NOT NULL,
		Nombre VARCHAR(25) NOT NULL,
		Observaciones VARCHAR(100),
		Precio MONEY NOT NULL,
		Created_at TIMESTAMP NOT NULL DEFAULT now(), 
		Updated_at TIMESTAMP,
		CONSTRAINT products_id_pk PRIMARY KEY(Id)
	)`
)

// PsqlProduct se usa para trabajar postgres - product
type PsqlProduct struct {
	db *sql.DB
}

// NewPsqlProduct retorna un nuevo puntero de PsqlProduct
func NewPsqlProduct(db *sql.DB) *PsqlProduct {
	return &PsqlProduct{db}

}

func (p *PsqlProduct) Migrate() error {
	t := time.Now()
	fecha := fmt.Sprintf("%02d-%02d-%d, a las %02d:%02d:%02d", t.Day(), t.Month(), t.Year(),
		t.Hour(), t.Minute(), t.Second())
	stmt, err := p.db.Prepare(psqlCreateProduct)

	if err != nil {
		return err
	}
	defer stmt.Close()
	_, err = stmt.Exec()

	if err != nil {
		return err
	}
	fmt.Println("Migración de producto ejecutada exitosamente el ", fecha)
	return nil

}
