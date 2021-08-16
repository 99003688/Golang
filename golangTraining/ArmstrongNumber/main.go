package main

import(
	"fmt"

)

func main(){
	var result, number int
	fmt.Println("enter the number you want to check")
	fmt.Scan("%d",&number)

	tempNumber := number

	for {
		remainder := tempNumber%10
		result += remainder*remainder*remainder		
		tempNumber /=10
		
		if(tempNumber==0){
			break  
		}
	}

	if(result==number){
		 fmt.Printf("%d is an Armstrong number.",number)
	}else{
		fmt.Printf("%d is not an Armstrong number.",number)
	}
}

