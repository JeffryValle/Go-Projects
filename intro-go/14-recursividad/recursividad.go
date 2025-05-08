package main

import "fmt"

func main() {
	fmt.Println("****** Recursividad ******")

	fmt.Println(factorial(5))
	fmt.Println("****************************")
	fmt.Println(fibonacci(5))

}

func factorial(n int) int {
	if n == 0 {
		return 1
	}
	fmt.Print(" ")
	fmt.Print(n, " ")
	return n * factorial(n-1)
}

func fibonacci(n int) int {
	if n < 2 {
		return n
	}
	return fibonacci(n-1) + fibonacci(n-2)
}
