package main

import "fmt"

func ThirdLetter(s string) string {
	if len(s) < 1 {
		return "\n"
	}
	return string(s[len(s)-1]) + "\n"
}
func main() {
	fmt.Println(ThirdLetter("hello"))
	fmt.Println(ThirdLetter("Go"))
	fmt.Println(ThirdLetter("A"))
	fmt.Println(ThirdLetter(""))
	fmt.Println(ThirdLetter("ChatGPT"))
}
