package main

import "fmt"

/*
func saludar() {
	fmt.Println("Hola Mundo")
}

func main() {
	// el punto de entrada siempre sera el main
	saludar()
}
*/

// FUNCIONES CON PARAMETROS
/*
func main() {
	// el punto de entrada siempre sera el main
	saludar("Gabo", 20)
}
func saludar(nombre string, edad uint8) {
	/*
	  fmt.Println("Hola", nombre, "tienes", edad, "años")
	*-/
	fmt.Printf("Hola %s tienes %d años de edad", nombre, edad)
}
*/

// FUNCIONES CON RETURN
func main() {
	// el punto de entrada siempre sera el main
	saludar("Gabo", 20)
}
func saludar(nombre string, edad uint8) {
	/*
	  fmt.Println("Hola", nombre, "tienes", edad, "años")
	*/
	fmt.Printf("Hola %s tienes %d años de edad", nombre, edad)
	if edad > 20 {
		return
	}
	fmt.Println("\nEres jovén")
}
