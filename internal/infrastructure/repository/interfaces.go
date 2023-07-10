package repository

import (
	"context"
	"shop-smart-api/internal/entity"
)

type (
	UserRepository interface {
		Store(ctx context.Context, u *entity.User) (*entity.User, error)
		GetByPhone(ctx context.Context, phone string) (*entity.User, error)
		GetById(ctx context.Context, id string) (*entity.User, error)
	}
)
