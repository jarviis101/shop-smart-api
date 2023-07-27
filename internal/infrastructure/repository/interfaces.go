package repository

import (
	"context"
	"shop-smart-api/internal/entity"
)

type (
	UserRepository interface {
		Store(ctx context.Context, phone string) (*entity.User, error)
		StoreWithData(
			ctx context.Context,
			phone, firstName, lastName, middleName string,
			roles []string,
		) (*entity.User, error)
		GetByPhone(ctx context.Context, phone string) (*entity.User, error)
		GetById(ctx context.Context, id string) (*entity.User, error)
		UpdateUser(ctx context.Context, userId, firstName, lastName, middleName string) (*entity.User, error)
	}
	OTPRepository interface {
		Store(owner, code string) (*entity.OTP, error)
		GetByOwnerAndCode(owner, code string) (*entity.OTP, error)
		UseOTP(otp *entity.OTP) error
	}
)
