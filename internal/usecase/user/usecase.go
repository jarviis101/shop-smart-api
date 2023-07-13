package user

import (
	"context"
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/usecase"
)

type useCase struct {
	auth      AuthService
	finder    Finder
	collector Collector
}

func CreateUserUseCase(a AuthService, f Finder, cs Collector) usecase.UserUseCase {
	return &useCase{a, f, cs}
}

func (uc *useCase) PreAuthenticate(ctx context.Context, phone string) (string, error) {
	return uc.auth.PreAuthenticate(ctx, phone)
}

func (uc *useCase) Authenticate(user *entity.User) (string, error) {
	return uc.auth.FullAuthenticate(user)
}

func (uc *useCase) Get(ctx context.Context, id string) (*entity.User, error) {
	return uc.finder.Find(ctx, id)
}

func (uc *useCase) GetByPhone(ctx context.Context, phone string) (*entity.User, error) {
	return uc.finder.FindByPhone(ctx, phone)
}
