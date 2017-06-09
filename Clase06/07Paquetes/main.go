package main

import (
	"fmt"

	"github.com/golang-es/edcurso-go/Clase06/07Paquetes/despedida"
	"github.com/golang-es/edcurso-go/Clase06/07Paquetes/saludar"
)

func main() {
	nombre := "juan gabriel"
	saludar.Saludar(nombre)
	saludar.MeVes = "valor de variable externa asignado desde main"
	fmt.Println(saludar.MeVes)

	despedida.Despedirse(nombre)
}
