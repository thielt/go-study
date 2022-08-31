package main

import (
	"fmt"
)

//example of using error

func main() {
	d, err := divide(5.0, 0.0)
	if err != nil {
		fmt.Println(err)
		return
	}
	fmt.Println(d)
}

func divide(a, b float64) (float64, error) { //end is return type
	if b == 0.0 {
		return 0.0, fmt.Errorf("Cannot divide by zero")
	}
	return a / b, nil
}
