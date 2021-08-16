package main

import "fmt"

func main() {

    var yr int

    fmt.Print("\nEnter the Year to Check for Leap = ")
    fmt.Scanln(&yr)

    if yr%400 == 0 || (yr%4 == 0 && yr%100 != 0) {
        fmt.Println(yr, " is a Leap Year")
    } else {
        fmt.Println(yr, " is Not a Leap Year")
    }
}