package main

import "fmt"

func main() {
	saludarVarios(18, "Juan", "Gabriel", "Neyer")
}

func saludarVarios(edad uint8, nombres ...string) {
	fmt.Printf("%T\n", nombres)
	for _, v := range nombres {
		fmt.Println("Hola", v, "tienes", edad, "años")
	}

}
