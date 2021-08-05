package main

import "fmt"

func main() {

  // String Array
  fmt.Println("1. String Array : ")

  var arr [5]string
  arr[0] = "Ubuntu"
  arr[1] = "CentOS"
  arr[2] = "RedHat"
  arr[3] = "Debian"
  arr[4] = "OpenBSD"

  personal := arr[1]// initilaising Centos to personal
  fmt.Println("mydistro = ", personal) //printing personal
  fmt.Printf("\n")
  fmt.Println("arr[2] = ", arr[2])// printing RedHat
  fmt.Printf("\n")
  fmt.Println("arr = ", arr)//Printing whole arr
  fmt.Printf("\n")
  fmt.Println("Number of arr = ", len(arr))

  // Integer Array (Numbers)
  fmt.Printf("\n")
  fmt.Println("2. Integer Array : ")

  var ids [5]int
  ids[0] = 1
  ids[1] = 2
  ids[2] = 3
  ids[3] = 4
  ids[4] = 5

  myid := ids[3]
  fmt.Println("myid = ", myid)
  fmt.Println("ids[2] = ", ids[2])
  fmt.Println("ids = ", ids)
  fmt.Println("Number of ids = ", len(ids))

  // Declare and Initialize Array at the same time
  fmt.Printf("\n")
  fmt.Println("3. Declare and Initialize Array at the same time : ")

  lang := []string{"golang", "go", "arrays"}
  lang = append(lang , "saujanya")
  fmt.Printf("\n")
  fmt.Println("lang = ",lang)
  fmt.Printf("\n")
  fmt.Println("lang = ",lang[1:4])
  fmt.Printf("\n")
  fmt.Println("Number of lang = ", len(lang))
  fmt.Printf("\n")
  

  fibonacci := []int{1, 1, 2, 3, 5, 8}
  fibonacci=append(fibonacci, 1)
  fmt.Println("fibonacci = ",fibonacci)
  fmt.Println(len(fibonacci))//

  // Multi-line Array Initialization Syntax
  fmt.Println()
  fmt.Println("4. Multi-line Array values in array ")

  temperature := [3]float64{
       98.5,
       65.5,
       83.2,
  }
  fmt.Println("temperature = ", temperature)

 

  // Default Values in an Array
  fmt.Println()
  fmt.Println("5. Default Values in an Array : ")

  empIds := [5]int{101, 102, 103}
  fmt.Println("empIds = ",empIds)

  empNames := [5]string{"saujanya","tailang"}
  fmt.Println("empNames = ",empNames)
  

  // Loop through Array using For and Range
  fmt.Println()
  fmt.Println("6. Loop through Array using For and Range : ")

  for index, value := range arr {
    fmt.Println(index, " = ", value)
  }

  // Loop through Array using For and Range (Ignore Index)
  fmt.Println()
  fmt.Println("7. Loop through Array using For and Range (Ignore Index) : ")

  total := 0
  for _, value := range ids {
    total = total + value
  }
  fmt.Println("total of all ids = ", total)

  // Initialize an integer array with sequence
  fmt.Println()
  fmt.Println("8. Initialize an integer array with sequence : ")
  var sequence [10]int
  counter := 10
  for index := range sequence {
    sequence[index] = counter
    counter = counter + 5
  }
  fmt.Println("sequence = ",sequence)

  // Multi dimensional array
  fmt.Println()
  fmt.Println("9. Multi dimensional array : ")

  count := 1
  var multi [4][2]int
   for i := 0; i < 4; i++ {
       for j := 0; j < 2; j++ {
           multi[i][j] = count
           count++
       }
   }
   fmt.Println("Array 4 x 2 : ", multi)

}