package main

import "fmt"

func IsNegative(n int) bool {
	return n < 0
}
func main() {
	fmt.Println(IsNegative(-5))
	fmt.Println(IsNegative(0))
	fmt.Println(IsNegative(10))
}
