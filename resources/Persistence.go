package resources

import (
	"log"

	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
)

func GetConexion() *gorm.DB {
	db, err := gorm.Open("postgres", "host=173.249.51.238 port=5432 user=postgres password=yeco2010 dbname=sch_seguridad")
	if err != nil {
		log.Fatal(err)
		panic("Error al abrir la conexion. ")
	}
	return db
}
