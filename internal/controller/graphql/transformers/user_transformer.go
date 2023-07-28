package transformers

import (
	"shop-smart-api/internal/controller/graphql/graph/model"
	"shop-smart-api/internal/entity"
	"strconv"
)

type UserTransformer interface {
	TransformToModel(u *entity.User) *model.User
	TransformManyToModel(u []*entity.User) []*model.User
}

type userTransformer struct {
	BaseTransformer
}

func CreateUserTransformer(bt BaseTransformer) UserTransformer {
	return &userTransformer{bt}
}

func (t *userTransformer) TransformManyToModel(u []*entity.User) []*model.User {
	var users []*model.User
	for _, user := range u {
		m := t.TransformToModel(user)

		users = append(users, m)
	}

	return users
}

func (t *userTransformer) TransformToModel(u *entity.User) *model.User {
	return &model.User{
		ID:         strconv.Itoa(int(u.ID)),
		FirstName:  u.FirstName,
		LastName:   u.LastName,
		MiddleName: u.MiddleName,
		Phone:      u.Phone,
	}
}
