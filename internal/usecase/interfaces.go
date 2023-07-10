package usecase

import (
	"context"
	"shop-smart-api/internal/entity"
)

type (
	UserUseCase interface {
		PreAuthenticate(ctx context.Context, phone string) (string, error)
		Authenticate(ctx context.Context, id string) (string, error)
		Get(ctx context.Context, id string) (*entity.User, error)
	}
	OTPUseCase interface {
		Send(ctx context.Context, phone string) error
		Verify(ctx context.Context, code string) bool
	}
)
