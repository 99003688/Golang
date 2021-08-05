package main

import "fmt"

type Vowel interface {
	IsVowel(c rune) bool

}

type GroupChars []rune

type Vow string

func (v Vow) IsVowel(c rune) bool {
	vowels := GroupChars{'a', 'e', 'i', 'o', 'u', 'A', 'E', 'I', 'O', 'U'}
	for _, value := range vowels {
		if value == c {
			return true
		}
	}
	return false
}


func main() {

	var v1 Vowel
	v1 = Vow("saujanya")
	t := 0

	for _, value := range Vow("saujanya") {
		if v1.IsVowel(value) {
			fmt.Printf("%c is Vowel\n", value)
			t++
		}
	}
	fmt.Printf("%d Vowels.", t)

}
