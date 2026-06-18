package main

import "fmt"

func Swap(a, b int) (int, int) {
	return b, a
}
func main() {
	fmt.Println(Swap(1, 2))
	fmt.Println(Swap(10, 7))
	fmt.Println(Swap(-3, 5))
}
