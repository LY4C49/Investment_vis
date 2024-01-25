package tools

import (
	"math/rand"
	"time"
)

func GetRandomInt(base int) int {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return r.Intn(base) - (base / 2)
}

func GetRandomFloat(base float64) float64 {
	s := rand.NewSource(time.Now().UnixNano())
	r := rand.New(s)

	return (r.Float64() - 0.5) * base
}
