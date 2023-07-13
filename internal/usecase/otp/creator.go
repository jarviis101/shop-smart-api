package otp

import (
	"context"
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/infrastructure/repository"
)

type Creator interface {
	Create(ctx context.Context, owner string) (*entity.OTP, error)
}

type creator struct {
	repository repository.OTPRepository
	generator  Generator
}

func CreateCreator(r repository.OTPRepository, g Generator) Creator {
	return &creator{r, g}
}

func (c *creator) Create(ctx context.Context, owner string) (*entity.OTP, error) {
	code := c.generator.Generate()
	return c.repository.Store(ctx, code, owner)
}
