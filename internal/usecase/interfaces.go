package usecase

import (
	"context"
	"shop-smart-api/internal/entity"
)

type (
	UserUseCase interface {
		PreAuthenticate(ctx context.Context, phone string) (string, error)
		Authenticate(user *entity.User) (string, error)
		Get(ctx context.Context, id string) (*entity.User, error)
		GetByPhone(ctx context.Context, phone string) (*entity.User, error)
	}
	OTPUseCase interface {
		Send(ctx context.Context, owner *entity.User) error
		Verify(ctx context.Context, code string) bool
	}
)
