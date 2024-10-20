package main

import (
	"encoding/json"
	"fmt"
	"math/rand"
	"net/http"
	"sync"
)

var (
	balance float64
	mutex   sync.Mutex
)

func init() {
	balance = 100.0 // Starting balance
}

func main() {
	http.HandleFunc("/spin", spinHandler)
	http.HandleFunc("/balance", balanceHandler)
	http.Handle("/", http.FileServer(http.Dir("./static")))

	fmt.Println("Server is running on http://localhost:9000")
	err := http.ListenAndServe("127.0.0.1:9000", nil)
	if err != nil {
		fmt.Println("Failed to start server:", err)
	}
}

func spinHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	win := rand.Float64() < 0.5 // 50% chance to win
	if win {
		balance += 10.0
		fmt.Fprintln(w, "You win! New balance:", balance)
	} else {
		balance -= 5.0
		fmt.Fprintln(w, "You lose. New balance:", balance)
	}
}

func balanceHandler(w http.ResponseWriter, r *http.Request) {
	mutex.Lock()
	defer mutex.Unlock()

	balanceInfo := map[string]float64{
		"balance": balance,
	}
	json.NewEncoder(w).Encode(balanceInfo)
}
