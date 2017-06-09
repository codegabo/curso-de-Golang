package main

import "fmt"

func main() {
	//Arrays

	//estructura basica de array
	/*
		var nombres [3]string
		nombres[0] = "Juan"
		nombres[1] = "Gabriel"
		nombres[2] = "Neyer"

		fmt.Println(nombres)
	*/

	//estructura de array con chortcuts
	nombres := [3]string{"Juan", "Gabriel", "Neyer"}
	//declaración de variable donde su valor es un solo campor de la variable madre, esto con el propósito de crear una jerarquia o dar nombre a valores de la variable madre
	nombre := nombres[1]
	size := len(nombres)
	fmt.Println("El tamaño del arreglo es de:", size)
	fmt.Println(nombre)
	//rango de valores de una variable
	/*fmt.Println(inicio:finalExcluyente)*/
	fmt.Println(nombres[0:2])

	/*fmt.Println(nombres)*/
	//con esta imprecion solo se llama el campo que nosotros desemos a través de la posición
	/*fmt.Println(nombres[1])*/

}
