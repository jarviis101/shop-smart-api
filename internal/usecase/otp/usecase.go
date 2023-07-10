package otp

import (
	"context"
	"shop-smart-api/internal/usecase"
)

type useCase struct {
}

func CreateOTPUseCase() usecase.OTPUseCase {
	return &useCase{}
}

func (oc *useCase) Send(ctx context.Context, phone string) error {
	return nil
}

func (oc *useCase) Verify(ctx context.Context, code string) bool {
	return false
}
