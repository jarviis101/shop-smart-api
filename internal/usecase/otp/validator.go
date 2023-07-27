package otp

import (
	"context"
	"errors"
	"shop-smart-api/internal/infrastructure/repository"
	"time"
)

const (
	devOTP = "1111"
)

type Validator interface {
	Validate(ctx context.Context, owner, code string) error
}

type validator struct {
	repository repository.OTPRepository
	isDebug    bool
}

func CreateValidator(r repository.OTPRepository, d bool) Validator {
	return &validator{r, d}
}

func (v *validator) Validate(ctx context.Context, owner, code string) error {
	if v.isDebug && code == devOTP {
		return nil
	}

	otp, err := v.repository.GetByOwnerAndCode(owner, code)
	if err != nil {
		return errors.New("code not found")
	}

	if otp.IsUsed {
		return errors.New("code not found")
	}

	if time.Now().After(otp.ExpiredAt) {
		return errors.New("code is expired")
	}

	return v.repository.Use(otp)
}
