package main

import "fmt"

func CountSpaces(s string) int {
	count := 0
	for _, c := range s {
		if c == ' ' {
			count++
		}
	}
	return count
}
func main() {
	fmt.Println(CountSpaces("Hello world"))
	fmt.Println(CountSpaces("H e l l o"))
	fmt.Println(CountSpaces("NoSpaces"))
}
