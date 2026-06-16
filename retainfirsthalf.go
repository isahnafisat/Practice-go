package main

import "fmt"

func RetainFirstHalf(str string) string {
	half := len(str) / 2
	if len(str) == 0 {
		return ""
	}
	if len(str) == 1 {
		return str
	}
	return str[:half]
}
func main() {
	fmt.Println(RetainFirstHalf("This is the 1st halfThis is the 2nd half"))
	fmt.Println(RetainFirstHalf("A"))
	fmt.Println(RetainFirstHalf(""))
	fmt.Println(RetainFirstHalf("Hello World"))
}
