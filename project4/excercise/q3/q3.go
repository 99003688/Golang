package main

import (
    "fmt"
	"time"
)

func main() {
	now := time.Now()
    my_chan := make(chan time.Time)
    chan_of_chans := make(chan chan time.Time)

    go func() {
        my_chan <- now
    }()

    go func() {
        chan_of_chans <- my_chan
    }()

    fmt.Println(<- <- chan_of_chans)
}

// package main

// import (
// 	"fmt"
// 	"time"
// )

// func currdate(dt time.Time, data chan time.Time) {
// 	dt = time.Now()
// 	data <- dt

// }

// func main() {

// 	var date time.Time
// 	ch := make(chan time.Time)
// 	go currdate(date, ch) 	read := <-ch
	// fmt.Println(read)

//  }