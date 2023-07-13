package mapper

import (
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/infrastructure/repository/mongo"
)

type OTPMapper interface {
	SchemaToEntity(o *mongo.OTP) *entity.OTP
}

type otpMapper struct {
	BaseMapper
}

func CreateOTPMapper(bm BaseMapper) OTPMapper {
	return &otpMapper{bm}
}

func (o *otpMapper) SchemaToEntity(otp *mongo.OTP) *entity.OTP {
	return &entity.OTP{
		ID:        otp.ID.Hex(),
		ExpiredAt: otp.ExpiredAt,
		CreatedAt: otp.CreatedAt,
		UpdatedAt: otp.UpdatedAt,
		Code:      otp.Code,
		OwnerID:   otp.OwnerID.Hex(),
	}
}
