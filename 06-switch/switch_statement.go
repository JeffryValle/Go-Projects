package main

import (
	"fmt"
	"time" // paquete para medir y mostrar el time
)

func main() {
	fmt.Println("--------- Switch ---------")

	i := 2
	fmt.Println("Write ", i, " as ")
	switch i {
	case 1:
		fmt.Println("uno")
	case 2:
		fmt.Println("dos")
	case 3:
		fmt.Println("tres")
	}

	switch time.Now().Weekday() {
	case time.Saturday, time.Sunday:
		fmt.Println("It's the weekend")
	default:
		fmt.Println("It's a weekday")
	}

	/* switch sin una expresión lógica es una alternativa
	de trabajar con la lógica de if/else
	también consideramos cuando los casos no son
	constantes.
	*/
	t := time.Now()
	switch {
	case t.Hour() < 12:
		fmt.Println("It's before noon")
	default:
		fmt.Println("It's after noon")
	}

	// Función que recibe cualquier entrada de cualquier tipo de dato
	// y determina que tipo de dato es
	whatAmI := func(i interface{}) {
		switch t := i.(type) {
		case bool:
			fmt.Println("I'm a bool")
		case int:
			fmt.Println("I'm a int")
		default:
			// %s para string
			// %d para enteros
			// %f para decimales
			// %T imprime el tipo de dato
			fmt.Printf("Don't know type %T\n", t)
		}
	}

	whatAmI(true)
	whatAmI(1)
	whatAmI("hey")
	whatAmI(25.2555)

}
