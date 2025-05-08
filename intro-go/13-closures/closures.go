package main

import "fmt"

func main() {
	fmt.Println("************ Closures ************ ")

	nextInt := intSeq()

	fmt.Println(nextInt())
	fmt.Println(nextInt())
	fmt.Println(nextInt())

	newInts := intSeq()
	fmt.Println(newInts())
}

// función que retorna otra función la cual definimos en el
// cuerpo de la función intSeq().
func intSeq() func() int {
	i := 0
	return func() int {
		i++
		return i
	}
}
