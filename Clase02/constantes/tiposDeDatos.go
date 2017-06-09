package main

import "fmt"

func main() {
var a int
var b int8
a = 121212
b = 5
j:="Juan"
m:="Mogollón Martínez"
// casting
c := a + int(b)
fmt.Println(c)
fmt.Printf("hola %s %s que hace?\n", j, m)
fmt.Printf("c es de tipo: %T\n b es: %T\n", c, b)
// prioridad aritmetica
//    () % * / + -
d := 6 + 6 * 6 - 6
fmt.Println(d)

}
