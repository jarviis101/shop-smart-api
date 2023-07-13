package user

import (
	"context"
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/infrastructure/repository"
	"shop-smart-api/internal/pkg/jwt"
)

type AuthService interface {
	PreAuthenticate(ctx context.Context, phone string) (string, error)
	FullAuthenticate(user *entity.User) (string, error)
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
	if err != nil {
		createdUser, _ := s.creatorService.Create(ctx, phone)
		return s.jwtManager.Generate(createdUser, false)
	}

	return s.jwtManager.Generate(user, false)
}

func (s *authService) FullAuthenticate(user *entity.User) (string, error) {
	return s.jwtManager.Generate(user, true)
}
