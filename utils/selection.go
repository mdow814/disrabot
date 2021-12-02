package utils

import "math/rand"

func GetRandomUpTo(upperLimit int) int {
	num := rand.Intn(upperLimit)
	return num
}
