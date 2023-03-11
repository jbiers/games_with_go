package main

import "fmt"

func main() {
	fmt.Println("Heyy")

	func() {
		fmt.Println("sjfdf")
	}()
}
