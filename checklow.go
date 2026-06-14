package main

import "fmt"

func CheckLower(s string) bool {
	for _, ch := range s {
		if ch >= 'a' && ch <= 'z' {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(CheckLower("HELLO"))
	fmt.Println(CheckLower("HeLLO"))
	fmt.Println(CheckLower("123"))
}
