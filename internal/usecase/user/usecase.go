package user

import (
	"context"
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/usecase"
)

type useCase struct {
	authService      AuthService
	finderService    Finder
	collectorService Collector
}

func CreateUserUseCase(a AuthService, f Finder, cs Collector) usecase.UserUseCase {
	return &useCase{a, f, cs}
}

func (uc *useCase) PreAuthenticate(ctx context.Context, phone string) (string, error) {
	return uc.authService.PreAuthenticate(ctx, phone)
}

func (uc *useCase) Authenticate(user *entity.User) (string, error) {
	return uc.authService.FullAuthenticate(user)
}

func (uc *useCase) Get(ctx context.Context, id string) (*entity.User, error) {
	return uc.finderService.Find(ctx, id)
}

func (uc *useCase) GetByPhone(ctx context.Context, phone string) (*entity.User, error) {
	return uc.finderService.FindByPhone(ctx, phone)
}
