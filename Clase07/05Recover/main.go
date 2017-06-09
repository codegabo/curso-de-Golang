package main

import "fmt"

func main() {
	f()
}

func f() {
	// función anónima para respaldar o soportar un panic o error con comportamientos de respaldo-soporte
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("%T\n", r)
			fmt.Println("Recuperado en f:", r)
		}
	}()
	fmt.Println("Llamando a g.")
	g(5)
	fmt.Println("Retorna normalment desde g")
}

func g(i int) {
	if i > 3 {
		fmt.Println("AAAAAAAAAiudaaa :v")
		panic("EL número no puede ser mayor que 3")
	}
	defer fmt.Println("Defer en la función g")
	fmt.Println("Imprimiendo en g", i)
}
