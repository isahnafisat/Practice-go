package main

import "fmt"

func CheckVowel(s string) bool {
	for _, ch := range s {
		// fmt.Printf("type of %x is %T", ch, ch)
		if isLowerCaseVowel(ch) || isUpperCaseVowel(ch) {
			return true
		}
	}
	return false
}

func isLowerCaseVowel(ch int32) bool {
	return (ch == 'a' || ch == 'e' || ch == 'i' || ch == 'o' || ch == 'u')
}

func isUpperCaseVowel(ch int32) bool {
	return (ch == 'A' || ch == 'E' || ch == 'I' || ch == 'O' || ch == 'U')
}

func main() {
	fmt.Println(CheckVowel("sky"))
	fmt.Println(CheckVowel("hello"))
	fmt.Println(CheckVowel("WHY"))
}
