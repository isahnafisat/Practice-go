package main

import "fmt"

func FirstWord(s string) string {
	result := ""
	for i := 0; i < len(s); i++ {
		if s[i] == ' ' {
			break
		}
		result = result + string(s[i])
	}
	return result + "\n"
}
func main() {
	fmt.Println(FirstWord("hello there"))
	fmt.Println(FirstWord(""))
	fmt.Println(FirstWord("hello   .........  bye"))
}
