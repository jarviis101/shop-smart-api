package repository

import (
	"shop-smart-api/internal/entity"
)

type (
	UserRepository interface {
		Get(id int64) (*entity.User, error)
		GetByPhone(phone string) (*entity.User, error)
		// GetByOrganization(id int64) ([]*entity.User, error)
		Store(phone, firstName, lastName, middleName string, roles []string) (*entity.User, error)
		UpdateUser(id int64, firstName, lastName, middleName string) (*entity.User, error)
	}
	OTPRepository interface {
		GetByOwnerAndCode(owner int64, code string) (*entity.OTP, error)
		Store(owner int64, code string) (*entity.OTP, error)
		Use(otp *entity.OTP) error
	}
	OrganizationRepository interface {
		Get(id int64) (*entity.Organization, error)
	}
)
