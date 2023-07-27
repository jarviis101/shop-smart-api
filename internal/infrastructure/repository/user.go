package repository

import (
	"context"
	"database/sql"
	"shop-smart-api/internal/entity"
)

type userRepository struct {
	database *sql.DB
}

func CreateUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (u userRepository) Store(ctx context.Context, phone string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) StoreWithData(ctx context.Context, phone, firstName, lastName, middleName string, roles []string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetByPhone(ctx context.Context, phone string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) GetById(ctx context.Context, id string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}

func (u userRepository) UpdateUser(ctx context.Context, userId, firstName, lastName, middleName string) (*entity.User, error) {
	//TODO implement me
	panic("implement me")
}
