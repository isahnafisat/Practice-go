package main

import "fmt"

func main() {
	var hunger int

	fmt.Print("Enter your hunger level (1-10): ")
	fmt.Scan(&hunger)

	if hunger > 8 {
		fmt.Println("Big meal 🍔")
	} else if hunger > 4 {
		fmt.Println("snack 🍕")
	} else {
		fmt.Println("Water 💧")
	}
}
