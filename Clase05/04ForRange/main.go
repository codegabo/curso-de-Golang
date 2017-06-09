package main

import "fmt"

func main() {
	numeros := []string{"cero", "uno", "dos", "tres"}
	// i = indice
	/*
	   for i, v := range numeros{}
	*/
	// _ = omicion de variables
	for _, v := range numeros {
		fmt.Println(v)
	}

	//for range con mapas
	nombres := map[string]string{"es": "españa", "co": "colombia", "br": "Brasil"}

	for i, v := range nombres {
		fmt.Println(i, v)
	}
	// for range con strings
	frase := "Hola Mundo"
	for posición, valor := range frase {
		// el casting string(valor) en este caso se usa para convertir los valores bit a string
		fmt.Println(posición, string(valor))
	}
	// SHOTCHUTS
	// range con SLICE
	for _, entero := range []int{15, 36, 24, 85} {
		fmt.Println(entero)
	}
	// Range con String
	for _, textoChort := range "GOLANG" {
		fmt.Println(string(textoChort))
	}
}
