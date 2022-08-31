package main

import (
	"fmt"
)

func main() {
	var w Writer = ConsoleWriter{} //intializing that ConsoleWriter has Writer behavior
	w.Write([]byte("Hello Go!"))
}

// polymorphic behavior
type Writer interface {
	Write([]byte) (int, error) //behavior, method definitions
}

type ConsoleWriter struct{} //implement Writer interface and its behaviors

// method called Write
// accepting slice of bytes
// returning and int and an error
// converts byte slice into string and return to console
func (cw ConsoleWriter) Write(data []byte) (int, error) {
	n, err := fmt.Println(string(data))
	return n, err
}
