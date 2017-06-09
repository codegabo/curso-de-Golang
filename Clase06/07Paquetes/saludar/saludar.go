package saludar

import "fmt"

// esta variable es para ser usada desde otro paquete
var MeVes string

// esta variable solo puede ser usada dentro de este paquete ya que no iinicia con mayuscula
var noMeVes string

// la "S" mayuscula quiere decir que sera una función exportada
// si el nombre de la función no comienza con mayuscula entonces solo se ppodra usar en este archivo
func Saludar(nombre string) {
	fmt.Println("hola", nombre)
}

func noVisible()
