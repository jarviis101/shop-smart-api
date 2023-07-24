package user

import (
	"context"
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/infrastructure/repository"
)

type Modifier interface {
	UpdateUser(ctx context.Context, user *entity.User, firstName, lastName, middleName string) (*entity.User, error)
}

type modifier struct {
	repository repository.UserRepository
}

func CreateModifier(r repository.UserRepository) Modifier {
	return &modifier{r}
}

func (m *modifier) UpdateUser(
	ctx context.Context,
	user *entity.User, firstName, lastName, middleName string,
) (*entity.User, error) {
	return m.repository.UpdateUser(ctx, user.ID, firstName, lastName, middleName)
}
