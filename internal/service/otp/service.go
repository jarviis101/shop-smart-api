package otp

import (
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/service"
)

type useCase struct {
	sender    Sender
	validator Validator
}

func CreateOTPUseCase(s Sender, v Validator) service.OTPUseCase {
	return &useCase{s, v}
}

func (oc *useCase) Send(owner *entity.User) error {
	return oc.sender.SendOTP(owner)
}

func (oc *useCase) Verify(owner *entity.User, code string) error {
	return oc.validator.Validate(owner.ID, code)
}
