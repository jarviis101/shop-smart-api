package user

import (
	"shop-smart-api/internal/infrastructure/repository"
)

type Collector interface {
}

type collector struct {
	repository repository.UserRepository
}

func CreateCollector(r repository.UserRepository) Collector {
	return &collector{r}
}
