package otp

import (
	"context"
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/usecase"
)

type useCase struct {
	sender    Sender
	validator Validator
}

func CreateOTPUseCase(s Sender, v Validator) usecase.OTPUseCase {
	return &useCase{s, v}
}

func (oc *useCase) Send(ctx context.Context, owner *entity.User) error {
	return oc.sender.SendOTP(ctx, owner)
}

func (oc *useCase) Verify(ctx context.Context, owner *entity.User, code string) error {
	return oc.validator.Validate(ctx, owner.ID, code)
}
