package usecase

import (
	"context"
	"shop-smart-api/internal/entity"
)

type (
	UserUseCase interface {
		PreAuthenticate(phone string) (string, error)
		Authenticate(user *entity.User) (string, error)
		Get(id string) (*entity.User, error)
		GetByPhone(phone string) (*entity.User, error)
		Update(user *entity.User, firstName, lastName, middleName string) (*entity.User, error)
		Create(
			phone, firstName, lastName, middleName string,
			roles []string,
		) (*entity.User, error)
	}
	OTPUseCase interface {
		Send(ctx context.Context, owner *entity.User) error
		Verify(ctx context.Context, owner *entity.User, code string) error
	}
	OrganizationUseCase interface {
	}
)
