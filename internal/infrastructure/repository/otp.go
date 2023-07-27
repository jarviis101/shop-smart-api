package repository

import (
	"database/sql"
	"shop-smart-api/internal/entity"
	"time"
)

type otpRepository struct {
	database *sql.DB
}

func CreateOTPRepository(db *sql.DB) OTPRepository {
	return &otpRepository{db}
}

func (r *otpRepository) Store(owner, code string) (*entity.OTP, error) {
	var otp entity.OTP

	err := r.database.QueryRow(
		`INSERT INTO otp (code, owner_id, expired_at) VALUES ($1, $2, $3) 
		RETURNING id, code, is_used, owner_id, created_at, updated_at, expired_at
		`,
		code, owner, time.Now().Add(time.Minute*5),
	).Scan(
		&otp.ID,
		&otp.Code,
		&otp.IsUsed,
		&otp.OwnerID,
		&otp.CreatedAt,
		&otp.UpdatedAt,
		&otp.ExpiredAt,
	)
	if err != nil {
		return nil, err
	}

	return &otp, nil
}

func (r *otpRepository) GetByOwnerAndCode(owner, code string) (*entity.OTP, error) {
	var otp entity.OTP

	err := r.database.QueryRow(
		"SELECT * FROM otp WHERE owner_id = $1 AND code = $2",
		owner, code,
	).Scan(
		&otp.ID,
		&otp.Code,
		&otp.IsUsed,
		&otp.OwnerID,
		&otp.CreatedAt,
		&otp.UpdatedAt,
		&otp.ExpiredAt,
	)
	if err != nil {
		return nil, err
	}

	return &otp, nil
}

func (r *otpRepository) Use(otp *entity.OTP) error {
	if _, err := r.database.Exec("UPDATE otp SET is_used = true WHERE id = $1", otp.ID); err != nil {
		return err
	}

	return nil
}
