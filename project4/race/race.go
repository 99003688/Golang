// package main

// import (
// 	"fmt"
// 	"time"
// )

// func execute(some string) {

// 	// initializing a infinite for loop
// 	for i := 1; true; i++ {

// 		// prints string and number
// 		fmt.Println(some, i)

// 		// this makes the program sleep for
// 		// 100 milliseconds, wiz 10 seconds
// 		time.Sleep(time.Millisecond * 100)

// 	}
// }

// func main() {

// 	// Simple go synchronous program
// 	// without any concurrency
// 	execute("First")
	
// 	// when this func is called it executes
// 	// and then passes on to next line
// 	execute("Second")
	
// 	// after the first second is to be executed  the problem is, execute function will
// 	// never finish execution, its trapped
// 	// in an infinite loop so the program will
// 	// never move to second execution.
// 	fmt.Println("program ends successfully")
	
// 	// if I'm wrong and both first and second
// 	// execute, then this line should print too
// 	// check the output
// }

// ---------------------------------------------------------------------------------------------------------

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func execute(some string) {

// 	// initializing a infinite for loop
// 	for i := 1; true; i++ {

// 		// prints string and number
// 		fmt.Println(some, i)

// 		// this makes the program sleep for
// 		// 100 milliseconds, wiz 10 seconds
// 		time.Sleep(time.Millisecond * 100)
// 	}
// }

// func main() {

// 	// Simple go program with concurrency
// 	go execute("First")

// 	// Placing the go command infront of the
// 	// func call simply creates a goroutine
// 	execute("Second")

// 	// The goroutine ensures that both functions
// 	// execute simultaneously & successfully
// 	fmt.Println("program ends successfully")

// 	// This statement still won't execute because
// 	// the func call is stuck in an infinite loop
// 	// check the output
// }

// // -----------------------------------------------------------------------------------------------------

// package main

// // one goroutine is the main
// // goroutine that comes by default
// import (
// 	"fmt"
// 	"runtime"
// 	"sync"
// )

// var wgIns sync.WaitGroup

// func main() {

// 	// shared variable
// 	counter := 0

// 	// the other 10 goroutines are
// 	// supposed to come from here
// 	wgIns.Add(10)
// 	for i := 0; i < 10; i++ {

// 		// goroutines are made
// 		go func() {
// 			for j := 0; j < 10; j++ {

// 				// shared variable execution
// 				counter += 1
// 				// 100 should be the counter value but
// 				// it may be equal to 100 or lesser
// 				// due to race condition
// 			}
// 			wgIns.Done()
// 		}()
// 	}

// 	// this value should actually be 11
// 	fmt.Println("The number of goroutines before wait = ", runtime.NumGoroutine())

// 	wgIns.Wait()

// 	// value should be 100
// 	fmt.Println("Counter value = ", counter)

// 	fmt.Println("The number of goroutines after wait = ", runtime.NumGoroutine())

// 	// this value is supposed to be 1
// 	// but lets see if these values
// 	// stay consistently same every
// 	// time we run the code
//

//-----------------------------------------------------------------------------------------------

// Golang program to fix the race
// condition using atomic package
package main

import (
	"fmt"
	"runtime"
	"sync"
	"sync/atomic"
)

// All goroutines will increment variable c
// waitgroup is waiting for the completion
// of program.
var (
	c		 int32
	waitgroup sync.WaitGroup
)

func main() {

	// with the help of Add() function add
	// one for each goroutine
	// a count of total 3
	waitgroup.Add(3)

	// increment with the help
	// of increment() function
	go increment("geeks")
	go increment("for")
	go increment("geeks")

	// waiting for completion
	// of goroutines.
	waitgroup.Wait()

	// print the counter
	fmt.Println("Counter:", c)

}

func increment(name string) {

	// Done() function used
	// to tell that it is done.
	defer waitgroup.Done()
    fmt.Println(name)
	for range name {

		// Atomic Functions
		// for fix race condition
		atomic.AddInt32(&c, 1)

		// enter thread in the line by line
		runtime.Gosched()// Gosched yields the processor, allowing other goroutines to run. It does not
		// suspend the current goroutine, so execution resumes automatically.
	}
}


