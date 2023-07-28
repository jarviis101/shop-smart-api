package seeder

import (
	"github.com/go-faker/faker/v4"
	"log"
	"shop-smart-api/internal/infrastructure/repository"
)

type Seeder interface {
	Seed() error
}

type manager struct {
	userRepository repository.UserRepository
}

func CreateSeeder(ur repository.UserRepository) Seeder {
	return &manager{ur}
}

func (s *manager) Seed() error {
	if err := s.seedUsers(); err != nil {
		return err
	}

	return nil
}

func (s *manager) seedUsers() error {
	if err := s.userRepository.Truncate(); err != nil {
		return err
	}

	for i := 0; i < 10; i++ {
		user := User{}
		if err := faker.FakeData(&user); err != nil {
			return err
		}

		if _, err := s.userRepository.Store(
			user.Phone,
			user.FirstName,
			user.LastName,
			"",
			[]string{"user"},
		); err != nil {
			return err
		}
	}

	log.Println("User seeding completed")
	return nil
}
