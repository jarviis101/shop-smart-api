package entity

import "time"

type Organization struct {
	CreatedAt time.Time
	UpdatedAt time.Time
	ID        string
	OwnerID   string
	Name      string
	ORGN      string
	KPP       string
	INN       string
}
