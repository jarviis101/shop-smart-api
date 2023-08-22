package otp

import (
	"shop-smart-api/internal/controller/http/types"
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/pkg/sms"
)

type Sender interface {
	SendOTP(owner *entity.User, channel types.Channel) error
}

type sender struct {
	creator Creator
	client  sms.Client
}

func CreateSender(c Creator, cl sms.Client) Sender {
	return &sender{c, cl}
}

func (s *sender) SendOTP(owner *entity.User, channel types.Channel) error {
	otp, err := s.creator.Create(owner.ID)
	if err != nil {
		return err
	}

	if channel == types.Phone {
		go s.client.Send(owner.Phone, otp.Code)
	}

	if channel == types.Email {
		// TODO: Complete
		return nil
	}

	return nil
}
