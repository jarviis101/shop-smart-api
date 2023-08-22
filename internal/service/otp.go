package service

import (
	"shop-smart-api/internal/controller/http/types"
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/service/otp"
)

type otpService struct {
	sender    otp.Sender
	validator otp.Validator
}

func CreateOTPService(s otp.Sender, v otp.Validator) OTPService {
	return &otpService{s, v}
}

func (oc *otpService) Send(owner *entity.User, channel types.Channel) error {
	return oc.sender.SendOTP(owner, channel)
}

func (oc *otpService) Verify(owner *entity.User, code string) error {
	return oc.validator.Validate(owner.ID, code)
}
