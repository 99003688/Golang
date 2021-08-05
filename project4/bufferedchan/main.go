package main

import "fmt"

func main() {

    messages := make(chan string, 2)

    messages <- "buffered"
    messages <- "channel"

    fmt.Println(<-messages)
    fmt.Println(<-messages)
}

// ----------------------------------------------------

// package main

// import "fmt"

// func main() {
// 	ch := make(chan int, 3)
// 	ch <- 1
// 	ch <- 2
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

// ---------------------------------------------------------

// package main

// import (
// 	"fmt"
// )

// func main() {

// 	// create a buffered channel
// 	// with a capacity of 2.
// 	ch := make(chan string, 2)
// 	ch <- "Saujanya"
// 	ch <- "Tailang"
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }

// -----------------------------------------------------------------

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func write(ch chan int) {
// 	for i := 0; i < 4; i++ {
// 		ch <- i
// 		fmt.Println("successfully wrote", i, "to ch")
// 	}
// 	close(ch)
// }
// func main() {

// 	// creates capacity of 2
// 	ch := make(chan int, 2)
// 	go write(ch)
// 	time.Sleep(2 * time.Second)
// 	for v := range ch {
// 		fmt.Println("read value", v, "from ch")
// 		time.Sleep(2 * time.Second)

// 	}
// }

// ----------------------------------------------------------------------------

// package main

// import (
// 	"fmt"
// )

// func main() {
// 	ch := make(chan string, 2)
// 	ch <- "saujanya"
// 	ch <- "guna"
// 	ch <- "amit"
// 	fmt.Println(<-ch)
// 	fmt.Println(<-ch)
// }
