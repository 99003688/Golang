package main

import (
	"fmt"
	// "strings"
	"time"

// "sync"
)
func count(str string){
	for i:=1; i<6;i++{
		fmt.Println(i,str)
		time.Sleep(time.Millisecond * 500)
	}
}
// func itarr(arr[] int){

// for i := 0; i < len(arr); i++ {
// 	fmt.Println(arr[i])
// }
// }

 func main(){
// var wg sync.WaitGroup
// wg.Add(1)
// go func (){
// 	count("sheep")
// 	wg.Done()
// }()
//  go count("fish")

// wg.Wait()

	go count("concurrency or go routine")
	 count("mapping")


//  var arr=[]int {1,2,3,4,5}
  
//  go itarr(arr)

//  itarr(arr)

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


}

