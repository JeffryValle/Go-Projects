package main

import "fmt"

func main() {
	fmt.Println("******************************")

	var a = "initial"
	fmt.Println(a)

	var b, c int = 1, 2
	fmt.Println(b, c)

	// Go infiere en el tipo de variable
	var d = true
	fmt.Println(d)

	// Variables sin inicializar, Go les asigna el valor 0
	var e int
	fmt.Println(e)

	/* La sintaxis := es una abreviatura para inicialización y declaración de
	   una variable, esta sintaxis trabaja solamente dentro de funciones
	*/
	f := "apple"
	fmt.Println(f)
}
