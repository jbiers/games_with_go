package main

import (
	"bufio"
	"fmt"
	"os"
)

func main() {

	scanner := bufio.NewScanner(os.Stdin)

	highestNumber := 100
	lowestNumber := 1

	numGuesses := 0
	currentGuess := (highestNumber + lowestNumber) / 2

	fmt.Printf("You need to think of a number between %d and %d...\n", lowestNumber, highestNumber)
	fmt.Println("And I will try to guess it!")

	for {
		fmt.Printf("%d? (higher/lower/yes) ", currentGuess)
		scanner.Scan()

		if numGuesses >= 100 {
			fmt.Println("I think you tricked me...")
			break
		}

		if scanner.Text() == "yes" {
			numGuesses++
			fmt.Printf("It took me %d guesses to get to %d!\n", numGuesses, currentGuess)
			break
		} else if scanner.Text() == "higher" {
			lowestNumber = currentGuess + 1

			numGuesses++
			currentGuess = (highestNumber + lowestNumber) / 2
			continue
		} else if scanner.Text() == "lower" {
			highestNumber = currentGuess - 1

			numGuesses++
			currentGuess = (highestNumber + lowestNumber) / 2
			continue
		} else {
			fmt.Println("Invalid input. Try again.")
		}
	}

}
