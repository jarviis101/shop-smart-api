package repository

import (
	"context"
	"shop-smart-api/internal/entity"
)

type (
	UserRepository interface {
		Store(ctx context.Context, phone string) (*entity.User, error)
		GetByPhone(ctx context.Context, phone string) (*entity.User, error)
		GetById(ctx context.Context, id string) (*entity.User, error)
		UpdateUser(ctx context.Context, userId, firstName, lastName, middleName string) (*entity.User, error)
	}
	OTPRepository interface {
		Store(ctx context.Context, owner, code string) (*entity.OTP, error)
		GetByOwnerAndCode(ctx context.Context, owner, code string) (*entity.OTP, error)
		UseOTP(ctx context.Context, otp *entity.OTP) error
	}
)
