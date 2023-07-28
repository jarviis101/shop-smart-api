package transaction

import (
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/infrastructure/repository"
)

type Finder interface {
	FindByOwner(id int64) ([]*entity.Transaction, error)
}

type finder struct {
	repository repository.TransactionRepository
}

func CreateFinder(r repository.TransactionRepository) Finder {
	return &finder{r}
}

func (f *finder) FindByOwner(id int64) ([]*entity.Transaction, error) {
	return f.repository.GetByOwner(id)
}
