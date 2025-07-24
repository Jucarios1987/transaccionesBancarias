/* Indicamos que este paquete esta  dentro de routes */
package db

import (
	"log"

	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

/* Variables globales, se pueden utilizar desde otros contextos */
// Variable data base string

/* Conexion Local */
var DSN = "host=localhost user=postgres password=1087989233 dbname=cuentasBancarias port=5432 sslmode=disable"

/* Conexion Docker */
// var DSN = "host=localhost user=postgres password=1087989233 dbname=cuentasbancarias port=5433 sslmode=disable TimeZone=UTC client_encoding=UTF8"

// Variable que retorna informacion de la base de datos
var DB *gorm.DB

// Si definimos una funcion y queremos que se pueda utilizar en otros paquetes, el nombre debe iniciar en mayuscula.
func DBConnection() {
	/* Variables locales solo se pueden utilizar dentro de esta funcion */
	// Variable que me retorna un error de conexion a una base de datos
	var error error

	// postgres: Driver que utilizaremos
	DB, error = gorm.Open(postgres.Open(DSN), &gorm.Config{})

	// Cuando error retorna algo ingresa por la condicion
	if error != nil {
		// Muestra el error por consola
		log.Fatal(error)
	} else {
		log.Println("DB Connected")
	}
}
