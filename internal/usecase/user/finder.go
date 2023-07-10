package user

import (
	"context"
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/infrastructure/repository"
)

type Finder interface {
	Find(ctx context.Context, id string) (*entity.User, error)
	FindByPhone(ctx context.Context, phone string) (*entity.User, error)
}

type finder struct {
	repository repository.UserRepository
}

func CreateFinder(repository repository.UserRepository) Finder {
	return &finder{repository}
}

func (f *finder) Find(ctx context.Context, id string) (*entity.User, error) {
	return f.repository.GetById(ctx, id)
}

func (f *finder) FindByPhone(ctx context.Context, phone string) (*entity.User, error) {
	return f.repository.GetByPhone(ctx, phone)
}
