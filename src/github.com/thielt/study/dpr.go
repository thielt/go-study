package main

//defer panic recover
//alters flow of execution

import (
	"fmt"
)

func main() {
	// DEFER:
	// fmt.Println("one")
	// defer fmt.Println("three")
	// fmt.Println("two")

	//if defer is put infront of all of them, then they are in LIFO order
	//doesn't execute  until function executes
	//but can waste memory, due to a bunch of resources being accessed

	// notes:
	// useful to group open and close functions together, probably not best for loops
	// variables will be called at defer time not at function time of execution

	a := "start"
	defer fmt.Println(a)
	a = "end" //doesn't print

	// PANIC:
	// b, c := 1, 0
	// ans := b / c
	// fmt.Println(ans)

	//similar to this:
	// fmt.Println("start")
	// panic("Something Bad Happened") //an error message
	// fmt.Println("end") //doesnt print "unreachable code"

	//order of execution:
	// execute main func
	// any defer statements
	// panic
	// return value

	// notes:
	// making an http request and then no result, is an error
	// panics are used when the program gets to a state it cant recover from i.e. 1/0
	// dont use when file cant be opened, unless critical
	// best used when cant obtain tcp port for web server
	// function will stop unless it has deferred statements

	//RECOVER:
	// used to recover from panics
	// only useful in deferred functions
	// current function will not continue, but higher functions in call stack will
}
