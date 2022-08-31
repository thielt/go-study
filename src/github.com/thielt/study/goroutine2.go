// package main

// import (
// 	"fmt"
// 	"sync"
// )

// var wg = sync.WaitGroup{}
// var counter = 0

// func main() {
// 	for i := 0; i < 10; i++ { //creating 20 go routines
// 		wg.Add(2)
// 		go sayHello()
// 		go increment()
// 	}
// 	wg.Wait()
// }

// func sayHello() {
// 	fmt.Printf("Hello #%v\n", counter)
// 	wg.Done()
// }

// func increment() {
// 	counter++
// 	wg.Done()
// }
//there is no syncronization and they are racing against each other, mutex is needed

// as many things can read, but only one can write at a time
// but if anything is reading, we can't write to it at all
// infinite readers, only writer
// writer will lock the mutex until its done
package main

import (
	"fmt"
	"runtime"
	"sync"
)

var wg = sync.WaitGroup{}
var counter = 0
var m = sync.RWMutex{}

func main() {
	runtime.GOMAXPROCS(100) //setting how many threads there are on the system being used
	//scheduler has to work harder because how many threads are managed, but the higher the faster
	//used to test where performance is best
	for i := 0; i < 10; i++ {
		wg.Add(2)
		m.RLock() //Read Lock on the mutex
		go sayHello()
		m.Lock() //Write Lock on the mutex
		go increment()
	}
	wg.Wait()
} //Pretty much makes the go routine useless, because we're constantly locking and unlocking

func sayHello() {
	fmt.Printf("Hello2 #%v\n", counter)
	m.RUnlock()
	wg.Done()
}

func increment() {
	counter++
	m.Unlock()
	wg.Done()
}
