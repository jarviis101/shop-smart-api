package usecase

import (
	"shop-smart-api/internal/entity"
)

type (
	UserUseCase interface {
		PreAuthenticate(phone string) (string, error)
		Authenticate(user *entity.User) (string, error)
		Get(id int64) (*entity.User, error)
		GetByPhone(phone string) (*entity.User, error)
		Update(user *entity.User, firstName, lastName, middleName string) (*entity.User, error)
		Create(
			phone, firstName, lastName, middleName string,
			roles []string,
		) (*entity.User, error)
	}
	OTPUseCase interface {
		Send(*entity.User) error
		Verify(owner *entity.User, code string) error
	}
	OrganizationUseCase interface {
	}
)
