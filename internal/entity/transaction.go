package entity

import "time"

type Transaction struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ActionAt  time.Time
	ID        string
	OwnerID   string
	TrxNumber string
	Value     string
	Status    bool
}
