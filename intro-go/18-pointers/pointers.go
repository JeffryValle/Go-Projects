package main

import "fmt"

func main() {
	fmt.Println("**** Apuntadores / Punteros ***")
	// Crear, diferenciar  punteros,
	// la función new
	// aprender a trabajar con nil
	// types con punteros internos

	var a int = 42
	var b *int = &a // creamos el apuntador y le asignamos la
	// dirección de memoria de la variable a
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, b) // Imprime la dirección
	// b = 14 // No es permitido
	*b = 14
	fmt.Println(a, *b) // Imprime el valor y no su dirección

	ax := [3]int{1, 2, 3}
	bx := &ax[0]
	cx := &ax[1]
	fmt.Printf("%v %p %p\n", ax, bx, cx)

	var ms *myStruct
	// ms = &myStruct{foo: 42}
	fmt.Println(ms)
	ms = new(myStruct) // parecido a la instanciación
	// ms: instancia de myStruct
	ms.foo = 42
	fmt.Println(ms.foo)

	// Un slice contiene:
	// 1. Un puntero al array subyacente
	// 2. len: cuantos elementos tiene el slice
	// 3. cap: cuantos elementos puede contener

	az := map[string]string{"foo": "bar", "baz": "buz"}
	bz := az
	fmt.Println(az, bz)
	az["foo"] = "qux"
	fmt.Println(az, bz)

}

type myStruct struct {
	foo int
}
