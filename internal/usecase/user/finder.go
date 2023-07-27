package user

import (
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/infrastructure/repository"
)

type Finder interface {
	Find(id string) (*entity.User, error)
	FindByPhone(phone string) (*entity.User, error)
}

type finder struct {
	repository repository.UserRepository
}

func CreateFinder(repository repository.UserRepository) Finder {
	return &finder{repository}
}

func (f *finder) Find(id string) (*entity.User, error) {
	return f.repository.GetById(id)
}

func (f *finder) FindByPhone(phone string) (*entity.User, error) {
	return f.repository.GetByPhone(phone)
}
