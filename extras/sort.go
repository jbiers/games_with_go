package main

import (
	"fmt"
	"sort"
)

type Cat struct {
	name       string
	tailLength int
}

type byTailLength []Cat

func (b byTailLength) Len() int {
	return len(b)
}

func (b byTailLength) Less(i, j int) bool {
	return b[i].tailLength < b[j].tailLength
}

func (b byTailLength) Swap(i, j int) {
	b[i], b[j] = b[j], b[i]
}

func main() {
	cats := []Cat{
		{"Lola", 40},
		{"Fifo", 33},
		{"Bartô", 37},
		{"Maria", 27},
		{"Risoto", 10},
		{"Zé Lelé", 30},
	}

	fmt.Println(cats)

	sort.Sort(byTailLength(cats))
	fmt.Println(cats)
}
