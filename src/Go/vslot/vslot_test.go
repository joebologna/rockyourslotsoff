package vslot

import (
	"math/rand"
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestSpin(t *testing.T) {
	a := assert.New(t)
	rand.Seed(0)

	a.Equal([3]int{rand.Intn(10) + 1, rand.Intn(10) + 1, rand.Intn(10) + 1}, Spin(0))
}
