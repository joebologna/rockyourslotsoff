package vslot

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

var (
	spin0, spin1               [3]int
	spin0_result, spin1_result bool
)

func TestSpin(t *testing.T) {
	a := assert.New(t)
	rand.Seed(0)

	v := MyVSlot{} // REVIEW Consider setting the seed using a New function to prevent using rand.Intn() before the seed is set. Maybe add an error return wherever rand.Intn() is called.

	spin_result := v.Spin(0)
	a.Equal(spin0, spin_result)
}

func TestBalance(t *testing.T) {
	a := assert.New(t)
	rand.Seed(0)

	v := MyVSlot{}
	a.Equal(v.GetBalance(), 0)

	spin_result := v.Spin(0)
	is_win := spin_result == spin0
	a.True(is_win)

	if is_win {
		v.UpdateBalance(10)
		a.Equal(10, v.GetBalance())
	} else {
		a.FailNow("should have won")
	}
}
