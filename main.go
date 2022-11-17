package main

//para importar desde las interfaces a traves del modulo
/*
	"github.com/ArriazaArtiga/go-db-postgreSQL/Productos"
	"github.com/ArriazaArtiga/go-db-postgreSQL/DetalleFacturas"
	"github.com/ArriazaArtiga/go-db-postgreSQL/EncabezadoFacturas"
*/
import (
	"log"

	"github.com/ArriazaArtiga/go-db-postgreSQL/Productos"
	"github.com/ArriazaArtiga/go-db-postgreSQL/Storage"
)

func main() {
	Storage.NewPostgresDB()
	storageProducto := Storage.NewPsqlProduct(Storage.Pool())
	ServicioProducto := Productos.NewService(storageProducto)

	if err := ServicioProducto.Migrate(); err != nil {
		log.Fatalf("Producto.Migrate: %v", err)
	}
}

//Para subir nuestro proyecto a github

/*Es necesario en la terminal ejecutar los siguientes comandos:
- git init    // para inicializar
- git add .   // para agregar todas las carpetas y archivos
- git commit  // para agregar un mensaje de los cambios
- Se crea el nuevo repositorio
- git remote add origin https://github.com/ArriazaArtiga/go-db-postgreSQL.git
- git branch -M main
- git push -u origin main
*/

/*
Para crear un nuevo modulo:
	- go mod init https://github.com/ArriazaArtiga/go-db-postgreSQL.git   // la url raiz del repositorio

*/
