package service

import (
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/service/user"
)

type userService struct {
	auth      user.AuthService
	finder    user.Finder
	collector user.Collector
	modifier  user.Modifier
	creator   user.Creator
}

func CreateUserService(
	a user.AuthService,
	f user.Finder,
	cs user.Collector,
	m user.Modifier,
	c user.Creator,
) UserService {
	return &userService{a, f, cs, m, c}
}

func (uc *userService) Get(id int64) (*entity.User, error) {
	return uc.finder.Find(id)
}

func (uc *userService) GetByPhone(phone string) (*entity.User, error) {
	return uc.finder.FindByPhone(phone)
}

func (uc *userService) GetByOrganization(id int64) ([]*entity.User, error) {
	return uc.finder.FindByOrganization(id)
}

func (uc *userService) PreAuthenticate(phone string) (string, error) {
	return uc.auth.PreAuthenticate(phone)
}

func (uc *userService) Authenticate(user *entity.User) (string, error) {
	return uc.auth.FullAuthenticate(user)
}

func (uc *userService) Update(
	user *entity.User,
	firstName, lastName, middleName string,
) (*entity.User, error) {
	return uc.modifier.UpdateUser(user, firstName, lastName, middleName)
}

func (uc *userService) Create(
	phone, firstName, lastName, middleName string,
	roles []string,
) (*entity.User, error) {
	return uc.creator.CreateWithData(phone, firstName, lastName, middleName, roles)
}
