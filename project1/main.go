package main

import (
	"log"
	"net/http"
	"project1/controller/productcontroller"
	
)

func main(){
     mux:=http.NewServeMux()
	mux.HandleFunc("/", productcontroller.Index)
	mux.HandleFunc("/product", productcontroller.Index)
	mux.HandleFunc("/product/index", productcontroller.Index)
	mux.HandleFunc("/product/add", productcontroller.Add)
	mux.HandleFunc("/product/addprocess", productcontroller.ProcessAdd)
	log.Fatal(http.ListenAndServe(":8085",mux))


	// -----------------------------

	

    
    
	
// 	var number int
// 	// var v int
// 	// var minim int
// 	// fmt.Println("enter the number of elements ")
// 	fmt.Scanf("%d",&number)
//    var arr = make([]int, number)
// //   for i:=0; i<number; i++ {
// // //   fmt.Printf("Enter %dth element: ", i)
// //   fmt.Scanf("%d", &arr[i])
// // }
// smallestNumber := arr[0]  
// for i, element := range arr {  
// 	fmt.Scanf("%d", &arr[i])
// 	fmt.Println("arr at 0",arr[i])

// if arr[i] < smallestNumber {  
// smallestNumber = element  

// }  
// }
// fmt.Println(smallestNumber)   

 

	

// }
    /*
// Sample code to perform I/O:

fmt.Scanf("%s", &myname)            // Reading input from STDIN
fmt.Println("Hello", myname)        // Writing output to STDOUT

// Warning: Printing unwanted or ill-formatted data to output will cause the test cases to fail
*/

// Write your code here

}