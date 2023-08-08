package main

import (
	"fmt"
	"math/rand"
)

const myNumber = 43

func main() {
	computerGuess := rand.Intn(101)
	count := 1
	for computerGuess != myNumber {
		if computerGuess > myNumber {
			fmt.Printf("Computer's Guess: %v. Too big!\n", computerGuess)

		} else {
			fmt.Printf("Computer's Guess: %v. Too small!\n", computerGuess)
		}

		// Let computer guess again
		count++
		computerGuess = rand.Intn(101)
	}
	fmt.Printf("Yay! Computer has successfully guessed my number after %v tries! Brute force win win!!🥇\n", count)

}
