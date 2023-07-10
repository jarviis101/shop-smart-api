package mapper

import (
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/infrastructure/repository/mongo"
)

type UserMapper interface {
	SchemaToEntity(user *mongo.User) *entity.User
}

type userMapper struct {
	BaseMapper
}

func CreateUserMapper(bm BaseMapper) UserMapper {
	return &userMapper{bm}
}

func (u *userMapper) SchemaToEntity(user *mongo.User) *entity.User {
	return &entity.User{
		ID:         user.ID.Hex(),
		FirstName:  user.FirstName,
		LastName:   user.LastName,
		MiddleName: user.MiddleName,
		CreatedAt:  user.CreatedAt,
		UpdatedAt:  user.UpdatedAt,
		Phone:      user.Phone,
	}
}
