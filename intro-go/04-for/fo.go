package main

import "fmt"

// En go 'for' es la única estructura de repetición
// No existe While, Do While, forEach...

func main() {
	fmt.Println("*********** For loop ***********")

	// i := 1 // asignamos una variable inicializadora
	// for i <= 3 {
	// 	fmt.Println(i)
	// 	i = i + 1
	// }
	// // sintaxis más manejable
	// for j := 0; j < 3; j++ {
	// 	fmt.Println(j)
	// }

	// for i := range 3 {
	// 	fmt.Println(i)
	// }

	// for {
	// 	fmt.Println("loop")
	// 	break
	// }
	// fmt.Println("****************************")
	// for n := range 6 {
	// 	if n%2 == 0 {
	// 		continue
	// 	}
	// 	fmt.Println(n)
	// }

	promedioNotas(91, 92, 93, 94, 95)

}

func promedioNotas(nums ...int) /* aqui se agrega el tipado cuando querramos devolver un valor */ {

	fmt.Println("********************************")
	fmt.Println("** Registro de Calificaciones **")
	fmt.Println("**                            **")
	fmt.Println("********************************")

	sum := 0
	nota := 0
	for i := 0; i < len(nums); i++ {
		sum = sum + nums[i]
		nota = nums[i]
		fmt.Println("Nota", i+1, " :", nota)
	}

	fmt.Println(" Promedio de Notas: ", sum/len(nums))
}
