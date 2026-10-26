package main

import (
	"fmt"
	"math/rand/v2"
)

func main() {
	secretNum := rand.IntN(50) + 1
	attempts := 5
	guessed := false

	var userNum int

	fmt.Println("I have guessed a number from 1 to 50.")
	fmt.Println("Guess it in 5 tries.")

	for !guessed && attempts > 0 {
		fmt.Printf("You have %d attempts left.\n", attempts)
		fmt.Print("Your Guess: ")
		_, err := fmt.Scan(&userNum)

		if err != nil {
			fmt.Println("Invalid input, please enter a number")
			continue
		}

		if userNum == secretNum {
			fmt.Println("Congratulations! Your guessed the secret number")
			guessed = true

		} else {
			attempts--

			if attempts == 0 {
				fmt.Printf("Game Over! The secret number was %d", secretNum)
			} else if userNum < secretNum {
				fmt.Println("Higher...")
			} else {
				fmt.Println("Lower")
			}

		}

	}
}
