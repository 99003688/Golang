// // package main


// // import (
// // 	"database/sql"
// // 	// "fmt"

// // 	_ "github.com/go-sql-driver/mysql"
// // )

// // func main(){

// // 	database, err := sql.Open("mysql", "root:root@tcp(localhost:3306)/test")
// // 	if err != nil {
// //         panic(err.Error())
// //     }

// // 	//Create
// // 	statement, _ := 
// // 		database.Prepare("CREATE TABLE IF NOT EXISTS people (id INTEGER PRIMARY KEY, firstname TEXT, lastname TEXT)")
// //     statement.Exec()

// // 	//insert
// // 	statement, _ = 
// // 		database.Prepare("INSERT INTO people (firstname, lastname) VALUES (?, ?)")
// //     statement.Exec("Saujanya", "Tailang")


// // }

// // ----------------------------------------------------------------------------------------------------------
// package main

// import (
//     "fmt"
//     "time"
// )

// type Fibonacci struct {
//     num    float64
//     answer float64
// }

// func newFibonacci(n float64) *Fibonacci {

//     f := new(Fibonacci)
//     f.num = n
//     c1 := make(chan float64)
//     c2 := make(chan float64)

//     if f.num <= 1 {
//         f.answer = n
//     } else {
//         go func() {
//             fib1 := newFibonacci(n - 1)
//             c2 <- fib1.answer
//         }()
//         go func() {
//             fib2 := newFibonacci(n - 2)
//             c1 <- fib2.answer   
//         }()

//         f.answer = <-c2 + <-c1
//     }
//     close(c1)
//     close(c2)

//     return f
// }

// func main() {

//     numbers := []float64{30, 35, 36, 37}
//     for _, value := range numbers{
//         start := time.Now()
//         fmt.Println("getting the ", value, " fibonacci number")
//         f := newFibonacci(value)
//         fmt.Println(f.answer)
//         end := time.Now()
//         totalTime := end.Sub(start)
//         fmt.Println("Fibonacci number: ", value, " took ", totalTime)

//     }

// }

// // ------------------------------------------

// package main

// import(
//     "fmt"
//     // "time"
// )

// func fib() chan int {
//     c := make(chan int)
  
//     go func() { 
//       for i, j := 0, 1; ; i, j = i+j,i {
//           c <- i
//       }
//     }()
  
//     return c
//   }

//   func main() {
//     c := fib()
//     for n := 0; n < 12 ; n++ {
//         fmt.Println(<- c)
//     }
// }

// ----------------------------------------------------------
// package main

// import (
// 	"fmt"
// 	// "strings"
// 	"time"

// "sync"
// )
// func count(str string){
// 	for i:=1; i<6;i++{
// 		fmt.Println(i,str)
// 		time.Sleep(time.Millisecond * 500)
// 	}
// }
// // func itarr(arr[] int){

// // for i := 0; i < len(arr); i++ {
// // 	fmt.Println(arr[i])
// // }
// // }

//  func main(){
// var wg sync.WaitGroup
// wg.Add(1)
// go func (){
// 	count("sheep")
// 	wg.Done()
// }()
//  go count("fish")

// wg.Wait()

// 	go count("concurrency or go routine")
// 	 count("mapping")


// //  var arr=[]int {1,2,3,4,5}
  
// //  go itarr(arr)

// //  itarr(arr)

// c1:= make(chan string)
// c2:= make(chan string)

// go func(){
// 	for i:=1; i<6;i++{
// 		c1<- "every 500ms"
// 		time.Sleep(time.Millisecond*500)
// 	}
// }()

// go func(){
// 	for i:=1; i<6;i++{
// 		c2<- "every 2 sec"
// 		time.Sleep(time.Second*2)
// 	}
// }()

// for{
// 	fmt.Println(<-c1)
// 	fmt.Println(<-c2)
// }

// fmt.Scanln()


// }

// // ------------------------------------------------------------------

// package main

// import (
//     "fmt"
//     "time"
// )

// type Fibonacci struct {
//     num    float64
//     answer float64
// }

// func newFibonacci(n float64) *Fibonacci {

//     f := new(Fibonacci)
//     f.num = n
//     c1 := make(chan float64)
//     c2 := make(chan float64)

//     if f.num <= 1 {
//         f.answer = n
//     } else {
//         go func() {
//             fib1 := newFibonacci(n - 1)
//             c2 <- fib1.answer
//         }()
//         go func() {
//             fib2 := newFibonacci(n - 2)
//             c1 <- fib2.answer   
//         }()

//         f.answer = <-c2 + <-c1
//     }
//     close(c1)
//     close(c2)

//     return f
// }

// func main() {

//     numbers := []float64{30, 35, 36, 37, 38, 39, 40}
//     for _, value := range numbers{
//         start := time.Now()
//         fmt.Println("getting the ", value, " fibonacci number")
//         f := newFibonacci(value)
//         fmt.Println(f.answer)
//         end := time.Now()
//         totalTime := end.Sub(start)
//         fmt.Println("Fibonacci number: ", value, " took ", totalTime)

//     }

// }

// -----------------------------------------------------------------------------------------
//  fibonacci series using select statement 
// package main

// import (
//     "fmt"
// )

// func finonaccin(n int, c chan int ){
//     x, y := 0,1
//     for i:=0; i<n; i++{
//         c<-x
//         x, y =y,x+y
//     }
//     close(c)

// }

// func main(){
//     c := make(chan int , 10)
//     go finonaccin(cap(c), c)
//     for i := range c{
//         fmt.Println(i)

//     }
// }

// ----------------------------------------------------------
// codes  gives 5 fibonacci series
package main

import(
    "fmt"
)

func fibonacci(c , quit chan int ){
    fmt.Println("fibonacci started")
    x, y := 0, 1
    for{ 
        fmt.Printf("x= %d\n",x)
        select{
            case  c <- x:
                x, y =y, x+y
                case quit_value:= <-quit:
                    fmt.Printf("quit_value = %d\n", quit_value)
                    return  
        }

}

}

func main(){
    c:=make(chan int)
    q:= make(chan int)

    go func(){
        fmt.Println("Goroutine satrted")
        for i:=0; i<5;i++{
            value := <-c
            fmt.Printf("recieved %d\n", value)
        }
        q <-000
    }()
    fmt.Println("calling fibonacci")
    fibonacci(c,q)
}

// ------------------------------------------------------------------------------------------------

// package main

// import (
//    "fmt"
// )
// // recursive 
// func fibo(n int )int {
//     if n<=1{
//         return n
//     }
//      return fibo(n-1) + fibo(n-2)
// }

// // iterative using closure
// func fibon() func () int{
//     x, y :=0,1
//     return func()int{
//         r := x
//         x, y = y, x+y
//         return r
//     }
// }

// // iterative using channel
// func fibonacci (n int, c chan int ){
//     x, y :=0,1
//     for i:= 0; i<n;i++{
//         c<-x
//         x, y  = y, x+y
//     }
//     close(c)

    
// }

// func main(){
//     n:=10
//     for i:=0;i<n;i++{
//         fmt.Println("???????? %d",fibo(i))
//     }
//     fmt.Println()
    
//     nextFibbo:= fibon()
//     for i:= 0; i<n;i++{
//         fmt.Println("*******%d",nextFibbo())

//     }
//     fmt.Println()
//     c:= make(chan int ,n)
//     go fibonacci(cap(c),c)
//     for i:=range c{
//         fmt.Println("---------%d",i)
//     }
//     fmt.Println()

// }