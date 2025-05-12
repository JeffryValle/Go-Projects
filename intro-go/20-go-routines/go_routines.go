package main

import (
	"fmt"
	// "runtime"
	"sync"
)

// Concurrencia: Capacidad de un
// sistema para realizar múltiples procesos o tareas
// Paralelismo: Capacidad de un
// sistema o más para realizar múltiples procesos
// o tareas simultaneamente

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {

	var msg = "Hello"
	wg.Add(1)
	go func() {
		fmt.Println(msg)
		wg.Done()
	}()
	msg = "Goodbye"
	wg.Wait()

	// for i := 0; i < 10; i++ {
	// 	wg.Add(2)
	// 	m.RLock()
	// 	go sayHello()
	// 	m.Lock()
	// 	go increment()
	// }
	// wg.Wait()

	// fmt.Println("Threads: %v\n", runtime.GOMAXPROCS(-1))
}

// func sayHello() {
// 	fmt.Printf("Hello #%v\n", counter)
// 	m.RUnlock()
// 	wg.Done()
// }

// func increment() {
// 	counter++
// 	m.Unlock()
// 	wg.Done()
// }
