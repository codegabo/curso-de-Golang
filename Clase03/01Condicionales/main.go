package main

import "fmt"

func main() {
	// Operadores:
	// && and
	// || or
	// !  not NIEGA LA EXPRESION, O LA PASA A SU VALOR CONTRARIO
	// == iqual
	// <=
	// >=
	/*if 5 < 10 {
		fmt.Println("verdadero")
	}
	fmt.Println("fin del programa")*/
	/* isValid := true */
	/*edad := 30*/

	if edad, nombre := 10, "juan"; edad < 14 {
		fmt.Println(nombre, " es un niño")
	} else if edad < 18 {
		fmt.Println(nombre, " es un adolescente")
	} else if edad < 30 {
		fmt.Println(nombre, " es un adulto")
	} else {
		fmt.Println("adulto emtrando a su prevejez")
	}

	/*
		if isValid {
			fmt.Println("Esto se encuentra el un bloque verdadero")
			fmt.Println(isValid)
						// condiciones anidadas
						if 5 < 10 {
							fmt.Println("5 es menor que 10")
							} else {
								fmt.Println("5 no es menor que 10")
							}

		} else {
			fmt.Println("Esto se encuentra en un bloque falso")
		}*/
	fmt.Println("Fin del programa")
}
