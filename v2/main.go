package main

import (
	"fmt"
	"math/rand"
	"time"

	"github.com/eiannone/keyboard"
)

// Define the symbols for the slot machine
var symbols = []string{"🍒", "🍋", "🍊", "🍉", "⭐", "🔔"}

func main() {
	rand.Seed(time.Now().UnixNano())
	playAgain := 'y'
	coins := 0

	// Ask the player to insert coins
	fmt.Print("Insert coins to play: ")
	fmt.Scan(&coins)

	for playAgain == 'y' && coins > 0 {
		// Deduct a coin for each play
		coins--

		// Generate random symbols for each slot
		slot1 := symbols[rand.Intn(len(symbols))]
		slot2 := symbols[rand.Intn(len(symbols))]
		slot3 := symbols[rand.Intn(len(symbols))]

		// Display the slot machine
		fmt.Print("\033[H\033[2J") // Clear the screen
		fmt.Println("🎰 Slot Machine 🎰")
		fmt.Printf("| %s | %s | %s |\n", slot1, slot2, slot3)

		// Check for winning combinations
		if slot1 == slot2 && slot2 == slot3 {
			fmt.Println("🎉 You win! 🎉")
		} else {
			fmt.Println("😢 Try again! 😢")
		}

		// Show the number of coins left
		fmt.Printf("Coins left: %d\n", coins)

		// Check if the player has coins left
		if coins > 0 {
			// Ask the player if they want to play again
			fmt.Print("Play again? (y/n): ")
			if char, _, err := keyboard.GetSingleKey(); err == nil {
				playAgain = char
			} else {
				fmt.Println("Error reading input:", err)
				break
			}
		} else {
			fmt.Println("No more coins left!")
			break
		}
	}

	fmt.Println("Thanks for playing!")
}
