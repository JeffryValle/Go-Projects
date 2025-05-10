package main

import (
	"fmt"
	"io"
	"log"
	"net/http"
	// "net/http"
)

// Las funciones Defer obedecen el orden LIFO Last in First out
// Ejecución Diferida
// Pospone la ejecución de una función hasta que la
// función que la contiene termine

func main() {
	fmt.Println("******** Defer ********")

	// defer fmt.Println("start")
	// defer fmt.Println("middle")
	// defer fmt.Println("end")

	res, err := http.Get("https://pokeapi.co/api/v2/pokemon?limit=3&offset=0")
	if err != nil {
		log.Fatal(err)
	}
	defer res.Body.Close()
	pokemons, err := io.ReadAll(res.Body)
	if err != nil {
		log.Fatal(err)
	}
	log.Println(string(pokemons))
	// fmt.Printf("Nombre %s", pokemons)
	// fmt.Printf("Url %s", pokemons)

	principal()
	recovery()
}

// Panic: detiene inmediatamente la ejecución normal del
// programa, empieza a correr los defer, y después sale
// del programa.
// Se usa para propagar errores que no se pueden o
// no deben manejar localmente.
func principal() {
	// fmt.Println(" ----Panic----- ")

	http.HandleFunc("/", func(w http.ResponseWriter, r *http.Request) {
		w.Write([]byte("Hello Go!"))
	})
	err := http.ListenAndServe(":8080", nil)
	if err != nil {
		panic(err.Error())
	}

	fmt.Println("start")
	defer fmt.Println("this was deferred")
	panic("Somethin bad happened")
	// fmt.Println("end") // No se imprime

	// panics happen after deferred statements are executed
}

func recovery() {
	// fmt.Println("----- Recover -----")
	// fmt.Println("start")
	// defer func() { // función anónima
	// 	if err := recover(); err != nil {
	// 		log.Println("Error:", err)
	// 	}
	// }()
	// panic("Somethin bad happened")
	// fmt.Println("end")

	fmt.Println("start")
	panicker()
	fmt.Println("end")
}

// recover - Capturar un panic
/* Se usa dentro de una función defer para capturar
y manejar un panic, evitando que el programa termine
abruptamente
*/

func panicker() {
	fmt.Println("about to panic")
	defer func() {
		if r := recover(); r != nil {
			log.Println("Error: ", r)
		}
	}()
	panic("Something bad happened")
	// fmt.Println("done panicking") // No lo alcanza
}
