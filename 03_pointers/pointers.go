package main

import "fmt"

type position struct {
	x float32
	y float32
}

type badGuy struct {
	name   string
	health int
	pos    position
}

func whereIsBadGuy(b *badGuy) {
	x := b.pos.x
	y := b.pos.y

	fmt.Println("x:", x, "y:", y)

}

func addOne(x *int) {
	*x = *x + 1
}

func main() {
	x := 5
	fmt.Println(x)

	// xPtr is of type pointer to an int
	var xPtr *int = &x
	fmt.Println(xPtr)

	addOne(xPtr)
	fmt.Println(x)

	p := position{3, 6}
	goomba := badGuy{"Goomba", 100, p}

	whereIsBadGuy(&goomba)

}
