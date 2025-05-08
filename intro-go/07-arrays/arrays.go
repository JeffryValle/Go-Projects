package main

import "fmt"

/* Los arreglos en lenguajes estaticamente tipados solo pueden almacenar
valores del mismo tipo de datos
*/

/*Cuando decimos que un lenguaje como Go es estáticamente tipado, significa que:
El tipo de cada variable debe conocerse (y validarse) en tiempo de compilación, no en tiempo de ejecución
*/

func main() {
	fmt.Println("************ Arrays ************")

	// Por defecto al declarar la cantidad de elementos
	// de tipo int sin ningun valor inicializado, Go
	// generará los valores cero 0 ada cada elemento
	// var a [5]int
	// fmt.Println("emp:", a)

	// a[4] = 100
	// fmt.Println("set:", a)
	// fmt.Println("get:", a[4])

	// fmt.Println("len:", len(a))

	// b := [5]int{1, 2, 3, 4, 5}
	// fmt.Println("dcl:", b)

	// b = [...]int{1, 2, 3, 4, 5}
	// fmt.Println("dcl:", b)

	// b = [...]int{100, 2: 300, 400, 500}
	// fmt.Println("idx:", b)
	// b = [...]int{100, 3: 400, 500}
	// fmt.Println("idx:", b)
	// b = [...]int{100, 4: 500}
	// fmt.Println("idx:", b)

	// var twoD [2][3]int
	// for i := 0; i < 2; i++ {
	// 	for j := 0; j < 3; j++ {
	// 		twoD[i][j] = i + j
	// 		fmt.Println(i, j)
	// 	}
	// }
	// fmt.Println("2d: ", twoD)

	// twoD = [2][3]int{
	// 	{1, 2, 3},
	// 	{1, 2, 3},
	// }
	// fmt.Println("2d: ", twoD)

	arrays()
}

func arrays() {

	// var students [5]string
	// fmt.Printf("Students: %v\n", students)
	// students[0] = "Jeffry"
	// students[1] = "Stephanie"
	// students[2] = "Pablo"
	// fmt.Printf("Students: %v\n", students)
	// fmt.Printf("Number of Students: %v\n", len(students))

	// var identityMatrix [3][3]int
	// identityMatrix[0] = [3]int{1, 0, 0}
	// identityMatrix[1] = [3]int{0, 1, 0}
	// identityMatrix[2] = [3]int{0, 0, 1}
	// fmt.Printf(" %v\n", identityMatrix)
	// fmt.Printf(" %v\n", len(identityMatrix))
	// // Imprime la matriz en 2D
	// for i := 0; i < len(identityMatrix); i++ {
	// 	for j := 0; j < len(identityMatrix); j++ {
	// 		fmt.Print(" ", identityMatrix[i][j])
	// 	}
	// 	fmt.Println()
	// }

	// array1 := [...]int{1, 2, 3}
	// array2 := array1
	// array2[1] = 5
	// fmt.Println(array1)
	// fmt.Println(array2)
	// fmt.Println("************Antes************")
	// array3 := [...]int{1, 2, 3}
	// array4 := &array3 // creamos un puntero al array3
	// // el array4 ahora es una referencia(puntero) al
	// // mismo arreglo en memoria
	// array4[1] = 5
	// fmt.Println(array3)
	// fmt.Println(array4)
	// fmt.Println("************Después************")

	/* Slice: Estructura de datos dinámica */
	// array1 := []int{1, 2, 3}
	// array2 := array1
	// array2[1] = 5
	// fmt.Println(array1)
	// fmt.Println(array2)
	// fmt.Printf("Length  : %v\n", len(array1))
	// fmt.Printf("Capacity: %v\n", cap(array1))

	// ax := []int{1, 2, 3, 4, 5, 6, 7, 8, 9, 10}
	// bx := ax[:]   // slice con todos los elementos
	// cx := ax[3:]  // slice from 4th element to end
	// dx := ax[:6]  // slice first 6 elements
	// ex := ax[3:6] // slice the 4th, 5th and 6th elements
	// // Basically:
	// // the first number is inclusive
	// // the second number is exclusive
	// ax[5] = 42
	// fmt.Println(ax)
	// fmt.Println(bx)
	// fmt.Println(cx)
	// fmt.Println(dx)
	// fmt.Println(ex)

	// function: make()
	// make makes a slice, maps or canals, it doesn't create arrays
	// ax := make([]int, 3, 100)
	// fmt.Println(ax)
	// fmt.Printf("Length  : %v\n", len(ax))
	// fmt.Printf("Capacity: %v\n", cap(ax))
	// ax = append(ax, 1)
	// fmt.Println(ax)
	// fmt.Printf("Length  : %v\n", len(ax))
	// fmt.Printf("Capacity: %v\n", cap(ax))
	// ax := []int{}
	// fmt.Println(ax)
	// fmt.Printf("Length  : %v\n", len(ax))
	// fmt.Printf("Capacity: %v\n", cap(ax))
	// ax = append(ax, 1)
	// fmt.Println(ax)
	// fmt.Printf("Length  : %v\n", len(ax))
	// fmt.Printf("Capacity: %v\n", cap(ax))
	// // el simbolo ... le dice al slice que separe
	// // cada valor individualmente
	// ax = append(ax, []int{2, 3, 4, 5}...)
	// fmt.Println(ax)
	// fmt.Printf("Length  : %v\n", len(ax))
	// fmt.Printf("Capacity: %v\n", cap(ax))
	// b := ax[:len(ax)-1]
	// fmt.Println(b)

	ax := []int{1, 2, 3, 4, 5}
	bx := append(ax[:2], ax[3:]...)
	fmt.Println(bx)
}
