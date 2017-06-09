package main

import (
	"errors"
	"fmt"
)

// uso normal
/*
func main() {
	//valor cero del tipo de daro error es nil
	err := errors.New("Mensaje de error")
	fmt.Printf("%T\n", err)
}
*/
func main() {
	//valor cero del tipo de daro error es nil
	resultado, err := division(100, 0)
	if err != nil {
		fmt.Println("Error:", err)
		// con el return no continuara ejecutando la sentencia
		return
	}
	fmt.Println(resultado, err)
}

//
func division(dividendo, divisor float64) (resultadoFinal float64, err error) {
	if divisor == 0 {
		err = errors.New("No se puee divividir por cero")
	} else {
		resultadoFinal = dividendo / divisor
	}
	return
}
