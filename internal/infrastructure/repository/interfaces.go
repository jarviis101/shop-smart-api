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
		GetById(id string) (*entity.User, error)
		UpdateUser(userId, firstName, lastName, middleName string) (*entity.User, error)
	}
	OTPRepository interface {
		Store(owner, code string) (*entity.OTP, error)
		GetByOwnerAndCode(owner, code string) (*entity.OTP, error)
		UseOTP(otp *entity.OTP) error
	}
)
