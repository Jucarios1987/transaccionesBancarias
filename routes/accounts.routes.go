/* Indicamos que este paquete esta  dentro de routes */
package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Jucarios1987/transaccionesBancarias/db"
	"github.com/Jucarios1987/transaccionesBancarias/models"
)

/* Funcion para obtener las cuentas */
func GetAccountsHandler(w http.ResponseWriter, r *http.Request) {
	var accounts []models.Accounts
	db.DB.Find(&accounts)
	json.NewEncoder(w).Encode(&accounts)
}

/* Funcion para obtener una cuenta */
func GetAccountHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get account"))
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

	//w.Write([]byte("post"))
}

/* Funcion para eliminar una cuenta*/
func DeleteAccountHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}
