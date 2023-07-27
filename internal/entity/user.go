package entity

import "time"

type User struct {
	CreatedAt      time.Time
	UpdatedAt      time.Time
	ID             string
	FirstName      string
	LastName       string
	MiddleName     string
	Phone          string
	Roles          []string
	OrganizationId *string
}
