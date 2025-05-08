package main

import (
	"fmt"
	"math"
)

/*
	El sombreado es una consecuencia directa del concepto de alcance en Go.
	Cuando se declara una variable dentro de un ámbito, ese ámbito "recuerda"
	esa variable y, si se encuentra una variable con el mismo nombre en un
	ámbito más interno, la variable del ámbito interno "sombra" a la del ámbito externo.
*/
/*
	Debido a que son constantes estan no ocupan espacio en memoria, funcionan en tiempo de
	ejecución.
	Pero si cambiaramos las constantes por variables, ahí ocuparía espacio en memoria.
	Aún ese caso sigue sucediendo el shadowing.
*/
const s string = "constant"
const a int16 = 17

func main() {
	fmt.Println("**************Constantes**************")
	// shadowing : sombre de nombres
	// Ocurre cuando en un ambito externo se declara un nombre a una variable
	// y en un ambito interno se declara una variable con el mismo nombre
	// haciendo que "oculte" temporalmente la variable del ámbito externo.
	const a int = 15
	const b int16 = 16
	fmt.Printf("%v %T\n ", a, a)
	// fmt.Printf("%v %T\n ", a + b, a) !!! Error: No coinciden los tipos de datos

	fmt.Println(s)

	const n = 500000000

	const d = 3e20 / n
	fmt.Println(d)
	fmt.Println(int64(d))

	num := (10.5 / 3.5)
	fmt.Println(num)
	fmt.Println(math.Ceil(num))

	fmt.Println(math.Sin(n))

	values()
}

const (
	x = iota //
	b = 1 << (10 * iota)
	c
	d
	e
	f
	g
)

func values() {
	fileSize := 4000000000.
	fmt.Println("Valor de b", b)
	fmt.Printf("%.2fd", fileSize/d)
}
