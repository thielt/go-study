package main

import (
	"fmt"
	"math"
)

func main() {
	if true {
		fmt.Println("The test is true")
	}
	//returns true
	//because the code in the curly braces is executed

	statePopulations := map[string]int{
		"California":   39250017,
		"Texas":        27862596,
		"Florida":      20612439,
		"New York":     19745289,
		"Pennsylvania": 12802503,
		"Illinois":     12801539,
		"Ohio":         11614373,
	}

	if pop, ok := statePopulations["Florida"]; ok {
		fmt.Println(pop)
	}
	//if florida is exists, print value

	// fmt.Println(pop) will not print outside of the block

	//go "short-cicuits"
	//if initial if statements are true, it wont evaluate the rest of the if statements

	// syntax:
	// if {

	// }else if {

	// }else {

	// }

	myNum := 0.123
	if math.Abs(myNum/math.Pow(math.Sqrt(myNum), 2)-1) < 0.001 {
		fmt.Println("These are the same")
	} else {
		fmt.Println("These are different")
	}
	// if you square a number then multiply it by itself it is the same
	//however it is evaluated by floating point numbers which arent always accurate, approximated
	// this code checks for within 1 thousandth of a float comparison, return true

	// for switch statements there can't be overlying cases

	switch i := 2 + 3; i {
	case 1, 5, 10:
		fmt.Println("1, 5 or 10")
	case 2, 4, 6:
		fmt.Println("2,4 or 6")
	default:
		fmt.Println("Some other")
		// break key word is implicitely placed
	}

	//in tagless syntax, they can overlap, but first case accepted will execute
	//unless you want it to go through all cases; use fallthrough

	switch j := 2 + 3; j {
	case 1, 5, 10:
		fmt.Println("1, 5 or 10")
		fallthrough
	case 2, 4, 6:
		fmt.Println("2,4 or 6")
	default:
		fmt.Println("Some other")
	}
	// //in this case, both are going to print, whether they are true or not

	var i interface{} = [3]int{} //can be any type
	switch i.(type) {
	case int:
		fmt.Println("int")
		//break
		//fmt.Println("This will not print, short-circuited")
	case float64:
		fmt.Println("float64")
	case string:
		fmt.Println("string")
	case [3]int:
		fmt.Println("[3]int")
	default:
		fmt.Println("some other type")
	}
}
