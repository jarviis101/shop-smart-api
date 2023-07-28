package service

import (
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/service/otp"
)

type otpService struct {
	sender    otp.Sender
	validator otp.Validator
}

func CreateOTPService(s otp.Sender, v otp.Validator) OTPUseCase {
	return &otpService{s, v}
}

func (oc *otpService) Send(owner *entity.User) error {
	return oc.sender.SendOTP(owner)
}

func (oc *otpService) Verify(owner *entity.User, code string) error {
	return oc.validator.Validate(owner.ID, code)
}
