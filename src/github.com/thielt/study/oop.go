package main

import (
	"fmt"
)

type Animal struct {
	Name string
	//tags example, used to describe a field
	//Name string `required max: "100"`
	//neet to import "reflect"
	Origin string
}

type Bird struct {
	Animal //embedding the Animal struct right into the Bird struct, without this they are independent
	// Animal Penguin, this embeds another field called Penguin
	SpeedKPH float32
	CanFly bool
}

func main() {
	b := Bird{}
	n.Name = "Emu"
	b.Origin = "Austrailia"
	b.SpeedKPH = 48
	b.CanFly = false
	fmt.Println(b)
	fmt.Println(b.Name) //go is handling it automatically, bird has animal like characteristics, but not an animal

	//have to be aware that you are using embedding only, no inheritance exists in go
	//better to use interfaces for in common behavior
	b := Bird{
		Animal: Animal{Name: "Emu,", Origin: "Austrailia"},
		SpeedKPH: 48,
		CanFly: false,
	}
	fmt.Println(b)
}