package otp

import (
	"context"
	"shop-smart-api/internal/entity"
)

type Sender interface {
	SendOTP(ctx context.Context, owner *entity.User) error
}

type sender struct {
	creator Creator
}

func CreateSender(c Creator) Sender {
	return &sender{c}
}

func (s *sender) SendOTP(ctx context.Context, owner *entity.User) error {
	_, err := s.creator.Create(ctx, owner.ID)
	if err != nil {
		return err
	}

	return nil
}
