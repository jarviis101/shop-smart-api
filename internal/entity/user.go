package entity

import (
	"fmt"
	"time"
)

type Role string

func (r *Role) Scan(src interface{}) error {
	if src == nil {
		*r = ""
		return nil
	}

	byteArray, ok := src.([]uint8)
	if !ok {
		return fmt.Errorf("expected []uint8, got %T", src)
	}

	*r = Role(byteArray)

	return nil
}

const (
	UserRole   Role = "user"
	EditorRole Role = "editor"
	OwnerRole  Role = "owner"
)

type User struct {
	CreatedAt      time.Time
	UpdatedAt      time.Time
	Roles          []Role
	ID             int64
	FirstName      string
	LastName       string
	MiddleName     string
	Phone          string
	OrganizationID *int64
}

func (u *User) IsUser() bool {
	return u.containsRole(UserRole)
}

func (u *User) IsOwner() bool {
	return u.containsRole(UserRole) && u.containsRole(OwnerRole)
}

func (u *User) IsEditor() bool {
	return u.containsRole(UserRole) && u.containsRole(EditorRole)
}

func (u *User) containsRole(role Role) bool {
	for _, r := range u.Roles {
		if r == role {
			return true
		}
	}
	return false
}
