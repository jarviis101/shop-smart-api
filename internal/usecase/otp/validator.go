package otp

import (
	"context"
	"errors"
	"shop-smart-api/internal/infrastructure/repository"
	"time"
)

type Validator interface {
	Validate(ctx context.Context, owner, code string) error
}

type validator struct {
	repository repository.OTPRepository
}

func CreateValidator(r repository.OTPRepository) Validator {
	return &validator{r}
}

func (v *validator) Validate(ctx context.Context, owner, code string) error {
	otp, err := v.repository.GetByOwnerAndCode(ctx, owner, code)
	if err != nil {
		return errors.New("code not found")
	}

	if otp.IsUsed {
		return errors.New("code not found")
	}

	if time.Now().After(otp.ExpiredAt) {
		return errors.New("code is expired")
	}

	return v.repository.UseOTP(ctx, otp)
}
