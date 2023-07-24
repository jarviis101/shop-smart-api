package mongo

import (
	"go.mongodb.org/mongo-driver/bson/primitive"
	"time"
)

type User struct {
	ID             primitive.ObjectID `bson:"_id,omitempty"`
	OrganizationID primitive.ObjectID `bson:"organization_id,omitempty"`
	CreatedAt      time.Time          `bson:"created_at,omitempty"`
	UpdatedAt      time.Time          `bson:"updated_at,omitempty"`
	FirstName      string             `bson:"first_name,omitempty"`
	LastName       string             `bson:"last_name,omitempty"`
	MiddleName     string             `bson:"middle_name,omitempty"`
	Phone          string             `bson:"phone,omitempty"`
	Roles          []string           `bson:"roles,omitempty"`
}

type OTP struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	OwnerID   primitive.ObjectID `bson:"owner_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	ExpiredAt time.Time          `bson:"expired_at,omitempty"`
	Code      string             `bson:"code,omitempty"`
	IsUsed    bool               `bson:"is_used,omitempty"`
}

type Organization struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	OwnerID   primitive.ObjectID `bson:"owner_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	Name      string             `bson:"name,omitempty"`
	ORGN      string             `bson:"orgn,omitempty"`
	KPP       string             `bson:"kpp,omitempty"`
	INN       string             `bson:"inn,omitempty"`
}

type Transaction struct {
	ID        primitive.ObjectID `bson:"_id,omitempty"`
	OwnerID   primitive.ObjectID `bson:"owner_id,omitempty"`
	CreatedAt time.Time          `bson:"created_at,omitempty"`
	UpdatedAt time.Time          `bson:"updated_at,omitempty"`
	ActionAt  time.Time          `bson:"actioned_at,omitempty"`
	TrxNumber string             `bson:"trx_number,omitempty"`
	Value     string             `bson:"value,omitempty"`
	Status    bool               `bson:"status,omitempty"`
}
