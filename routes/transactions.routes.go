package routes

import (
	"encoding/json"
	"net/http"

	"github.com/Jucarios1987/transaccionesBancarias/db"
	"github.com/Jucarios1987/transaccionesBancarias/models"
)

/* Funcion para obtener las cuentas */
func GetTransactionsHandler(w http.ResponseWriter, r *http.Request) {
	var transactions []models.Transactions
	db.DB.Find(&transactions)
	json.NewEncoder(w).Encode(&transactions)
}

/* Funcion para obtener una cuenta */
func GetTransactionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("get transaction"))
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
	}

	json.NewEncoder(w).Encode(&transaction)
}

/* Funcion para eliminar una cuenta*/
func DeleteTransactionHandler(w http.ResponseWriter, r *http.Request) {
	w.Write([]byte("delete"))
}
