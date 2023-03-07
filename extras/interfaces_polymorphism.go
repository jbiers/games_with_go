package main

import (
	"fmt"
)

type Person struct {
	name     string
	lastName string
	age      int
}

type Dentist struct {
	Person
	fixedTeeth int
	salary     int
}

type Architect struct {
	Person
	constructionType string
	howCrazy         int
}

func (x Dentist) helloGoodmorning() {
	fmt.Println("Hello, goodmorning! My name is ", x.name, ". I'm a dentist.")
	fmt.Println("I have fixed ", x.fixedTeeth, " teeth.")
}

func (x Architect) helloGoodmorning() {
	fmt.Println("Hello, goodmorning! My name is ", x.name, ". I'm an architect.")
	fmt.Println("I work with", x.constructionType, "!")
}

type Guy interface {
	helloGoodmorning()
}

func humanBeing(g Guy) {
	g.helloGoodmorning()
}

func main() {
	mrTooth := Dentist{
		Person: Person{
			"Dino",
			"Tooth",
			42,
		},
		fixedTeeth: 695,
		salary:     1000000,
	}

	mrBuilding := Architect{
		Person: Person{
			"Thomas",
			"Building",
			49,
		},
		constructionType: "urbanism",
		howCrazy:         10,
	}

	mrBuilding.helloGoodmorning()
	mrTooth.helloGoodmorning()

	humanBeing(mrBuilding)

}

// https://gobyexample.com/interfaces
