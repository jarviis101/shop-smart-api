package seeder

import (
	"log"
	"shop-smart-api/internal/usecase"
)

type Seeder interface {
	Seed() error
}

type manager struct {
	uc usecase.UserUseCase
}

func CreateSeeder(uc usecase.UserUseCase) Seeder {
	return &manager{uc}
}

func (s *manager) Seed() error {

	log.Println("Seeding successfully complete")
	return nil
}
