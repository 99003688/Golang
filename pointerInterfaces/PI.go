// package main

// import "fmt"

// package main

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
// type Bird interface {
// 	fly()
// }

// type B struct {
// 	name string
// }

// // implement it
// func (b *B) fly() {
// 	fmt.Println("Flying...",b.name)
// }

// func main() {

// 	var b Bird
// 	a := B{"Peacock"}

// 	b = &a

// 	// b := &a
// 	// fmt.Println(b)    // 0x40c138
// 	// fmt.Println(*b)   // {Peacock}
// 	b.fly()
// 	a.fly()
// }

// type saujanya interface {
// 	empid(int)
// }

// func empid(s saujanya, id int) {
// 	s.empid(id)
// }

// type details struct {
// 	emp int
// }

// func (c details) empid(id int) {
// 	c.emp = id
// }

// type number struct {
// 	ps int
// }

// func (p *number) empid(id int) {
// 	p.ps = id
// }

// func main() {

// 	first := details{emp: 2499}
// 	second := number{ps: 9999}

// 	empid(first, 1999)

// 	empid(&second, 7999)

// 	fmt.Println("employee  id:", first.emp)
// 	fmt.Println("personal number:", second.ps)
// }

      
package main  
import (  
   "fmt"  
   "net/http"  
)  
func main() {  
   http.HandleFunc("/",MyHandler1)  
   http.HandleFunc("/John",MyHandler2)  
   http.ListenAndServe(":8080",nil)  
}  
func MyHandler1(w http.ResponseWriter,r *http.Request){  
   fmt.Fprint(w,"Hello World\n")  
}  
func MyHandler2(w http.ResponseWriter,r *http.Request){  
   fmt.Fprint(w,"Hello John\n")  
} 
