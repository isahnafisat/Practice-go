package main

import "fmt"

func main() {
	for {
		var hunger int

		fmt.Print("Enter hunger level(1-10, 0 to exit): ")
		fmt.Scan(&hunger)

		if hunger == 0 {
			fmt.Println("Goodbye 👋")
			break
		}
		if hunger > 8 {
			fmt.Println("Big meal 🍔")
		} else if hunger > 4 {
			fmt.Println("Snack 🍕")
		} else {
			fmt.Println("Water 💧")
		}
	}
}
