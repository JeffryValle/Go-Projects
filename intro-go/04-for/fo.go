package main

import "fmt"

// En go 'for' es la única estructura de repetición
// No existe While, Do While, forEach...

func main() {
	fmt.Println("*********** For loop ***********")

	i := 1 // asignamos una variable inicializadora
	for i <= 3 {
		fmt.Println(i)
		i = i + 1
	}
	// sintaxis más manejable
	for j := 0; j < 3; j++ {
		fmt.Println(j)
	}

	for i := range 3 {
		fmt.Println(i)
	}

	for {
		fmt.Println("loop")
		break
	}
	fmt.Println("****************************")
	for n := range 6 {
		if n%2 == 0 {
			continue
		}
		fmt.Println(n)
	}
}
