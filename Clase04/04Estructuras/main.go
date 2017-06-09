package main

import "fmt"

//  Persona
// el oreden en el que se imprime aparecera segun el orden de la declaracion de los campos en la estructura
type Persona struct {
	Nombre string
	Edad   uint8
}

func main() {
	var persona1 Persona
	persona1.Nombre = "Juan"
	persona1.Edad = 20
	/*
	  fmt.Println(persona1)
	*/
	// para imprimir valores especificos se puede usar esto:
	fmt.Println(persona1.Nombre)

	//SHORTCUTS
	/*
		  persona2 := Persona{}
			persona2.Nombre = "Gabriel"
			persona2.Edad = 93
			fmt.Println(persona2)
	*/
	/*
		  persona2 := Persona{Nombre: "Neyer", Edad: 50}
			fmt.Println(persona2)
	*/
	//SHORTCUTS que muestra los valores de los campos en la posicion de los mismos sin necesidad de volver a declararlos
	/*
		  persona2 := Persona{"Neyer", 50}
			fmt.Println(persona2)
	*/
}
