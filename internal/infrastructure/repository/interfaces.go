package repository

import (
	"shop-smart-api/internal/entity"
)

type (
	UserRepository interface {
		Store(
			phone, firstName, lastName, middleName string,
			roles []string,
		) (*entity.User, error)
		GetByPhone(phone string) (*entity.User, error)
		GetById(id int64) (*entity.User, error)
		UpdateUser(id int64, firstName, lastName, middleName string) (*entity.User, error)
	}
	OTPRepository interface {
		Store(owner int64, code string) (*entity.OTP, error)
		GetByOwnerAndCode(owner int64, code string) (*entity.OTP, error)
		Use(otp *entity.OTP) error
	}
)
