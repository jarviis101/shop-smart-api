package seeder

import (
	"fmt"
	"math/rand"
)

func generateRandomCode() string {
	return fmt.Sprintf("%08d", rand.Intn(100000000))
}

func generateRandomFloatValue(min, max float64) float64 {
	return min + rand.Float64()*(max-min)
}
