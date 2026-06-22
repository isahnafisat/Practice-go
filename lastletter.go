package main

import "fmt"

func LastLetter(s string) string {
	if s == "" {
		return "\n"
	}
	return string(s[len(s)-1]) + "\n"
}
func main() {
	fmt.Print(LastLetter("Nafisat"))
	fmt.Print(LastLetter("Isah"))
	fmt.Print(LastLetter("I love me "))
	fmt.Print(LastLetter("Okay"))
}
