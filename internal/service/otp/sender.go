package otp

import (
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/pkg/sms"
)

type Sender interface {
	SendOTP(owner *entity.User) error
}

type sender struct {
	creator Creator
	client  sms.Client
}

func CreateSender(c Creator, cl sms.Client) Sender {
	return &sender{c, cl}
}

func (s *sender) SendOTP(owner *entity.User) error {
	otp, err := s.creator.Create(owner.ID)
	if err != nil {
		return err
	}

	go s.client.Send(owner.Phone, otp.Code)

	return nil
}
