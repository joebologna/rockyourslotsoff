package vslot

import (
	"fmt"
	"math/rand"
)

// VSlot interface defines the required methods for implementing a virtual slot machine.
type VSlot interface {
	Spin() (reels [3]int, is_winner bool, err error)
	GetBalance() (balance int)
	GetReels() (reels [3]int)
	Reset()
	UpdateBalance(amount int) (balance int)
}

// MyVSlot struct implements the VSlot interface.
type MyVSlot struct {
	seed                    int64
	balance, initialBalance int
	reels                   [3]int
}

var WinningReels = [3]int{7, 7, 7}

const (
	WinAmount   = 100
	LoseAmount  = 10
	WinningSeed = int64(589)
)

// Ensure MySpinner implements the VSlot interface.
var _ VSlot = (*MyVSlot)(nil) // This line enforces the implementation at compile time.

func NewMyVSlot(initialSeed int64, initialBalance int) *MyVSlot {
	rand.Seed(initialSeed)
	return &MyVSlot{seed: initialSeed, balance: initialBalance, initialBalance: initialBalance}
}

// Spin sets the reels to random values and returns them.
func (vs *MyVSlot) Spin() (reels [3]int, is_winner bool, err error) {
	if vs.balance < LoseAmount {
		return [3]int{}, false, fmt.Errorf("insufficient balance")
	}
	vs.reels[0] = rand.Intn(10) + 1
	vs.reels[1] = rand.Intn(10) + 1
	vs.reels[2] = rand.Intn(10) + 1
	if vs.reels == WinningReels {
		vs.balance += WinAmount
		return vs.reels, true, nil
	}
	vs.balance -= LoseAmount
	return vs.reels, false, nil
}

// GetReels returns the current reels.
func (vs *MyVSlot) GetReels() [3]int {
	return vs.reels
}

// Reset resets the spinner's state.
func (vs *MyVSlot) Reset() {
	vs.balance = vs.initialBalance
	vs.reels = [3]int{}
	rand.Seed(vs.seed)
}

// UpdateBalance updates the balance by the given amount and returns the new balance.
func (vs *MyVSlot) UpdateBalance(amount int) int {
	vs.balance = amount
	return vs.balance
}

// GetBalance returns the current balance.
func (vs *MyVSlot) GetBalance() int {
	return vs.balance
}
