package main

import "fmt"

func main() {
	// Esta función tiene () despés de } para que se pueda autoejecutar, esto cuando la funcion no tiene nombre
	func() {
		fmt.Println("Me imprimo")
	}()

	// Función declarada dentro de una variable
	anonima := func() {
		fmt.Println("Me imprimo dentro de una variable llamada anonima")
	}
	anonima()

	// se declara cuna variable que tiene como valor una función
	secuencia1 := secuencia()
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println(secuencia1())
	fmt.Println("___________")
	secuencia2 := secuencia()
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())
	fmt.Println(secuencia2())

}

// funciones que retornan funciones
func secuencia() func() int32 {
	var x int32
	return func() int32 {
		x++
		return x
	}
}
