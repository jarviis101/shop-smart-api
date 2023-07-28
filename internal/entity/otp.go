package entity

import "time"

type OTP struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ExpiredAt time.Time
	ID        int64
	Code      string
	OwnerID   string
	IsUsed    bool
}
