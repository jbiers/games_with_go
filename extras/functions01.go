package main

import "fmt"

func returnInt() int {
	return 90
}

func returnIntString() (int, string) {
	return 50, "hello"
}

func main() {
	fmt.Println(returnInt())
	fmt.Println(returnIntString())
}
