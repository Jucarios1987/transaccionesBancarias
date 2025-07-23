package main

import (
	/* En esta zona importamos los paquetes que utilizaremos  */

	"net/http"

	"github.com/gorilla/mux"

	/* Ubicacion en le cual se encuentra la carpeta routes que contiene el archivo "index.routes.go" en el cual se definio la funcion "HomeHandler" */
	"github.com/Jucarios1987/transaccionesBancarias/routes"
)

func main() {
	// Variable 'r' para crear un nuevo enrutador HTTP.
	r := mux.NewRouter()

	// Inicializamos un router
	r.HandleFunc("/", routes.HomeHandler)

	// Inicializamos el ervidor
	http.ListenAndServe(":3000", r)

}
