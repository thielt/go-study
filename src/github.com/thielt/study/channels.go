package main

import (
	"fmt"
	"sync"
)

var wg = sync.WaitGroup{}

// func main() {
// 	ch := make(chan int)
// 	wg.Add(2)
// 	go func() {
// 		i := <-ch //receiving the value out of i
// 		fmt.Println(i)
// 		wg.Done()
// 	}()

// 	go func() {
// 		i := 42
// 		ch <- i //sending copy of the value into channel
// 		i = 27
// 		wg.Done()
// 	}()
// 	wg.Wait()
// }

// another way:
func main() {
	ch := make(chan int, 50) //create a channel that can store 50 integers
	wg.Add(2)

	go func(ch <-chan int) {
		i := <-ch
		fmt.Println(i)
		i = <-ch       //so the new information(27) can be outputted, this is a buffer channel, channel blocks sender side until receiver available and vice versa
		fmt.Println(i) // buffer channels are usually used when theres asymmetric loading
		wg.Done()
	}(ch)

	go func(ch chan<- int) {
		ch <- 42
		ch <- 27 // the 50 will lose this information
		wg.Done()
	}(ch)
	wg.Wait()
}

// example:
// earthquakes will send data every hour in the seismonitor
// stores data in chunks and sends burst of information
// sender or receiver needs time to process, because of a little delay

//research select statements
//go routine monitors  multiple channels at once
//blocks if all channels block
//if there multiple channels that receive the same value at the same time, behavior is undef
