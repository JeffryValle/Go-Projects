package main

import "fmt"

func main() {
	fmt.Println("**** Multiple Return Values ****")

	a, b := vals()
	fmt.Println("Primer digito: ", a)
	fmt.Println("Segundo digito: ", b)

	c, _ := vals()
	fmt.Println("Tercer digito: ", c) // si solo ocupáramos uno de los digitos, utilizamos el identificador blank _

}

func vals() (int, int) { // funciones que devuelve 2 enteros
	return 3, 7
}
