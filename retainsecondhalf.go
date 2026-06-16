package main

import "fmt"

func RetainSecondHalf(str string) string {
	half := len(str) / 2
	if len(str) == 0 {
		return " "
	}
	if len(str) == 1 {
		return str
	}
	return str[half:]
}
func main() {
	fmt.Println(RetainSecondHalf("HelloWorld"))
	fmt.Println(RetainSecondHalf("abcdef"))
	fmt.Println(RetainSecondHalf("A"))
	fmt.Println(RetainSecondHalf(""))
}
