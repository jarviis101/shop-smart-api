package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID         primitive.ObjectID `bson:"_id,omitempty"`
	CreatedAt  time.Time          `bson:"created_at,omitempty"`
	UpdatedAt  time.Time          `bson:"updated_at,omitempty"`
	FirstName  string             `bson:"first_name,omitempty"`
	LastName   string             `bson:"last_name,omitempty"`
	MiddleName string             `bson:"middle_name,omitempty"`
	Phone      string             `bson:"phone,omitempty"`
}

type OTP struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	OwnerID   primitive.ObjectID `bson:"owner_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	ExpiredAt time.Time          `bson:"expired_at,omitempty"`
	Code      string             `bson:"code,omitempty"`
}
