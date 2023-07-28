package entity

import "time"

type User struct {
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ID             int64
	FirstName      string
	LastName       string
	MiddleName     string
	Phone          string
	Roles          []string
	OrganizationId *int64
}
