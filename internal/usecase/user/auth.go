package user

import (
	"context"
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/infrastructure/repository"
	"shop-smart-api/internal/pkg/jwt"
)

type AuthService interface {
	PreAuthenticate(ctx context.Context, phone string) (string, error)
	Auth(ctx context.Context, id string) (string, error)
}

type authService struct {
	repository     repository.UserRepository
	jwtManager     jwt.Manager
	creatorService Creator
}

func CreateAuthService(r repository.UserRepository, j jwt.Manager, c Creator) AuthService {
	return &authService{r, j, c}
}

func (s *authService) PreAuthenticate(ctx context.Context, phone string) (string, error) {
	user, err := s.repository.GetByPhone(ctx, phone)
	if err != nil && user == nil {
		user, _ := s.creatorService.CreateUser(ctx, &entity.User{Phone: phone})
		return s.jwtManager.Generate(user, false)
	}

	return s.jwtManager.Generate(user, false)
}

func (s *authService) Auth(ctx context.Context, id string) (string, error) {
	user, err := s.repository.GetById(ctx, id)
	if err != nil {
		return "", err
	}

	return s.jwtManager.Generate(user, true)
}
