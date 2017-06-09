package main

func main() {

	// caso de switch básico
	/*var a int
	a = 5
	switch a {
	case 1:
		fmt.Println("El valor es 1")
	case 2:
		fmt.Println("El valor es 1")
	default:
		fmt.Println("No es ni 1 ni 2")
	}*/

	// Switch que me diga que diga que dia semana es de acuerdo al valor númerico de la variable
	/*
		var a uint8
		a = 5
		switch a {
		case 1:
			fmt.Println("Lunes")
		case 2:
			fmt.Println("Martes")
		case 3:
			fmt.Println("Miercoles")
		case 4:
			fmt.Println("Jueves")
		case 5:
			fmt.Println("Viernes")
		case 6:
			fmt.Println("Sábado")
		case 7:
			fmt.Println("Domingo")
		default:
			fmt.Println("No es un dia de la semana")
		}
	*/

	// Switch que me diga que diga si el valor númerico de la variable se encuentra en el fin de la semana
	/*
		  var a uint8
			a = 6
			switch a {
			case 1:
				fallthrough
			case 2:
				fallthrough
			case 3:
				fallthrough
			case 4:
				fallthrough
			case 5:
				fmt.Println("Estás entre semana")
			case 6:
				fallthrough
			case 7:
				fmt.Println("Estás en fin de semana, disfruta! :D")
			default:
				fmt.Println("Este no es un día de la semana valido")
			}*/

	// Switch que me diga que diga si el valor númerico de la variable se encuentra en el fin de la semana Recortando variables y haciendo un codigo mas limpio, declarando y asignando las variables y expresiones junto a la palabra reservada case
	/*
		switch a := 3; {
		case a >= 0 && a <= 5:
			fmt.Println("Estás entre semana")
		case a >= 6 && a <= 7:
			fmt.Println("Estás en fin de semana, disfruta!, :D")
		default:
			fmt.Println("Tu valor no es un día valido")

		}
	*/
}
