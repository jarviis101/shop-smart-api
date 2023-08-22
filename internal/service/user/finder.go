package user

import (
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/infrastructure/repository"
)

type Finder interface {
	Find(id int64) (*entity.User, error)
	FindByPhone(phone string) (*entity.User, error)
	FindByEmail(email string) (*entity.User, error)
	FindByOrganization(id int64) ([]*entity.User, error)
}

type finder struct {
	repository repository.UserRepository
}

func CreateFinder(r repository.UserRepository) Finder {
	return &finder{r}
}

func (f *finder) Find(id int64) (*entity.User, error) {
	return f.repository.Get(id)
}

func (f *finder) FindByPhone(phone string) (*entity.User, error) {
	return f.repository.GetByPhone(phone)
}

func (f *finder) FindByEmail(email string) (*entity.User, error) {
	return f.repository.GetByEmail(email)
}

func (f *finder) FindByOrganization(id int64) ([]*entity.User, error) {
	return f.repository.GetByOrganization(id)
}
