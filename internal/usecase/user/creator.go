package user

import (
	"context"
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/infrastructure/repository"
)

type Creator interface {
	Create(ctx context.Context, phone string) (*entity.User, error)
	CreateWithData(
		ctx context.Context,
		phone, firstName, lastName, middleName string,
		roles []string,
	) (*entity.User, error)
}

type creator struct {
	repository repository.UserRepository
}

func CreateCreator(r repository.UserRepository) Creator {
	return &creator{r}
}

func (c *creator) Create(ctx context.Context, phone string) (*entity.User, error) {
	return c.repository.Store(ctx, phone)
}

func (c *creator) CreateWithData(
	ctx context.Context,
	phone, firstName, lastName, middleName string,
	roles []string,
) (*entity.User, error) {
	return c.repository.StoreWithData(ctx, phone, firstName, lastName, middleName, roles)
}
