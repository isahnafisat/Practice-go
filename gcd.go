package main

import "fmt"

func Gcd(a, b uint) uint {
	if a == 0 || b == 0 {
		return 0
	}
	start := min(a, b)
	for i := start; i >= 1; i-- {
		if a%i == 0 && b%i == 0 {
			return i
		}
	}
	return 1
}
func main() {
	fmt.Println(Gcd(42, 10))
	fmt.Println(Gcd(42, 12))
	fmt.Println(Gcd(14, 77))
	fmt.Println(Gcd(17, 3))
}
