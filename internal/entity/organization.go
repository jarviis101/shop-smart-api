package entity

import "time"

type Organization struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        int64
	OwnerID   int64
	Name      string
	ORGN      string
	KPP       string
	INN       string
}
