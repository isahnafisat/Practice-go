package main

import "fmt"

func CountDownSteps(n int) int {
	count := 0
	if n < 0 {
		return -1
	}
	if n == 0 {
		return 0
	}
	for n != 0 {
		count++
		n = n - 1
	}
	return count
}
func main() {
	fmt.Println(CountDownSteps(5))  // 5
	fmt.Println(CountDownSteps(1))  // 1
	fmt.Println(CountDownSteps(0))  // 0
	fmt.Println(CountDownSteps(-3)) // -1
}
