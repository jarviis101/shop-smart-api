package otp

import (
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/infrastructure/repository"
)

type Creator interface {
	Create(owner int64) (*entity.OTP, error)
}

type creator struct {
	repository repository.OTPRepository
	generator  Generator
}

func CreateCreator(r repository.OTPRepository, g Generator) Creator {
	return &creator{r, g}
}

func (c *creator) Create(owner int64) (*entity.OTP, error) {
	code := c.generator.Generate()

	return c.repository.Store(owner, code)
}
