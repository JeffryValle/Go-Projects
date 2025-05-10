package main

import "fmt"

func main() {
	fmt.Println("*********** Funciones *********** ")
	res := plus(25, 24)
	fmt.Println("Suma: ", res)

	sum := plusPlus(10, 20, 30)
	fmt.Println("Suma: ", sum)
	// for i := range 5 {
	// 	sayMessage("Hello Go!", i)
	// }

	// sayGreeting("Hello", "Stacey")

	greeting := "Hello"
	name := "Stacey"
	sayGreeting(&greeting, &name)
	fmt.Println(name)

	s := suma(1, 2, 3, 4, 5)
	fmt.Println("The sum is: ", s)

	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)

}

func divide(a, b float64) (float64, error) {
	if b == 0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}

func suma(values ...int) (result int) {
	fmt.Println(values)
	for _, v := range values {
		result += v
	}
	return
}

// func sayGreeting(greeting, name string) {
// 	fmt.Println(greeting, name)
// }
// Aqui lo que realmente se pasa es una copia del valor, no la referencia original
// Así que dentro de sayGreeting, name es una copia local de "Stacey"
// func sayGreeting(greeting, name string) {
// 	fmt.Println(greeting, name)
// 	name = "Ted"
// 	fmt.Println(name)
// }
func sayGreeting(greeting, name *string) {
	fmt.Println(*greeting, *name)
	*name = "Ted"
	fmt.Println(*name)
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("The value of the index is: ", idx)
}

func plus(a int, b int) int {
	return a + b
}

func plusPlus(a, b, c int) int {
	return a + b + c
}
