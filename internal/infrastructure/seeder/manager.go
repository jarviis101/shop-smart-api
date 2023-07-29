package seeder

import (
	"github.com/go-faker/faker/v4"
	"log"
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/infrastructure/repository"
	"time"
)

const (
	countOfUsers        = 10
	countOfTransactions = 10
)

type Seeder interface {
	Seed() error
}

type manager struct {
	userRepository         repository.UserRepository
	organizationRepository repository.OrganizationRepository
	transactionRepository  repository.TransactionRepository
	users                  []*entity.User
}

func CreateSeeder(
	ur repository.UserRepository,
	or repository.OrganizationRepository,
	tr repository.TransactionRepository,
) Seeder {
	return &manager{userRepository: ur, organizationRepository: or, transactionRepository: tr}
}

func (s *manager) Seed() error {
	if err := s.seedUsers(); err != nil {
		return err
	}

	if err := s.seedOrganization(); err != nil {
		return err
	}

	if err := s.seedTransactions(); err != nil {
		return err
	}

	return nil
}

func (s *manager) seedUsers() error {
	for i := 0; i < countOfUsers; i++ {
		model := User{}
		if err := faker.FakeData(&model); err != nil {
			return err
		}

		user, err := s.userRepository.Store(
			model.Phone,
			model.FirstName,
			model.LastName,
			"",
			[]entity.Role{entity.UserRole},
		)
		if err != nil {
			return err
		}

		s.users = append(s.users, user)
	}

	log.Println("User seeding completed")

	return nil
}

func (s *manager) seedOrganization() error {
	model := Organization{}
	if err := faker.FakeData(&model); err != nil {
		return err
	}

	owner := s.users[0]
	role := entity.OwnerRole
	code := generateRandomCode()
	organization, err := s.organizationRepository.Store(model.Name, code, code, code, owner.ID)
	if err != nil {
		return err
	}

	for i, user := range s.users {
		if i == 0 {
			if _, err := s.userRepository.AddOrganization(user.ID, organization.ID, &role); err != nil {
				return err
			}
			continue
		}
		if _, err := s.userRepository.AddOrganization(user.ID, organization.ID, nil); err != nil {
			return err
		}
	}

	log.Println("Organization seeding completed")

	return nil
}

func (s *manager) seedTransactions() error {
	for _, user := range s.users {
		for i := 0; i < countOfTransactions; i++ {
			model := Transaction{}
			if err := faker.FakeData(&model); err != nil {
				return err
			}

			value := generateRandomFloatValue(10, 200)
			actionedAt := time.Now()

			if i%2 == 0 {
				if _, err := s.transactionRepository.Store(
					user.ID,
					model.TrxNumber,
					value,
					nil,
					false,
				); err != nil {
					return err
				}

				continue
			}

			if _, err := s.transactionRepository.Store(
				user.ID,
				model.TrxNumber,
				value,
				&actionedAt,
				true,
			); err != nil {
				return err
			}
		}
	}

	log.Println("Transaction seeding completed")

	return nil
}
