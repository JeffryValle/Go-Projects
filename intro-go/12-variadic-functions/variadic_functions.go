package main

import (
	"fmt"
)

func main() {
	fmt.Println("******** Variadic Functions ********")

	sum(1, 2, 3, 4, 5, 6, 7, 8, 9)
	sum(82, 92, 93, 93, 98)

	nums := []int{1, 2, 3, 4}
	sum(nums...)
}

// función que toma una arbitraria cantidad de enteros
// como argumento
func sum(nums ...int) { // sum internamente se trata como un slice de int
	fmt.Println(nums, " ")
	total := 0

	for _, num := range nums { // _ ignora el indice y empieza desde num (que es el valor actual)
		total += num //
	}
	fmt.Println("Total:    ", total)
	fmt.Println("Promedio: ", total/len(nums))
}
