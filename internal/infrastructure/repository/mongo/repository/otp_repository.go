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

type otpRepository struct {
	BaseRepository
	collection *mongo.Collection
	mapper     mapper.OTPMapper
}

func CreateOTPRepository(br BaseRepository, c *mongo.Collection, m mapper.OTPMapper) repository.OTPRepository {
	return &otpRepository{br, c, m}
}

func (r *otpRepository) Store(ctx context.Context, code, owner string) (*entity.OTP, error) {
	ownerId, err := primitive.ObjectIDFromHex(owner)
	if err != nil {
		return nil, err
	}

	result, err := r.collection.InsertOne(ctx, &schema.OTP{
		ID:        primitive.NewObjectID(),
		Code:      code,
		OwnerID:   ownerId,
		ExpiredAt: time.Now().Add(time.Minute * 5),
		CreatedAt: time.Now(),
		UpdatedAt: time.Now(),
	})
	if err != nil {
		return nil, err
	}

	var otp *schema.OTP
	if err := r.collection.FindOne(ctx, bson.M{"_id": result.InsertedID}).Decode(&otp); err != nil {
		return nil, err
	}

	return r.mapper.SchemaToEntity(otp), nil
}
