package entity

import "time"

type Transaction struct {
	CreatedAt  time.Time
	UpdatedAt  time.Time
	ActionedAt time.Time
	ID         int64
	OwnerID    int64
	Value      float64
	TrxNumber  string
	Status     bool
}
