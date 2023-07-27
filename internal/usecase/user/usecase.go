package user

import (
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/usecase"
)

type useCase struct {
	auth      AuthService
	finder    Finder
	collector Collector
	modifier  Modifier
	creator   Creator
}

func CreateUserUseCase(a AuthService, f Finder, cs Collector, m Modifier, c Creator) usecase.UserUseCase {
	return &useCase{a, f, cs, m, c}
}

func (uc *useCase) PreAuthenticate(phone string) (string, error) {
	return uc.auth.PreAuthenticate(phone)
}

func (uc *useCase) Authenticate(user *entity.User) (string, error) {
	return uc.auth.FullAuthenticate(user)
}

func (uc *useCase) Get(id int64) (*entity.User, error) {
	return uc.finder.Find(id)
}

func (uc *useCase) GetByPhone(phone string) (*entity.User, error) {
	return uc.finder.FindByPhone(phone)
}

func (uc *useCase) Update(
	user *entity.User,
	firstName, lastName, middleName string,
) (*entity.User, error) {
	return uc.modifier.UpdateUser(user, firstName, lastName, middleName)
}

func (uc *useCase) Create(
	phone, firstName, lastName, middleName string,
	roles []string,
) (*entity.User, error) {
	return uc.creator.CreateWithData(phone, firstName, lastName, middleName, roles)
}
