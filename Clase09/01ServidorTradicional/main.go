// SERVIDOR DE ARCHIVOS ESTATICOS

package main

import (
	"log"
	"net/http"
)

func main() {
	// Conexión tradicional
	/*
		// ruta al servidor, y que debe hacer el servidor cuando se haga la petición
		http.Handle("/", http.FileServer(http.Dir("public")))
		// carga del servidor en el puerto
		log.Println("Ejecutando server en http://localhost:8080")
		log.Println(http.ListenAndServe(":8080", nil))
	*/

	// Conexión con un multiplexor, o como se denomina en golang MUX
	//Un miltiplexor o Mux es una herramienta que permite recibir y procesar varias herramientas al tiempo
	// El mux esta diseñado para hacer el programa mas eficiente
	mux := http.NewServeMux()
	fs := http.FileServer(http.Dir("public"))
	mux.Handle("/", fs)
	log.Println("Ejecutando server en http://localhost:8080")
	log.Println(http.ListenAndServe(":8080", mux))
}
