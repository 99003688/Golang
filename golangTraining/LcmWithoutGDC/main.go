package main
import "fmt"

/* Function without return declaration*/
func main() { 
	var n1,n2,temp1,temp2 int
    fmt.Println("Enter two positive integers : ")
    fmt.Scanln(&n1)
    fmt.Scanln(&n2)
	
    var lcmnum int =1
    if(temp1>temp2)    {
        lcmnum=temp1
    }else{
        lcmnum=temp2
    }
    /* Use of For Loop as a While Loop*/
    for {        
        if(lcmnum%temp1==0 && lcmnum%temp2==0) {    // And operator
            /*  Print Statement with multiple variables   */
            fmt.Printf("LCM of %d and %d is %d",temp1,temp2,lcmnum)            
            break
        }
        lcmnum++
    }
	fmt.Printf("the lcm of two positive numbber is %d",lcmnum)
   
}