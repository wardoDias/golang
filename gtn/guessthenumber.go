package main

import (
	"bufio"
	"fmt"
	"math/rand/v2"
	"os"
	"strconv"
	"strings"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)

	const minVal, maxVal, maxAttempts int = 1, 50, 5
	gameFinished := false

GameLoop:
	for !gameFinished {
		attempts := maxAttempts
		secretNum := rand.IntN(maxVal-minVal+1) + minVal

		guessed := false

		fmt.Printf("I have guessed a number from %d to %d.\n", minVal, maxVal)
		fmt.Printf("Guess it in %d tries.\n", maxAttempts)

		for !guessed && attempts > 0 {
			fmt.Printf("You have %d attempts left.\n", attempts)
			fmt.Print("Your Guess: ")

			if !scanner.Scan() {
				if scanner.Err() != nil {
					fmt.Println("Error reading input")
				}
				fmt.Println("Exiting the game. Deuces!")
				break GameLoop
			}
			userInput := scanner.Text()
			userNum, err := strconv.Atoi(strings.TrimSpace(userInput))

			if err != nil {
				fmt.Println("Invalid input, please enter a number")
				continue
			}

			if userNum == secretNum {
				fmt.Printf("Congratulations! You guessed the secret number, which is %d\n", secretNum)
				guessed = true

			} else {

				if userNum < minVal || userNum > maxVal {
					fmt.Printf("Out of range. No attempt point charges for this. Your number should be in %d-%d range\n", minVal, maxVal)
					continue
				}

				attempts--

				if attempts == 0 {
					fmt.Printf("Game Over! The secret number was %d\n", secretNum)
				} else if userNum < secretNum {
					fmt.Println("Higher...")
				} else {
					fmt.Println("Lower")
				}

			}

		}
		fmt.Println("Wanna play again? (y/n)")
		if scanner.Scan() && strings.ToLower(strings.TrimSpace(scanner.Text())) == "y" {
			fmt.Println("Restarting the game...")
		} else {
			fmt.Println("Thanks for playing! Goodbye!")
			gameFinished = true
		}

	}
}
