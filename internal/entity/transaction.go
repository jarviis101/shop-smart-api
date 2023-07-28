package entity

import "time"

type Transaction struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ActionAt  time.Time
	ID        int64
	OwnerID   int64
	TrxNumber string
	Value     string
	Status    bool
}
