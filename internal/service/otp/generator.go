package otp

import (
	"math/rand"
	"strconv"
)

type Generator interface {
	Generate() string
}

type generator struct {
}

func CreateGenerator() Generator {
	return &generator{}
}

func (g *generator) Generate() string {
	var code string

	for i := 0; i < 4; i++ {
		code += strconv.Itoa(rand.Intn(10))
	}

	return code
}
