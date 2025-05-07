package main

import "fmt"

func main() {
	fmt.Println("********** if / else statement **********")
	// No se necesitan los parentesis en Go

	if 7%2 == 0 {
		fmt.Println("7 is even") // es par
	} else {
		fmt.Println("7 is odd") // es impar
	}

	if 8%4 == 0 {
		fmt.Println("8 es divisible por 4")
	}

	if 8%2 == 0 && 7%2 == 0 {
		fmt.Println("either 8 or 7 are even")
	}

	if num := 9; num < 0 {
		fmt.Println(num, "is negative")
	} else if num < 10 {
		fmt.Println(num, "has 1 digit")
	} else {
		fmt.Println(num, "has multiple digits")
	}

	// No existe el operador ternario en Go :')
}
