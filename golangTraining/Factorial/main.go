package main
import "fmt"

var factVal int = 1  
var i int = 1
var n int


func factorial(n int) int {   
    if(n < 0){
        fmt.Print("Factorial of negative number doesn't exist.")    
    }else{        
        for i:=1; i<=n; i++ {
            factVal *= int(i)  
        }
        
    }    
    return factVal 
}

func main(){    
    fmt.Print("Enter a positive integer between 0 - 50 : ")
    fmt.Scan(&n)   
    fmt.Print("Factorial is: ",factorial(n))
}