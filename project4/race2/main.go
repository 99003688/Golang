// package main

// import (
//     "fmt"
//     "sync"
// )

// var Wait sync.WaitGroup
// var Counter int = 0

// func main() {

//     for routine := 1; routine <= 2; routine++ {

//         Wait.Add(1)
//         go Routine(routine)
//     }

//     Wait.Wait()
// //     fmt.Printf("Final Counter: %d\n", Counter)
// // }

// // func Routine(id int) {

// //     for count := 0; count < 2; count++ {

// //         value := Counter
// //         value++
// //         Counter = value
// //     }

// //     Wait.Done()
// // }

// // ------------------------------------------------------------------------

// package main

// import (
//     "fmt"
//     "sync"
//     "time"
// )

// var Wait sync.WaitGroup
// var Counter int = 0

// func main() {

//     for i := 1; i <= 2; i++ {

//         Wait.Add(1)
//         go Routine(i)
//     }

//     Wait.Wait()
//     fmt.Printf("Final Counter: %d\n", Counter)
// }

// func Routine(id int) {

//     for i := 0; i < 2; i++ {

//         value := Counter
//         time.Sleep(1 * time.Nanosecond)
//         value++
//         Counter = value
//     }

//     Wait.Done()
// }

// ----------------------------------------------------------------------------------

// package main

// import (
//     "fmt"
//     "sync"
//     "time"
// )

// var Wait sync.WaitGroup
// var Counter int = 0

// func main() {

//     for routine := 1; routine <= 2; routine++ {

//         Wait.Add(1)
//         go Routine(routine)
//     }

//     Wait.Wait()
//     fmt.Printf("Final Counter: %d\n", Counter)
// }

// func Routine(id int) {

//     for count := 0; count < 2; count++ {

//         Counter = Counter + 1
//         time.Sleep(1 * time.Nanosecond)
//     }

//     Wait.Done()
// }
// Copy the value of Counter to BX
// Increment the value of BX
// Move the new value to Counter
// -------------------------------------------------------------------------------

package main

import (
    "fmt"
    "sync"
    "time"
)

var Wait sync.WaitGroup
var Counter int = 0
var Lock sync.Mutex

func main() {

    for routine := 1; routine <= 2; routine++ {

        Wait.Add(1)
        go Routine(routine)
		fmt.Println("hello")
    }

    Wait.Wait()
    fmt.Printf("Final Counter: %d\n", Counter)
}

func Routine(id int) {

    for count := 0; count < 2; count++ {

        Lock.Lock()
		fmt.Println("hii")

        value := Counter
        time.Sleep(1 * time.Nanosecond)
        value++
        Counter = value

        Lock.Unlock()
    }

    Wait.Done()
}