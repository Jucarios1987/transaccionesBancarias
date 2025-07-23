package routes

/* Importamos el modulo HTTP */
import "net/http"

// HomeHandler: Funcion que maneja solicitudes HTTP dirigidas a la ruta principal ("/") de una aplicación web.
// 'writer': Indica como se respondera el cliente
// 'reader': Permite acceder a los parametros que envia el cliente
func HomeHandler(writer http.ResponseWriter, reader *http.Request) {
	/* Retornamos un objeto al cliente: En este caso retornamos un arreglo de byte con un texto */
	writer.Write([]byte("Hello Transacciones Bancarias Air"))
}
