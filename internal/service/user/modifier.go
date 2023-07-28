package user

import (
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/infrastructure/repository"
)

type Modifier interface {
	UpdateUser(user *entity.User, firstName, lastName, middleName string) (*entity.User, error)
}

type modifier struct {
	repository repository.UserRepository
}

func CreateModifier(r repository.UserRepository) Modifier {
	return &modifier{r}
}

func (m *modifier) UpdateUser(
	user *entity.User, firstName, lastName, middleName string,
) (*entity.User, error) {
	return m.repository.UpdateUser(user.ID, firstName, lastName, middleName)
}
