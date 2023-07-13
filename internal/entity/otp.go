package entity

import "time"

type OTP struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiredAt time.Time
	ID        string
	Code      string
	OwnerID   string
}
