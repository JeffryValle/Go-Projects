package main

import (
	"fmt"
	"math" // paquete para funciones matematicas
)

const s string = "constant"

func main() {
	fmt.Println("**************Constantes**************")
	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))

	num := (10.5 / 3.5)
	fmt.Println(num)
	fmt.Println(math.Ceil(num))

	fmt.Println(math.Sin(n))

}
