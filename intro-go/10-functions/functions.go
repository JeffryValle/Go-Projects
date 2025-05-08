package main

import "fmt"

func main() {
	fmt.Println("*********** Funciones *********** ")
	res := plus(25, 24)
	fmt.Println("Suma: ", res)

	sum := plusPlus(10, 20, 30)
	fmt.Println("Suma: ", sum)
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}
