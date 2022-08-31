package main

import (
	"fmt"
	//pointer arithmatic: import "unsafe"
)

func main() {
	a := 42
	b := a
	fmt.Println(a, b)
	a = 27
	fmt.Println(a, b)
	//proving that they are NOT pointed to the same memory

	var c int = 42
	var d *int = &c
	fmt.Println(c, *d)
	//output: 42 0x14000016078
	//* declares its a pointer
	//& "at the address of"
	//* at a pointer is a dereferencing operator, pulling the value at that pointer
	c = 320
	fmt.Println(c, *d)
	//320 320
	*d = 67
	fmt.Println(c, *d)
	//67 67

	//a pointer not initialized, automatically assigned nil
	//the derencing operator has a lower precedence then dot operator
	//i.e. (*ms).foo = "some value"
	//so we need parenthesis for derefencing when dot operator used

	// slices and maps have some parts that automatically use pointers if not initialized in size
}
