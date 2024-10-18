package main

import (
	"encoding/json"
	"fmt"
	"log"
	"math/rand"
	"net/http"
	"net/url"
	"strings"

	"github.com/eiannone/keyboard"
	"github.com/r3labs/sse"
)

// ANSI escape codes for colors
const (
	reset  = "\033[0m"
	red    = "\033[31m"
	green  = "\033[32m"
	yellow = "\033[33m"
	blue   = "\033[34m"
)

// Initialize win and loss counters, credits, and bank balance to zero
var wins, bigWins, losses, credits, bank int
var reels = []string{"apple", "orange", "pear"}
var maxLength = max(len(reels[0]), len(reels[1]), len(reels[2]))
var top = blue + "┌" + strings.Repeat("─", maxLength+2) + "┐" + reset
var bottom = blue + "└" + strings.Repeat("─", maxLength+2) + "┘" + reset

// Variables to store the reels between spins
var (
	reel1 = reels[0]
	reel2 = reels[1]
	reel3 = reels[2]
)

func main() {
	// Fetch the initial bank balance from the HTTP endpoint
	var err error
	bank, err = fetchBankBalance()
	if err != nil {
		log.Fatalf("Failed to fetch bank balance: %v", err)
	}

	// Print the initial screen
	printScreen()

	// Initialize the keyboard
	if err := keyboard.Open(); err != nil {
		fmt.Println("Failed to open keyboard:", err)
		return
	}
	defer keyboard.Close()

	// Start listening for SSE updates in a separate goroutine
	go listenForSSEUpdates()

	for {
		fmt.Print("Enter 's' to spin, 'c' to cash out, 'i' to cash in, '?' to refresh, or 'q' to quit: ")
		char, _, err := keyboard.GetSingleKey()
		if err != nil {
			fmt.Println("Failed to read keyboard input:", err)
			return
		}

		if char == 's' {
			if credits <= 0 {
				printScreen("Insufficient credits! Please cash in from the bank.")
			} else {
				reel1, reel2, reel3 = spinReels()
				if reel1 == reel2 && reel2 == reel3 {
					wins++
					bigWins++
					credits += 10
					printScreen("You win 10 credits!")
				} else if reel1 == reel2 || reel2 == reel3 || reel1 == reel3 {
					wins++
					credits++
					printScreen("You win 1 credit!")
				} else {
					losses++
					credits--
					printScreen("You lose!")
				}
			}
		} else if char == 'c' {
			// Cash out credits to bank
			bank += credits
			credits = 0

			// Update the bank balance via HTTP POST request
			err := updateBankBalance(bank)
			if err != nil {
				log.Fatal("Failed to update bank balance:", err)
			}
			printScreen("Cashed out credits to bank!")
		} else if char == 'i' {
			// Cash in from bank to credits
			if bank > 0 {
				amount := 10 // Amount to cash in
				if bank < amount {
					amount = bank
				}
				bank -= amount
				credits += amount

				// Update the bank balance via HTTP POST request
				err := updateBankBalance(bank)
				if err != nil {
					log.Fatal("Failed to update bank balance:", err)
				}
				printScreen(fmt.Sprintf("Cashed in %d credits!\n", amount))
			} else {
				printScreen("Bank is empty!")
			}
		} else if char == '?' {
			// Fetch the current bank balance from the HTTP endpoint
			var err error
			bank, err = fetchBankBalance()
			if err != nil {
				log.Fatalf("Failed to fetch bank balance: %v", err)
			}
			// Refresh the screen
			printScreen()
		} else if char == 'q' {
			if credits > 0 {
				printScreen(fmt.Sprintf("Abandoned game with %d credits!", credits))
			} else {
				printScreen("Thanks for playing!")
			}
			break
		} else {
			fmt.Println()
		}
	}
}

func printScreen(msg ...string) {
	fmt.Print("\033[H\033[2J")
	fmt.Println()
	displayBoxes()
	fmt.Printf("Wins: %d, Big Wins: %d, Losses: %d, Credits: %d, Bank: %d\n", wins, bigWins, losses, credits, bank)
	fmt.Println(red + strings.Join(msg, "\n") + reset)
}

// Function to display the boxes with random words
func displayBoxes() {
	middle1 := blue + "│ " + centerText(reel1, maxLength) + " │" + reset
	middle2 := blue + "│ " + centerText(reel2, maxLength) + " │" + reset
	middle3 := blue + "│ " + centerText(reel3, maxLength) + " │" + reset

	fmt.Println(top + " " + top + " " + top)
	fmt.Println(middle1 + " " + middle2 + " " + middle3)
	fmt.Println(bottom + " " + bottom + " " + bottom)
}

func spinReels() (reel1, reel2, reel3 string) {
	reel1 = reels[rand.Intn(len(reels))]
	reel2 = reels[rand.Intn(len(reels))]
	reel3 = reels[rand.Intn(len(reels))]
	return reel1, reel2, reel3
}

// Helper function to fetch the bank balance from the HTTP endpoint
func fetchBankBalance() (int, error) {
	resp, err := http.Get("http://localhost:9000/balance")
	if err != nil {
		return 0, err
	}
	defer resp.Body.Close()

	var result struct {
		Balance int `json:"balance"`
	}
	if err := json.NewDecoder(resp.Body).Decode(&result); err != nil {
		return 0, err
	}
	return result.Balance, nil
}

// Helper function to update the bank balance via HTTP POST request
func updateBankBalance(balance int) error {
	data := url.Values{}
	data.Set("balance", fmt.Sprintf("%d", balance))

	resp, err := http.PostForm("http://localhost:9000/balance", data)
	if err != nil {
		return err
	}
	defer resp.Body.Close()

	if resp.StatusCode != http.StatusOK {
		return fmt.Errorf("failed to update balance, status code: %d", resp.StatusCode)
	}
	return nil
}

// Helper function to center text within a given width
func centerText(text string, width int) string {
	padding := (width - len(text)) / 2
	return strings.Repeat(" ", padding) + text + strings.Repeat(" ", width-len(text)-padding)
}

// Helper function to listen for SSE updates
func listenForSSEUpdates() {
	client := sse.NewClient("http://localhost:9000/sse")

	client.Subscribe("messages", func(msg *sse.Event) {
		if string(msg.Event) == "balance" {
			var err error
			bank, err = fetchBankBalance()
			if err != nil {
				log.Fatalf("Failed to fetch bank balance: %v", err)
			}
			// Refresh the screen
			printScreen("Bank balance updated!")
		}
	})
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
