package main

import (
	"fmt"
)

func factorial(number int) int {
	var result int = 1
	for i := 2; i <= number; i++ {
		result = result * i
	}
	return result

}

func combination(num1 int, num2 int) int {
	return factorial(num1) / (fact(num2) * fact(num1-num2))
}

func main() {
	var num1, num2 int
	fmt.Println("Enter the value of n and r")
	fmt.Scanln(&num1)
	fmt.Scanln(&num2)
	fmt.Println(combination(num1, num2))

}
