package main

import (
	"fmt"
	"slices"
)

func main() {

	var s []string
	fmt.Println("uninit: ", s, s == nil, len(s))

	s = make([]string, 3)
	fmt.Println("emp: ", s, "len: ", len(s), "cap: ", cap(s))

	s[0] = "a"
	s[1] = "b"
	s[2] = "c"

	fmt.Println("set: ", s)
	fmt.Println("get: ", s[2])

	fmt.Println("len: ", len(s))

	s = append(s, "d")
	s = append(s, "e", "f")
	fmt.Println("apd: ", s)

	c := make([]string, len(s))
	copy(c, s) // copia un slice

	fmt.Println("cpy", c)
	fmt.Println("/////////////////////")
	l := s[2:5] // copia del slice del elemento x hasta y, x:y
	fmt.Println("sl1:", l)

	l = s[:5] // excluye el elemento en la posición 5
	fmt.Println("sl2:", l)

	l = s[:2] // excluye los elementos superiores desde la posición 2 hasta la última
	fmt.Println("sl3:", l)

	// slices: son tipos diferentes de arreglos

	t := []string{"g", "h", "i"}
	fmt.Println("dcl:", t)

	t2 := []string{"g", "h", "i"}
	if slices.Equal(t, t2) {
		fmt.Println("t == t2")
	} else {
		fmt.Println("t != t2")
	}

	twoD := make([][]int, 3) // declaramos un array bidimensional
	fmt.Println("2d: ", twoD)
	for i := 0; i < 3; i++ {
		innerLen := i + 1               // contador
		twoD[i] = make([]int, innerLen) // creará arrays dentro del array inicial
		for j := 0; j < innerLen; j++ { // asignará valores segun la posición
			twoD[i][j] = i + j
		}
	}
	fmt.Println("2d: ", twoD)

}
