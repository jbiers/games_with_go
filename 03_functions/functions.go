package main

import "fmt"

func sayHello(name string) {
	fmt.Println("Hello", name)
}

func sayGoodbye(name string) {
	fmt.Println("Goodbye", name)
}

func beNice(name string) {
	sayHello(name)
	sayGoodbye(name)
}

func addThree(number int) int {
	return number + 3
}

func main() {
	beNice("Julia")

	x := 902
	x = addThree(addThree(x))
	fmt.Println(x)
}
