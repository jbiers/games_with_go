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

func whereIsBadGuy(b badGuy) {
	x := b.pos.x
	y := b.pos.y

	fmt.Println("x:", x, "y:", y)

}

func main() {
	p := position{3, 6}

	goomba := badGuy{"Goomba", 100, position{2, 34}}

	fmt.Println(p)
	fmt.Println(goomba)

	whereIsBadGuy(goomba)
}
