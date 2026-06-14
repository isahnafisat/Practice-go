package main

import "fmt"

func CountBoth(s string) (int, int) {
	lettercount := 0
	numbercount := 0
	for _, c := range s {
		if (c >= 'a' && c <= 'z') || (c >= 'A' && c <= 'Z') {
			lettercount++
		}
		if c >= '0' && c <= '9' {
			numbercount++
		}
	}
	return lettercount, numbercount
}
func main() {
	l1, n1 := CountBoth("H1e2l3lo")
	fmt.Println("letters:", l1, "numbers:", n1)

	l2, n2 := CountBoth("Hello123")
	fmt.Println("letters:", l2, "numbers:", n2)

	l3, n3 := CountBoth("Go Lang 99")
	fmt.Println("letters:", l3, "numbers:", n3)
}
