package main

import "fmt"

func main() {
	//SLICES
	//1. Apuntador a un array
	//2. Tamaño
	//3. Capacidad

	//estructura basica para la creacion de SLICES
	/*
	     var nombres []string
	   	fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	   	nombres = append(nombres, "juan")
	   	fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	   	nombres = append(nombres, "gabriel")
	   	fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	   	nombres = append(nombres, "neyer")
	   	fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))

	   	//para editar o asignar un valor a un indice-campo-valor se realiza de la siguiente manera
	   	nombres[2] = "Azucena"
	   	fmt.Println(nombres)
	*/
	//otra estructura para la creacion de SLICES cuando ya se saben cuales son los valores a asignar en la variable declarada
	nombres := []string{
		"Juan",
		"Gabriel",
		"Neyer",
	}
	fmt.Println(nombres)

	//estructura de SLICES con la funcion preconstrudia MAKE()
	//make(tipo de dato del slice, tamaño inicial, capacidad inicialOpcional)
	//generalmente se coloca un tamañpo inicial de cero por que generlamente no se sabe que cantidad de información de se va a agregar
	/*
		  nombres := make([]string, 0)

			fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
			nombres = append(nombres, "juan")
			fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
			nombres = append(nombres, "gabriel")
			fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
			nombres = append(nombres, "neyer")
			fmt.Printf("su tamaño es: %d y su capacidad es: %d\n", len(nombres), cap(nombres))
	*/
}
