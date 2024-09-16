package vslot

import "math/rand"

func Spin(seed int64) [3]int {
	rand.Seed(seed)
	return [3]int{rand.Intn(10) + 1, rand.Intn(10) + 1, rand.Intn(10) + 1}
}
