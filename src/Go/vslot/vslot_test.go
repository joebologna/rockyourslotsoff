package vslot

import (
	"flag"
	"math/rand"
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
	myVSlot = NewMyVSlot(0)
	spin0 = [3]int{rand.Intn(10) + 1, rand.Intn(10) + 1, rand.Intn(10) + 1}
	spin1 = [3]int{rand.Intn(10) + 1, rand.Intn(10) + 1, rand.Intn(10) + 1}
}

func TestSpin(t *testing.T) {
	a := assert.New(t)

	myVSlot = NewMyVSlot(0)

	spin_result := myVSlot.Spin()
	a.Equal(spin0, spin_result)
}

func TestTwoSpins(t *testing.T) {
	a := assert.New(t)

	myVSlot = NewMyVSlot(0)

	spin_result := myVSlot.Spin()
	a.Equal(spin0, spin_result)

	spin_result = myVSlot.Spin()
	a.Equal(spin1, spin_result)
}

func TestBalance(t *testing.T) {
	a := assert.New(t)

	myVSlot = NewMyVSlot(0)
	a.Equal(myVSlot.GetBalance(), 0)

	spin_result := myVSlot.Spin()
	a.Equal(spin0, spin_result)

	myVSlot.UpdateBalance(10)
	a.Equal(10, myVSlot.GetBalance())
}

func TestReset(t *testing.T) {
	a := assert.New(t)

	myVSlot = NewMyVSlot(0)
	a.Equal(myVSlot.GetBalance(), 0)

	spin_result := myVSlot.Spin()
	a.Equal(spin0, spin_result)

	myVSlot.UpdateBalance(10)
	a.Equal(10, myVSlot.GetBalance())

	myVSlot.Reset()
	a.Equal(0, myVSlot.GetBalance())
}
