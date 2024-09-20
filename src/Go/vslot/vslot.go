package vslot

import "math/rand"

// VSlot interface defines the required methods for implementing a virtual slot machine.
type VSlot interface {
	Spin(seed int64) [3]int   // Spin accepts a seed and returns a slice of ints
	Reset()                   // Reset resets the spinner's state
	UpdateBalance(amount int) // UpdateBalance updates the balance by the given amount
	GetBalance() int          // GetBalance returns the current balance
}

// MyVSlot struct implements the Spinner interface.
type MyVSlot struct {
	balance int
}

// Ensure MySpinner implements the Spinner interface.
var _ VSlot = (*MyVSlot)(nil) // This line enforces the implementation at compile time.

// Spin generates a slice of random integers based on the seed.
func (s *MyVSlot) Spin(seed int64) [3]int {
	// Example implementation: generate a slice of integers based on the seed
	rand.Seed(seed)
	return [3]int{rand.Intn(10) + 1, rand.Intn(10) + 1, rand.Intn(10) + 1}
}

// Reset resets the spinner's state.
func (s *MyVSlot) Reset() {
	s.balance = 0 // Reset balance to 0
}

// UpdateBalance updates the balance by the given amount.
func (s *MyVSlot) UpdateBalance(amount int) {
	s.balance += amount
}

// GetBalance returns the current balance.
func (s *MyVSlot) GetBalance() int {
	return s.balance
}
