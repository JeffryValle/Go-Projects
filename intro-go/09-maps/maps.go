package main

import (
	"fmt"
)

func main() {
	fmt.Println("************ Mapas ************")

	// map[string] -> Tipo de dato de la llave : key-type
	// map[string]int -> Tipo de dato del valor : value-type
	m := make(map[string]int)

	m["k1"] = 7
	m["k2"] = 13

	fmt.Println("map: ", m)

	v1 := m["k1"]
	fmt.Println("v1:", v1)

	v3 := m["k3"]
	fmt.Println("v3:", v3) // Si la llave no existe, returna 0

	fmt.Println("len:", len(m))

	delete(m, "k2") // elimina los pares llave/valor del mapa
	fmt.Println("map: ", m)

	clear(m)                // borra todos los pares llave/valor del mapa
	fmt.Println("map: ", m) // devuelve un mapa vacío

	_, prs := m["k2"]
	fmt.Println("prs:", prs)

	// Alternativa para inicializar un mapa
	n := map[string]int{"foo": 1, "bar": 2}
	fmt.Println("map: ", n)

}
