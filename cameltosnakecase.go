package main

import "fmt"

func CamelToSnakeCase(s string) string {
	result := ""

	for _, ch := range s {
		if ch >= '0' && ch <= '9' {
			return s
		}
	}
	for i, ch := range s {
		if ch >= 'A' && ch <= 'Z' {
			if i != 0 {
				result += "_"
			}
		}
		result += string(ch)
	}
	return result
}
func main() {
	fmt.Println(CamelToSnakeCase("HelloWorld"))
	fmt.Println(CamelToSnakeCase("helloWorld"))
	fmt.Println(CamelToSnakeCase("camelCase"))
	fmt.Println(CamelToSnakeCase("CAMELtoSnackCASE"))
	fmt.Println(CamelToSnakeCase("camelToSnakeCase"))
	fmt.Println(CamelToSnakeCase("hey2"))
}
