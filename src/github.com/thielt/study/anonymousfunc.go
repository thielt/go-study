package main

import (
	"fmt"
)

// methods
func main() {
	for i := 0; i < 5; i++ {
		func(i int) {
			fmt.Println(i)
		}(i) //an anonymous func is short lived, executed and defined at the same time
	}

	//another way
	var f func() = func() {
		fmt.Println("Hello")
	}
	f()
}
