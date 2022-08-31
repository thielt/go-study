package main

import (
	"fmt"
)

func main() {
	//basic for loop
	for i := 0; i < 5; i++ {
		fmt.Println(i)
	}

	//2 for loops
	for i, j := 0, 0; i < 5; i, j = i+1, j+1 {
		fmt.Println(i, j)
	}

	// j's will be 0's
	for i, j := 0, 0; i < 5; i++ {
		fmt.Println(i, j)
	}

	//try not to change loops inside block, they execute in different order
	//here it is also a basic for loop, but scoped to the main function
	i := 0
	for ; i < 5; i++ {
		fmt.Println(i) //will print out 1 to 5
	}
	fmt.Println(i) //just prints out 5, if no println in loop block

	//another alternate
	j := 0
	for j < 5 {
		fmt.Println(j)
		j++
	}

	//another alternate
	k := 0
	for {
		fmt.Println(k)
		k++
		if k == 5 {
			break
		}
	}

	//odd numbers
	for l := 0; l < 10; l++ {
		if l%2 == 0 {
			continue
		}
		fmt.Println(l)
	}

	//There is also Loop and break Loop to break out of certain loops in nested loops

	//enumeration in python is equivalent to this for slices and arrays(collections):
	s := []int{1, 2, 3}
	for k, v := range s {
		fmt.Println(k, v)
	} //for strings: values are unicode outputted

	//for keys only:
	for a := range s {
		fmt.Println(a)
	} //output: 0,1,2

	//for values only:
	for _, b := range s {
		fmt.Println(b)
		// }output: 1,2,3
	}
}
