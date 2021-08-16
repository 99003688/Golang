package main

import(
	"fmt"
)

func main(){
	var num1, num2, num3, num4 int
	fmt.Println("enter four numbers ")
	fmt.Scan("%d %d %d %d",&num1,&num2,&num3,&num4)

	fmt.Printf("The Entered three numbers are %d %d %d %d\n", num1, num2, num3, num4)  
  
 if num1 >= num2 && num1 >= num3 && num1>= num4 {  
  fmt.Println("Largest of three numbers: ", num1)  
 } else if num2 >= num1 && num2 >= num3 && num2>= num4 {  
  fmt.Println("Largest of three numbers: ", num2)  
 }  else if num3 >= num1 && num3 >= num2 && num3>= num4 {  
	fmt.Println("Largest of three numbers: ", num3)
 }else {  
  fmt.Println("Largest of three numbers: ", num4)  
 }  

}