package main

import (
	/* En esta zona importamos los paquetes que utilizaremos  */

	"net/http"

	"github.com/gorilla/mux"

	/* Ubicacion en le cual se encuentra la carpeta routes que contiene el archivo "index.routes.go" en el cual se definio la funcion "HomeHandler" */
	"github.com/Jucarios1987/transaccionesBancarias/db"
	"github.com/Jucarios1987/transaccionesBancarias/models"
	"github.com/Jucarios1987/transaccionesBancarias/routes"
)

func main() {

	// Iniciamos la coneccion a la base de datos.
	db.DBConnection()

	// Migracion a la DB
	db.DB.AutoMigrate(models.Accounts{})
	db.DB.AutoMigrate(models.Transactions{})

	// Variable 'r' para crear un nuevo enrutador HTTP.
	r := mux.NewRouter()

	/* Inicializamos un router */
	r.HandleFunc("/", routes.HomeHandler)

	// Establecemos las rutas para las peticiones http de las cuentas
	r.HandleFunc("/accounts", routes.GetAccountsHandler).Methods("GET")
	r.HandleFunc("/accounts/{id}", routes.GetAccountHandler).Methods("GET")
	r.HandleFunc("/accounts", routes.PostAccountHandler).Methods("POST")
	r.HandleFunc("/accounts", routes.DeleteAccountHandler).Methods("DELETE")

	// Establecemos las rutas para las peticiones http de las transacciones
	r.HandleFunc("/transactions", routes.GetTransactionsHandler).Methods("GET")
	r.HandleFunc("/transactions/{id}", routes.GetTransactionHandler).Methods("GET")
	r.HandleFunc("/transactions", routes.PostTransactionHandler).Methods("POST")
	r.HandleFunc("/transactions", routes.DeleteTransactionHandler).Methods("DELETE")

	// Inicializamos el ervidor
	http.ListenAndServe(":3000", r)

}
