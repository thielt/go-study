package main

import (
	"fmt"
)

func main() {
	for i := 0; i < 5; i++ {
		sayMessage("Hello Go!", i)
	}

	//for variadic example
	s := sum("The Sum is", 1, 2, 3, 4, 5)
	fmt.Println("sum", s)
}

func sayMessage(msg string, idx int) {
	fmt.Println(msg)
	fmt.Println("index value: ", idx)
}

//all parameters are the same type, then type declaration not neccesary
// func someFunc(msg, idx string) assumes all are the same

//using pointers in call function can alter main function

// variadic paramter, can only have one and has to be at the end
func sum(msg string, values ...int) int { //add return values type at end
	fmt.Println(values) // just to make sure we're printing the right values
	result := 0
	for _, v := range values {
		result += v
	}
	//fmt.Println(msg, result)
	return result
}
