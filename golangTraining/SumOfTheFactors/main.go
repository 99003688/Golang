package main

import "fmt"

func main() {

    var factorsnum, i,sum int
	 sum =0

    fmt.Print("Enter any Number to find Factors = ")
    fmt.Scanln(&factorsnum)

    fmt.Println("The Factors of the ", factorsnum, " are = ")
    for i = 1; i <= factorsnum; i++ {
        if factorsnum%i == 0 {
            fmt.Println(i)
		    sum = sum + i
			
        }
		
}
fmt.Println("sum of the factors are",sum)
    }