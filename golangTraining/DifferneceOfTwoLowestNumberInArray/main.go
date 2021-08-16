package main

import (
	"fmt"
	"sort"
)

func main(){
	var n int
	fmt.Println("Enter the number of element in array")
	fmt.Scan("%d",&n)
	var arr = make([]int, n)
	for i:=0; i<n; i++ {
	   fmt.Printf("Enter %dth element: ", i)
	   fmt.Scanf("%d", &arr[i])
	}
	fmt.Println("Your array is: ", arr)
	sort.Sort(sort.IntSlice(arr))
	minimumdiffernece:=arr[1]-arr[0]
	fmt.Printf("the minimum difference in array is %d",minimumdiffernece)
}