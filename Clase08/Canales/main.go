package main

import "fmt"

func main() {
	bodegaOrigen := []string{"php", "javascript", "golang", "html", "java", "git"}

	bodegaDestino := []string{}

	fotocopiadora := make(chan string)
	fotocopiado := make(chan string)

	// go rutina que deja los libros en la bodega destino

	go func() {
		for _, libro := range bodegaOrigen {
			fotocopiadora <- libro
		}
	}()

	//  go rutina que se encargue de realizar copia de los libros
	go func() {
		for {
			// acá fotocopeadora le esta entregando su valos a la variable libro
			libro := <-fotocopiadora
			fotocopiado <- libro

			// agregar el libro a la bodega destino
			bodegaDestino = append(bodegaDestino, libro)

			if len(bodegaOrigen) == len(bodegaDestino) {
				// cerrar el canal fotocopiado
				close(fotocopiado)
			}
		}
	}()

	fmt.Println("** Listado de libros fotocopiados**")
	for {
		if libro, ok := <-fotocopiado; ok {
			fmt.Println(libro)
		} else {
			break
		}
	}
}
