package vslot

import (
	"flag"
	"os"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	spin0, spin1 [3]int
	myVSlot      *MyVSlot
)

// might need to run with -parallel=1
func TestMain(m *testing.M) {
	flag.Parse()
	setup()

	code := m.Run()

	// teardown()

	os.Exit(code)
}

func setup() {
	myVSlot = NewMyVSlot(WinningSeed, LoseAmount)
	spin0 = WinningReels
	spin1 = [3]int{4, 5, 2}
}

func TestSpin(t *testing.T) {
	a := assert.New(t)

	myVSlot = NewMyVSlot(WinningSeed, LoseAmount)

	spinResult, isWinner, err := myVSlot.Spin()
	a.Equal(spin0, spinResult)
	a.True(isWinner)
	a.Nil(err)
}

func TestTwoSpins(t *testing.T) {
	a := assert.New(t)

	myVSlot = NewMyVSlot(WinningSeed, LoseAmount)

	spinResult, isWinner, err := myVSlot.Spin()
	a.Equal(spin0, spinResult)
	a.True(isWinner)
	a.Nil(err)

	spinResult, isWinner, err = myVSlot.Spin()
	a.Equal(spin1, spinResult)
	a.False(isWinner)
	a.Nil(err)
}

func TestUpdateBalance(t *testing.T) {
	a := assert.New(t)

	myVSlot = NewMyVSlot(WinningSeed, LoseAmount)
	a.Equal(LoseAmount, myVSlot.GetBalance())

	spinResult, isWinner, err := myVSlot.Spin()
	a.Equal(spin0, spinResult)
	a.Equal(LoseAmount+WinAmount, myVSlot.GetBalance())
	a.True(isWinner)
	a.Nil(err)

	myVSlot.UpdateBalance(10)
	a.Equal(10, myVSlot.GetBalance())
}

func TestReset(t *testing.T) {
	a := assert.New(t)

	myVSlot = NewMyVSlot(WinningSeed, LoseAmount)
	a.Equal(LoseAmount, myVSlot.GetBalance())

	spinResult, isWinner, err := myVSlot.Spin()
	a.Equal(spin0, spinResult)
	a.Equal(LoseAmount+WinAmount, myVSlot.GetBalance())
	a.True(isWinner)
	a.Nil(err)

	myVSlot.Reset()
	a.Equal(LoseAmount, myVSlot.GetBalance())
	a.Equal([3]int{}, myVSlot.GetReels())
}

func TestFindWinner(t *testing.T) {
	a := assert.New(t)

	seed := WinningSeed
	for {
		myVSlot = NewMyVSlot(seed, 100)
		_, isWinner, _ := myVSlot.Spin()
		if isWinner {
			break
		}
		seed += 1
	}
	a.True(true)
}

func TestWinner(t *testing.T) {
	a := assert.New(t)

	myVSlot := NewMyVSlot(WinningSeed, LoseAmount)
	reels, isWinner, err := myVSlot.Spin()

	a.True(isWinner)
	a.Equal(WinningReels, reels)
	a.Nil(err)
}

func TestLoser(t *testing.T) {
	a := assert.New(t)

	myVSlot := NewMyVSlot(0, LoseAmount)
	reels, isWinner, err := myVSlot.Spin()

	a.False(isWinner)
	a.Equal([3]int{5, 5, 4}, reels)
	a.Nil(err)
}

func TestNoFunds(t *testing.T) {
	a := assert.New(t)

	myVSlot := NewMyVSlot(WinningSeed, 0)
	reels, isWinner, err := myVSlot.Spin()

	a.False(isWinner)
	a.Equal([3]int{}, reels)
	a.NotNil(err)
}
