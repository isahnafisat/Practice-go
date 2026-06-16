package main

import "fmt"

func CountVowels(s string) int {
	count := 0
	for _, c := range s {
		if isLowerCaseVowel(c) || isUpperCaseVowel(c) {
			count++
		}
	}
	return count
}
func isLowerCaseVowel(c int32) bool {
	return (c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u')
}
func isUpperCaseVowel(c int32) bool {
	return (c == 'A' || c == 'E' || c == 'I' || c == 'O' || c == 'U')
}
func main() {
	fmt.Println(CountVowels("hello"))
	fmt.Println(CountVowels("AEIOU"))
	fmt.Println(CountVowels("sky"))
}
