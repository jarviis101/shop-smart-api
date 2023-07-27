package seeder

import (
	"log"
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
	log.Println("Seeding successfully complete")
	return nil
}

func (s *manager) seedUsers() error {
	//s.userUseCase.
	return nil
}
