package main

import "fmt"

func main() {
	var n1, n2 int8
	n1 = 15
	n2 = 3
	r := suma(n1, n2)
	fmt.Println(r)

	var edad uint8
	edad = 17
	fmt.Println(tipoEdad(edad))

	n := []int8{52, 63, 47, 5, 5, 3, 7, 6, 100, 2, 47, -5}
	maximo, minimo := maxymin(n)
	fmt.Println("Máximo", maximo)
	fmt.Println("Mínimo", minimo)

}

// esta linea     \\\\\func suma(a int8, b int8) int8 //// se conoce como la firma de la función
// SHORTCUT de los parametros de la función
// Al lado de los parametros se coloca el tipo de dato de la función que se retornara
func suma(a, b int8) int8 {
	return a + b
}

func tipoEdad(edad uint8) string {
	var tipo string
	switch {
	case edad < 12:
		tipo = "niñez"
	case edad < 18:
		tipo = "Adolescencia"
	default:
		tipo = "Madurez"
	}
	return tipo
}

// RETORNO DE MULTIPLES variables
// maximo y minimo de numeros dados
// valores de retornos nombrados, se nombran cuando se coloca el tipo  de dato para las variables
// cuando las variables sean nombradas en la firma DE LA FUNCIÓN no tendran  que ser retornadas con RETURN
// tampoco la declaración de las variables
func maxymin(numeros []int8) (max int8, min int8) {
	// como es una función nombrada no es necesario declarar la siguiente linea
	/*
	  var max, min int8
	*/
	for _, v := range numeros {
		if v > max {
			max = v
		}
		if v < min {
			min = v
		}
	}
	// como es una función nombrada no es necesario declarar la siguiente linea de retorno con las variables, solo el retorno
	/*
	  return max, min
	*/
	return
}
