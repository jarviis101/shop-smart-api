package entity

import "time"

type Transaction struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ActionAt  time.Time
	ID        int64
	OwnerID   string
	TrxNumber string
	Value     string
	Status    bool
}
