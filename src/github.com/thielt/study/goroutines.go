package main

import (
	"fmt"
	"sync"
)

//concurrency, out of order, partial order computing, overlapping time periods rather than sequentially, execution not all at the same time
//- such as multiple clients accessing a server at the same time
//parallelism, computing multiple things at the same time, execution happens at the same time

// possible to have parallelism without concurrency
// and concurrency without parallelism

// parallelism needs multiple cores

var wg = sync.WaitGroup{}

//used to synchronize main function to the anonymous function go routine
//this is a way to avoid using time, which is based on a real-world clock rather than timing it with the time to execute the go routine

func main() {
	//go sayHello()                      //running on an os thread, smaller abstraction of thread to save memory for create and destroy go routines
	//time.Sleep(100 * time.Millisecond) //not good practice

	//second way, may cause issues if arguements not used
	var msg string = "Hello"

	wg.Add(1) //to syncronize to the below function
	go func(msg string) {
		fmt.Println("Hello2")
		wg.Done() //decrements the number of waitgroups to reach 0, letting the computer know its done
	}(msg)
	msg = "Goodbye"
	wg.Wait() //waiting on the wait group/go routine above to finish
}

// func sayHello() {
// 	fmt.Println("Hello")
// }

// go run -race src/sample/goroutines.go
// we get all additional information where race conditions are shown
