package repository

import (
	"database/sql"
	"log"
	"shop-smart-api/internal/entity"
	"time"
)

type otpRepository struct {
	database *sql.DB
}

func CreateOTPRepository(db *sql.DB) OTPRepository {
	return &otpRepository{db}
}

func (o otpRepository) Store(owner, code string) (*entity.OTP, error) {
	var otp entity.OTP

	err := o.database.QueryRow(
		"INSERT INTO otp (code, owner_id, expired_at) values ($1, $2, $3) returning (*)",
		code, owner, time.Now().Add(time.Minute*5),
	).Scan(&otp)
	if err != nil {
		return nil, err
	}

	return &otp, nil
}

func (o otpRepository) GetByOwnerAndCode(owner, code string) (*entity.OTP, error) {
	var otp entity.OTP

	rows, err := o.database.Query(
		"SELECT * FROM otp WHERE owner_id = $1 AND code = $2",
		owner, code,
	)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	for rows.Next() {
		err := rows.Scan(&otp)
		if err != nil {
			log.Printf("Error: %s", err)
			continue
		}
	}

	return &otp, nil
}

func (o otpRepository) UseOTP(otp *entity.OTP) error {
	if _, err := o.database.Exec("UPDATE otp SET is_used = true WHERE id = $1", otp.ID); err != nil {
		return err
	}

	return nil
}
