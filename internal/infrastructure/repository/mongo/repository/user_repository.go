package repository

import (
	"context"
	"go.mongodb.org/mongo-driver/bson"
	"go.mongodb.org/mongo-driver/bson/primitive"
	"go.mongodb.org/mongo-driver/mongo"
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/infrastructure/repository"
	schema "shop-smart-api/internal/infrastructure/repository/mongo"
	"shop-smart-api/internal/infrastructure/repository/mongo/mapper"
	"time"
)

type userRepository struct {
	BaseRepository
	collection *mongo.Collection
	mapper     mapper.UserMapper
}

func CreateUserRepository(br BaseRepository, c *mongo.Collection, m mapper.UserMapper) repository.UserRepository {
	return &userRepository{br, c, m}
}

func (r *userRepository) Store(ctx context.Context, phone string) (*entity.User, error) {
	var user *schema.User
	result, err := r.collection.InsertOne(ctx, &schema.User{
		ID:        primitive.NewObjectID(),
		Phone:     phone,
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})

	if err != nil {
		return nil, err
	}

	if err := r.collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&user); err != nil {
		return nil, err
	}

	return r.mapper.SchemaToEntity(user), nil
}

func (r *userRepository) GetByPhone(ctx context.Context, phone string) (*entity.User, error) {
	var user *schema.User
	err := r.collection.FindOne(ctx, bson.M{"phone": phone}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return r.mapper.SchemaToEntity(user), nil
}

func (r *userRepository) GetById(ctx context.Context, id string) (*entity.User, error) {
	var user *schema.User
	objectId, err := primitive.ObjectIDFromHex(id)
	if err != nil {
		return nil, err
	}

	err = r.collection.FindOne(ctx, bson.M{"_id": objectId}).Decode(&user)
	if err != nil {
		return nil, err
	}

	return r.mapper.SchemaToEntity(user), nil
}

//func (r *userRepository) GetByIds(ctx context.Context, ids []string) ([]*entity.User, error) {
//	var users []*schema.User
//	var usersEntity []*entity.User
//	objectIds := r.fromStringToObjectId(ids)
//	cur, err := r.collection.Find(ctx, bson.M{"_id": bson.M{
//		"$in": objectIds,
//	}})
//	if err != nil {
//		return nil, err
//	}
//
//	for cur.Next(ctx) {
//		var user *schema.User
//		if err := cur.Decode(&user); err != nil {
//			return nil, err
//		}
//
//		users = append(users, user)
//	}
//
//	for _, u := range users {
//		user := r.mapper.SchemaToEntity(u)
//		usersEntity = append(usersEntity, user)
//	}
//
//	return usersEntity, nil
//}
