package otp

import (
	"shop-smart-api/internal/controller/http/types"
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/pkg/email"
	"shop-smart-api/internal/pkg/sms"
)

type Sender interface {
	SendOTP(owner *entity.User, channel *types.Channel) error
}

type sender struct {
	creator Creator
	client  sms.Client
	mailer  email.Mailer
}

func CreateSender(c Creator, cl sms.Client, m email.Mailer) Sender {
	return &sender{c, cl, m}
}

func (s *sender) SendOTP(owner *entity.User, channel *types.Channel) error {
	otp, err := s.creator.Create(owner.ID)
	if err != nil {
		return err
	}

	if channel.IsPhone() {
		go s.client.Send(owner.Phone, otp.Code)
	}

	if channel.IsEmail() {
		go s.mailer.Send(owner.Email, otp.Code)
	}

	return nil
}
