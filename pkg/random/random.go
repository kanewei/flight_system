package random

import "math/rand"

func RandRange(min, max int) int {
	return min + rand.Intn(max-min)
}
