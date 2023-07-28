package seeder

import (
	"shop-smart-api/internal/service"
)

type Seeder interface {
	Seed() error
}

type manager struct {
	userUseCase service.UserUseCase
}

func CreateSeeder(uc service.UserUseCase) Seeder {
	return &manager{uc}
}

func (s *manager) Seed() error {
	if err := s.seedUsers(); err != nil {
		return err
	}

	return nil
}

func (s *manager) seedUsers() error {
	return nil
}
