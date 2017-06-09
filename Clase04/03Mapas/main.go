package main

import "fmt"

func main() {
	//estructura basica de declaracion de un mapa
	/*idiomas := make(map[string]string)
	idiomas["es"] = "Español"
	idiomas["en"] = "Inglés"
	idiomas["fr"] = "Francés"

	fmt.Println(idiomas)
	*/

	//otra forma de declaración de un mapa asignando directamente un valor
	/*
		  idiomas := map[string]string{
				"es": "Español",
				"en": "Inglés",
				"fr": "francés",
				"pt": "Portugués",
			}
	*/

	//con esta impresion se muestra solo una posicion
	/*fmt.Println(idiomas["es"])*/
	/*fmt.Println(idiomas)*/

	//Eliiminar posiciones de un mapa
	//delete(variable, "posición")
	/*
		  delete(idiomas, "en")
			fmt.Println(idiomas)
	*/

	//para verificar si las posiciones existen se usa lo siguiente:
	//si la posición existe retonara un valor booleano,
	//Mapa retorna o devuelve dos valores y por ello se deben colocar dos variables, se puede colocar una variable cualquiera, se suele colocar "ok"
	/*if idioma, ok := idiomas["pt"]; ok {
		fmt.Println("la posición pt SI existe y representa a", idioma)
	} else {
		fmt.Println("la posición pt NO existe", idioma)
	}*/

	numeros := map[int]string{
		1:    "Uno es un número pequño\n",
		2:    "Dos es mayor que uno\n",
		2016: "En este caso este número entero 2016 no se ve afectado ya que no es obligatorio llevar el orden númerico de los enteros\n",
		-4:   "Los negativos tampoco se ven a fectados para mostrarse en mapas\n",
	}
	fmt.Println(numeros)
}
