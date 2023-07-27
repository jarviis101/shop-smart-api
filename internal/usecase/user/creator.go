package user

import (
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/infrastructure/repository"
)

type Creator interface {
	Create(phone string) (*entity.User, error)
	CreateWithData(
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

func (c *creator) Create(phone string) (*entity.User, error) {
	return c.repository.Store(phone, "", "", "", []string{"user"})
}

func (c *creator) CreateWithData(
	phone, firstName, lastName, middleName string,
	roles []string,
) (*entity.User, error) {
	return c.repository.Store(phone, firstName, lastName, middleName, roles)
}
