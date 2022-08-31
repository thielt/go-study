package main

import (
	"fmt"
)

//struct fields describe any type of data you want
//collection of many different data types

type Doctor struct {
	number     int
	actorName  string
	companions []string //slice of strings
	episodes   []string
	//if there is another value initialized then it should be
	//put in the main function in that same order, else it will be return a different value
	//unless explicitely defined
}

func main() {
	aDoctor := Doctor{
		number:    3,
		actorName: "John Pertwee",
		episodes:  []string{}, //empty place holder, but not neccessary when explicitely defined by field
		companions: []string{
			"Liz Shaw",
			"Jo Grant",
			"Sarah Jane Smith",
		},
	}

	//anonymous struct , usually short lived
	//intialized and assigned at the same time in the main function
	anotherDoc := struct{ name string }{name: "John Pertwee"}

	fmt.Println(aDoctor)
	fmt.Println(aDoctor.actorName)     //prints out actor name only
	fmt.Println(aDoctor.companions[1]) //Jo Grant
	fmt.Println(aDoctor.companions)
	fmt.Println(anotherDoc)
}
