package repository

import (
	"shop-smart-api/internal/entity"
	"time"
)

type (
	UserRepository interface {
		Get(id int64) (*entity.User, error)
		GetByPhone(phone string) (*entity.User, error)
		GetByEmail(email string) (*entity.User, error)
		GetByOrganization(id int64) ([]*entity.User, error)
		GetAll() ([]*entity.User, error)
		Store(phone, email string, roles []entity.Role) (*entity.User, error)
		UpdateUser(id int64, email string) (*entity.User, error)
		AddOrganization(id, organization int64, role *entity.Role) (*entity.User, error)
	}
	OTPRepository interface {
		GetByOwnerAndCode(owner int64, code string) (*entity.OTP, error)
		Store(owner int64, code string) (*entity.OTP, error)
		Use(otp *entity.OTP) error
	}
	OrganizationRepository interface {
		Get(id int64) (*entity.Organization, error)
		Store(name, kpp, orgn, inn string, owner int64) (*entity.Organization, error)
	}
	TransactionRepository interface {
		GetByOwner(id int64) ([]*entity.Transaction, error)
		Store(
			owner int64,
			trxNumber string,
			value float64,
			actionedAt *time.Time,
			status bool,
		) (*entity.Transaction, error)
	}
)
