package utils

import (
	"math/rand"
	"time"
)

func GetRandomUpTo(upperLimit int) int {
	rand.Seed(time.Now().UnixNano())
	num := rand.Intn(upperLimit)
	return num
}
