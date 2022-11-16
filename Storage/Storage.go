package Storage

// para crear el metodo singlenton

import (
	"database/sql"
	"fmt"
	"log"
	"sync"

	_ "github.com/lib/pq"
)

var (
	db   *sql.DB
	once sync.Once
)

func NewPostgresDB() {
	once.Do(func() {
		var err error
		cadenaConexion := "postgres://userPg:123456789@localhost:5432/dbTEst?sslmode=disable"
		db, err = sql.Open("postgres", cadenaConexion)
		if err != nil {
			log.Fatalf("No se puede abrir la base de datos, erro: %v", err)
		}
		if err = db.Ping(); err != nil {
			log.Fatalf("No se pudo hacer ping, erro: %v", err)
		}
		fmt.Println("Conectado a Postgres")
	})

}

// retorna una unica instancia de DB
func Pool() *sql.DB {
	return db
}
