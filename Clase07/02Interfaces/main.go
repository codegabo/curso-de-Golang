package main

import "github.com/golang-es/edcurso-go/Clase07/02Interfaces/animales"

func main() {
	var p animales.Perro
	var g animales.Gato
	p.Nombre = "Firulais"
	g.Nombre = "Garfield"

	/*
		  AdoptarPerro(p)
			AdoptarGato(g)
	*/
	AdoptarMascota(p)
	AdoptarMascota(g)
}

// de esta manera se llaman metodos externos
/*
func AdoptarPerro(p animales.Perro) {
	p.Comunicarse()
}

func AdoptarGato(g animales.Gato) {
	g.Comunicarse()
}
*/

// usando interfaces para el llamado de metodos externos

func AdoptarMascota(m animales.Mascota) {
	m.Comunicarse()
}
