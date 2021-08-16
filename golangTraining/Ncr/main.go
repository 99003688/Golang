package main

import (
	"fmt"
)

func fact(n int) int {
	var res int = 1
	for i := 2; i <= n; i++ {
		res = res * i
	}
	return res

}

func nCr(n int, r int) int {
	return fact(n) / (fact(r) * fact(n-r))
}

func main() {
	var n, r int
	fmt.Println("Enter the value of n and r")
	fmt.Scanln(&n)
	fmt.Scanln(&r)
	fmt.Println(nCr(n, r))

}
