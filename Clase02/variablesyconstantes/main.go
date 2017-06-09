package main

import "fmt"

func main() {
	//para declarar una variable
	//1. Siempre colocar la palabra reservada var.
	//2. El nombre de la variable.
	//3.Tipo de dato
	//SE PUEDE DECLARA MAS DE UNA VARIABLE,
	//LAS CUALES CADA UNA VA A SER UN ESPACIO INDEPENDIENTE DE MEMORIA, O SEA QUE NO ESTAN ASOCIADAS ENTRE SI.
	var nombre, apellido string
	//acá se asigna el valor de la variable, en este caso, el valor para la variable "nombre"
	nombre = "Juan Gabriel"
	apellido = "Mogollón Martínez"
	fmt.Println(nombre, apellido)
	/*tambien se pueden declara las variables de esta manera
	nombre, apellido := "Juan Gabriel", "Mogollón Martínez"
	nombre, edad := "pedro", 20
	*/

}
