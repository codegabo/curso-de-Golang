package main

import "fmt"

func main() {
	//en la sentencia de inicio se coloca la variable que va a controlar los ciclos
	/*
		  for i := 0; i < 5; i++ {
				fmt.Println(i)
			}
	*/

	// la variable i no se ejecutara cuando su valor sea 3 y continuara con la iteración
	/*
		for i := 4; i >= 0; i-- {
			if i == 3 {
				continue
			}
			fmt.Println(i)
		}
	*/

	// la iteración se va a romper o detener cuando el valor de la variable i sea 2, esto gracias a la palabra reservada break
	/*
		  for i := 0; i < 5; i++ {
				if i == 2 {
					fmt.Println("acá la variable i vale 2, por ende se rompe el ciclo")
					break
				}
				fmt.Println(i)
			}
	*/

	// for para llenar un array de arrays, osea un array multidimencional

	matriz := [3][3]int{}
	valor := 0
	//este for recorrera las filas
	for i := 0; i < 3; i++ {
		// for anidado
		//este for recorrera las columnas
		for j := 0; j < 3; j++ {
			valor++
			matriz[i][j] = valor
		}
	}
	/*
		fmt.Println(matriz)
	*/
	for fila := 0; fila < 3; fila++ {
		for columna := 0; columna < 3; columna++ {
			fmt.Println(matriz[fila][columna])
		}
	}
}
