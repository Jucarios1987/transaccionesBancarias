package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Jucarios1987/transaccionesBancarias/db"
	"github.com/Jucarios1987/transaccionesBancarias/models"
	"github.com/gorilla/mux"
)

/* Funcion para obtener las cuentas */
func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	var transactions []models.Transactions
	db.DB.Find(&transactions)
	json.NewEncoder(w).Encode(&transactions)
}

/* Funcion para obtener una cuenta */
func GetTransactionHandler(w http.ResponseWriter, r *http.Request) {

	var transaction models.Transactions
	params := mux.Vars(r)
	//fmt.Println(params)

	db.DB.Find(&transaction, params["id"])

	if transaction.Id == 0 {
		/* Establecemos un codigo de estado para la peticion HTTP */
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Transaction not found"))
		return
	}

	json.NewEncoder(w).Encode(&transaction)
}

/* Funcion para crear una cuentra */
func PostTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transactions
	json.NewDecoder(r.Body).Decode(&transaction)

	createTransaction := db.DB.Create(&transaction)
	err := createTransaction.Error
	if err != nil {
		w.WriteHeader(http.StatusBadRequest)
		w.Write([]byte(err.Error()))
		return
	}

	json.NewEncoder(w).Encode(&transaction)
}

/* Funcion para eliminar una cuenta*/
func DeleteTransactionHandler(w http.ResponseWriter, r *http.Request) {
	var transaction models.Transactions
	params := mux.Vars(r)
	db.DB.First(&transaction, params["id"])

	if transaction.Id == 0 {
		w.WriteHeader(http.StatusNotFound)
		w.Write([]byte("Accoun not foun"))

		return
	}

	db.DB.Unscoped().Delete(&transaction)
	w.WriteHeader(http.StatusNoContent)
}
