package main

import (
	"fmt"
)

func main() {
	fmt.Println("****** Interfaces ******")
	// var w Writer = ConsoleWriter{}
	// w.Write([]byte("Hello Go"))
	// myInt := IntCounter(0)
	// var inc Incrementer = &myInt
	// for i := 0; i < 10; i++ {
	// 	fmt.Println(inc.Increment())
	// }

	var i interface{} = "0"
	switch i.(type) {
	case int:
		fmt.Println("i is an integer")
	case string:
		fmt.Println("i is a string")
	default:
		fmt.Println("I don't know what i is ")
	}
}

/*
Las interfaces en Go son una herramienta
fundamental para escribir código flexible, desacoplado y reutilizable
Una interfaz en Go es un tipo que define un conjunto de métodos, p
ero no implementa ninguno. Cualquier tipo (como un struct) que implemente
todos esos métodos, automáticamente satisface esa interfaz.
*/
// type Writer interface {
// 	Write([]byte) (int, error)
// }
// type ConsoleWriter struct{}

//	func (cw ConsoleWriter) Write(data []byte) (int, error) {
//		n, err := fmt.Println(string(data))
//		return n, err
//	}
// type IntCounter int

// func (ic *IntCounter) Increment() int {
// 	*ic++
// 	return int(*ic)
// }

// type Incrementer interface {
// 	Increment() int
// }
