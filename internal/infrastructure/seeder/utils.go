package seeder

import (
	"fmt"
	"math/rand"
)

func generateRandomCode() string {
	return fmt.Sprintf("%08d", rand.Intn(100000000))
}
