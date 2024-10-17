package main

import (
	"fmt"
	"math/rand"
	"strings"
	"time"

	"github.com/eiannone/keyboard"
)

func main() {
	// ANSI escape codes for colors
	const (
		reset  = "\033[0m"
		red    = "\033[31m"
		green  = "\033[32m"
		yellow = "\033[33m"
		blue   = "\033[34m"
	)

	// The strings to display
	words := []string{"apple", "orange", "pear"}

	// Calculate the length of the longest text
	maxLength := max(len(words[0]), len(words[1]), len(words[2]))

	// Create the top parts of the boxes
	top := blue + "┌" + strings.Repeat("─", maxLength+2) + "┐" + reset

	// Create the bottom parts of the boxes
	bottom := blue + "└" + strings.Repeat("─", maxLength+2) + "┘" + reset

	// Initialize win and loss counters
	wins := 0
	bigWins := 0
	losses := 0
	credits := 0
	bank := 100

	// Function to display the boxes with random words
	displayBoxes := func() (string, string, string) {
		rand.Seed(time.Now().UnixNano())
		word1 := words[rand.Intn(len(words))]
		word2 := words[rand.Intn(len(words))]
		word3 := words[rand.Intn(len(words))]

		middle1 := blue + "│ " + centerText(word1, maxLength) + " │" + reset
		middle2 := blue + "│ " + centerText(word2, maxLength) + " │" + reset
		middle3 := blue + "│ " + centerText(word3, maxLength) + " │" + reset

		fmt.Println(top + " " + top + " " + top)
		fmt.Println(middle1 + " " + middle2 + " " + middle3)
		fmt.Println(bottom + " " + bottom + " " + bottom)

		return word1, word2, word3
	}

	// Display the boxes initially
	// Clear the screen
	fmt.Print("\033[H\033[2J")
	fmt.Println()
	displayBoxes()
	fmt.Printf("Wins: %d, Big Wins: %d, Losses: %d, Credits: %d, Bank: %d\n", wins, bigWins, losses, credits, bank)

	// Initialize the keyboard
	if err := keyboard.Open(); err != nil {
		fmt.Println("Failed to open keyboard:", err)
		return
	}
	defer keyboard.Close()

	for {
		fmt.Print("Enter 's' to spin, 'c' to cash out, 'i' to cash in, or 'q' to quit: ")
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			fmt.Println("Failed to read keyboard input:", err)
			return
		}

		if char == 's' {
			if credits <= 0 {
				fmt.Println("Insufficient credits! Please cash in from the bank.")
			} else {
				// Clear the screen
				fmt.Print("\033[H\033[2J")

				// Print the boxes side-by-side
				fmt.Println()
				word1, word2, word3 := displayBoxes()
				if word1 == word2 && word2 == word3 {
					wins++
					bigWins++
					credits += 10
					fmt.Println("You win 10 credits!")
				} else if word1 == word2 || word2 == word3 || word1 == word3 {
					wins++
					credits++
					fmt.Println("You win 1 credit!")
				} else {
					losses++
					credits--
					fmt.Println("You lose!")
				}
				fmt.Printf("Wins: %d, Big Wins: %d, Losses: %d, Credits: %d, Bank: %d\n", wins, bigWins, losses, credits, bank)
			}
		} else if char == 'c' {
			// Cash out credits to bank
			bank += credits
			credits = 0
			fmt.Println("Cashed out!")
			fmt.Printf("Wins: %d, Big Wins: %d, Losses: %d, Credits: %d, Bank: %d\n", wins, bigWins, losses, credits, bank)
		} else if char == 'i' {
			// Cash in from bank to credits
			if bank > 0 {
				amount := 10 // Amount to cash in
				if bank < amount {
					amount = bank
				}
				bank -= amount
				credits += amount
				fmt.Printf("Cashed in %d credits!\n", amount)
			} else {
				fmt.Println("Bank is empty!")
			}
			fmt.Printf("Wins: %d, Big Wins: %d, Losses: %d, Credits: %d, Bank: %d\n", wins, bigWins, losses, credits, bank)
		} else if char == 'q' {
			break
		}
		fmt.Println() // Output a carriage return and newline
	}
}

// Helper function to find the maximum of three integers
func max(a, b, c int) int {
	if a > b {
		if a > c {
			return a
		}
		return c
	}
	if b > c {
		return b
	}
	return c
}

// Helper function to center text within a given width
func centerText(text string, width int) string {
	padding := (width - len(text)) / 2
	return strings.Repeat(" ", padding) + text + strings.Repeat(" ", width-len(text)-padding)
}
