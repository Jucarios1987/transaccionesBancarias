/* Indicamos que este paquete esta  dentro de routes */
package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Jucarios1987/transaccionesBancarias/db"
	"github.com/Jucarios1987/transaccionesBancarias/models"
	"github.com/gorilla/mux"
)

/* Funcion para obtener las cuentas */
func GetAccountsHandler(w http.ResponseWriter, r *http.Request) {
	var accounts []models.Accounts
	db.DB.Find(&accounts)
	json.NewEncoder(w).Encode(&accounts)
}

/* Funcion para obtener una cuenta */
func GetAccountHandler(w http.ResponseWriter, r *http.Request) {

	var account models.Accounts
	params := mux.Vars(r)

	db.DB.Find(&account, params["id"])

	if account.Id == 0 {
		/* Establecemos un codigo de estado para la peticion HTTP */
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Account not found"))
		return
	}

	json.NewEncoder(w).Encode(&account)
}

/* Funcion para crear una cuentra */
func PostAccountHandler(w http.ResponseWriter, r *http.Request) {

	var account models.Accounts
	json.NewDecoder(r.Body).Decode(&account)

	// Aqui adicionamos un nuevo registro en la tabla accounts y almacenamos el resultado en la variable "createdAccount"
	createdAccount := db.DB.Create(&account)

	// Validamos si tenemos un error
	err := createdAccount.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
	}

	json.NewEncoder(w).Encode(&account)

}

/* Funcion para eliminar una cuenta*/
func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	var account models.Accounts
	/* Extraemos las variables que bienen desde mux */
	// mux.Vars(r): Obtenemos lo que biene del request de la solicitud HTTP
	params := mux.Vars(r)
	// Buscamos una cuenta y la almacenamos en "account"
	db.DB.First(&account, params["id"])

	// Validamos si la cuenta existe antes de eliminar
	if account.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Accoun not foun"))

		return
	}

	// Esta instruccion solo actializa la fecha de eliminacion en la DB lo que hace que el registro no se tenga en cuenta en busquedas.
	// db.DB.Delete(&account)

	// Esta instruccion borra el dato de la base dedatos
	db.DB.Unscoped().Delete(&account)
	w.WriteHeader(http.StatusOK)

	// w.Write([]byte("delete"))
}
