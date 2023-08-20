package service

import (
	"shop-smart-api/internal/entity"
)

type (
	UserService interface {
		Get(id int64) (*entity.User, error)
		GetByPhone(phone string) (*entity.User, error)
		GetByOrganization(id int64) ([]*entity.User, error)
		PreAuthenticate(phone string) (string, error)
		Authenticate(user *entity.User) (string, error)
		Update(user *entity.User, email string) (*entity.User, error)
	}
	OTPService interface {
		Send(*entity.User) error
		Verify(owner *entity.User, code string) error
	}
	OrganizationService interface {
		Get(id int64) (*entity.Organization, error)
	}
	TransactionService interface {
		GetTransactions(owner *entity.User) ([]*entity.Transaction, error)
	}
)
