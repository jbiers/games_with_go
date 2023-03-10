// add a function to add a page after a given page
// add a function to delete a page

package main

import (
	"bufio"
	"fmt"
	"os"
)

type bookPage struct {
	text     string
	nextPage *bookPage
}

func playStory(page *bookPage) {
	scanner := bufio.NewScanner(os.Stdin)

	fmt.Println(page.text)
	scanner.Scan()

	if page.nextPage != nil {
		playStory(page.nextPage)
	}
}

func addPage(prevPage *bookPage, text string) {
	newPage := bookPage{text, prevPage.nextPage}

	prevPage.nextPage = &newPage
}

func removePage(page *bookPage) {
	page.text = page.nextPage.text
	page.nextPage = page.nextPage.nextPage
}

func main() {
	page1 := bookPage{"It was a dark and stormy night.", nil}
	page2 := bookPage{"You are alone, and you need to find the sacred helmet.", nil}
	page3 := bookPage{"You see a troll ahead.", nil}

	page1.nextPage = &page2
	page2.nextPage = &page3

	addPage(&page3, "The place smelled of foulness.")
	removePage(&page3)

	playStory(&page1)
}
