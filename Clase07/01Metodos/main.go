package main

import "fmt"

// estructura persona
type Persona struct {
	nombre string
	edad   int8
}

// METODO
// para asociar funciones a estructuras la firma de debe ser de la siguiente manera
func (p Persona) saludar() {
	fmt.Println("Hola, soy una persona")
}

type Numero int

func (n Numero) presentarse() {
	fmt.Println("Hola, yo soy un número")
}

func (p *Persona) asignarEdad(i int8) {
	if i >= 0 {
		p.edad = i
	} else {
		fmt.Println("No es valida una edad negativa")
	}
}

func main() {
	// se llama al metodo
	var persona Persona
	persona.saludar()

	var numero Numero
	numero.presentarse()

	// para edad negativa
	/*
	  persona.asignarEdad(-35)
	*/
	persona.asignarEdad(35)
	fmt.Println(persona)
}
