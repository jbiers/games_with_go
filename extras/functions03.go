package main

import "fmt"

type Person struct {
	name     string
	lastName string
	age      int
}

func (p Person) PresentYourself() {
	fmt.Printf("I'm %v %v, %v years old.\n", p.name, p.lastName, p.age)
}

func main() {
	me := Person{
		name:     "Julia",
		lastName: "Bier",
		age:      21,
	}

	me.PresentYourself()
}
