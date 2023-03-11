package main

import "fmt"

type Person struct {
	name      string
	last_name string
	age       int
}

func changeMe(p *Person) {
	p.name = "Mary"
}

func main() {
	me := Person{
		name:      "Julia",
		last_name: "Bier",
		age:       21}

	changeMe(&me)

	fmt.Println(me.name)
}
