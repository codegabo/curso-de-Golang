package main

import "fmt"

func main() {
	a := 3
	// imprimir la direccion en memoria del  valor de una variable
	// el valor cero de un puntero es 1000
	fmt.Println(&a)
	fmt.Println("Antes de duplicar, a =", a)
	duplicar(&a)
	fmt.Println("Después de duplicar, a =", a)
}

func duplicar(a *int) {
	*a *= 2
	fmt.Println("Dentro de la función duplicar a =", *a)
}
