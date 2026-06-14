package main

import "fmt"

func CheckUpper(arg string) bool {
	for _, ch := range arg {
		if ch >= 'A' && ch <= 'Z' {
			return true
		}
	}
	return false
}
func main() {
	fmt.Println(CheckUpper("hello"))
	fmt.Println(CheckUpper("heLlo"))
	fmt.Println(CheckUpper("GO"))
}
