package main

import (
	"fmt"
	"reflect"
)

type Doctor struct {
	number     int
	actorName  string
	companions []string
}

func main() {

	fmt.Println("****** Struct *******")
	// un struct es como una clase sin métodos o
	// podría entenderse como de forma similar a un
	// objeto en otros lenguajes, GO no es
	// orientada a objetos
	aDoctor := Doctor{
		number:    3,
		actorName: "Jon Pertwee",
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}
	fmt.Println(aDoctor)

	// struct anónimo
	bDoctor := struct{ name string }{name: "Jeffry"}
	fmt.Println(bDoctor)
	fmt.Println(bDoctor.name)

	function()
}

type Animal struct {
	Name   string `required max:"100"`
	Origin string
}
type Bird struct {
	// Animal Animal
	Animal   // tipo embebido - similar a la herencia
	SpeedKPH float32
	CanFly   bool
}

func function() {
	t := reflect.TypeOf(Animal{})
	field, _ := t.FieldByName("Name")

	// instancia de Bird
	// aBird := Bird{
	// 	Animal:   Animal{Name: "Carpinter", Origin: "Austria"},
	// 	SpeedKPH: 22.22,
	// 	CanFly:   true,
	// }
	b := Bird{}
	b.Name = "Emu"
	b.Origin = "Australia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(field.Tag)

}
