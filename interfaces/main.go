package main

import (
	"fmt"
  s "interfaces/new"

)

// type Vowel interface
// 	IsVowel(c rune) bool
// }

// type Vow string
func main() {

	var v1 s.Vowel
	v1 = s.Vow("saujanya")
	t := 0

	for _, value := range  s.Vow("saujanya") {

		if v1.IsVowel(value) {

			fmt.Printf("%c is Vowel\n", value)
			t++
		}
	}
	fmt.Printf("%d Vowels.", t)

}
