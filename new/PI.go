package main

import "fmt"





 
// import (
//     "fmt"
// )
 
// func main() {   
//     var a interface{}
//     b := &a
//     fmt.Println(b)    // 0x40c138
//     fmt.Println(*b)   // <nil>
// }
  

// package main
 
// import (
//     "fmt"
// )
 
// // declare interface
// type saujnaya interface{
//     golang()
// }
 
// type B struct{
//     name string
// }
 
// // implement it
// func (b B)golang() {
//     fmt.Println("learning golang...")
// }
 
// func main() {   
//     var a saujanya = B{"interfaces"}
//     b := &a
//     fmt.Println(b)    // location 
//     fmt.Println(*b)   // {interfaces}
// }

type CoursePrice interface {
    show(int)
}
  

func show(cp CoursePrice, fee int) {
    cp.show(fee)
}
  

type Dsa struct {
    Price int
}
  
func (c Dsa) show(fee int) {
    c.Price = fee
}
  

type Placement struct {
    Price int
}

func (p *Placement) show(fee int) {
    p.Price = fee
}
  

func main() {
  
    first := Dsa{Price: 2499}
    second := Placement{Price: 9999}
  

    show(first, 1999)
  
    
    show(&second, 7999)
  
    fmt.Println("DSA Course Fee:", first.Price)
    fmt.Println("Placement100 Course Fee:", second.Price)
}