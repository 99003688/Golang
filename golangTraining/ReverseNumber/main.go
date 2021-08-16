package main
import "fmt"
func reverseNumber(num int) (int , int){

   res := 0
   sum:=0
   for num>0 {
      remainder := num % 10
      res = (res * 10) + remainder
      num /= 10
	  sum = sum+remainder
	  }
   return res,sum
}

func main(){
   var n int
   fmt.Scanf("%d",&n)
   fmt.Println(reverseNumber(n))
  
}