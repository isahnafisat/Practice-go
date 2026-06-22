package main

import "fmt"

func FirstLetter(s string) string {
	if s == "" {
		return "\n"
	}
	return string(s[0]) + "\n"
}
func main() {
	fmt.Println(FirstLetter("hello"))
	fmt.Println(FirstLetter("Go"))
	fmt.Println(FirstLetter(""))
	fmt.Println(FirstLetter("ChatGPT"))
}
