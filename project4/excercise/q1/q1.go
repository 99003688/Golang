package main

import (
	"fmt"
	// "time"
	"sync"
)
var wait sync.WaitGroup
func square(i int, squarechan chan int) {
	fmt.Println("hello go routine 1 is going to sleep")
	// time.Sleep(4 * time.Nanosecond)
	fmt.Println("hello  go  routine 1 awake and going to print square")
	sq:=i*i
    squarechan <- sq
    // wait.Wait()
}
func cube(i int, cubechan chan int){
	fmt.Println("hello go routine 2 is going to sleep")
	// time.Sleep(4 * time.Nanosecond)
	fmt.Println("hello  go routine 2 awake and going to print cube")

    c:=i*i*i
	cubechan <- c
}
func main(){
	var i int 
	wait.Add(2)
	fmt.Scanf("%d",&i)
	sq:= make(chan int)
	c:= make(chan int)
    go square(i,sq)
	go cube(i,c)
    square,cube:=<-	sq,<-c
	fmt.Println(square,cube)
	fmt.Println(square+cube)



}