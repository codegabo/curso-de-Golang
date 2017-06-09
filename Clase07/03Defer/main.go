package main

import "fmt"

func main() {
	// los defer se ejecutan en orden inverso
	fmt.Println("Conectando a la base de datos...")
	//supuestas operaciones de la conexión
	defer cerrarBD()

	fmt.Println("Consultamos información, set de datos")
	defer cerrarSetDeDatos()
}

func cerrarBD() {
	fmt.Println("Cerrar la base de datos")
}

func cerrarSetDeDatos() {
	fmt.Println("Cerrar el set de datos")
}
