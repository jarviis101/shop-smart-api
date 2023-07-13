package otp

import (
	"context"
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/usecase"
)

type useCase struct {
	creator Creator
	sender  Sender
}

func CreateOTPUseCase(c Creator, s Sender) usecase.OTPUseCase {
	return &useCase{c, s}
}

func (oc *useCase) Send(ctx context.Context, owner *entity.User) error {
	_, err := oc.creator.Create(ctx, owner.ID)
	if err != nil {
		return err
	}

	return oc.sender.SendOTP()
}

func (oc *useCase) Verify(ctx context.Context, code string) bool {
	return false
}
