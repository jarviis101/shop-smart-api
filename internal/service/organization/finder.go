package organization

import (
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/infrastructure/repository"
)

type Finder interface {
	Find(id int64) (*entity.Organization, error)
}

type finder struct {
	repository repository.OrganizationRepository
}

func CreateFinder(r repository.OrganizationRepository) Finder {
	return &finder{r}
}

func (f *finder) Find(id int64) (*entity.Organization, error) {
	return f.repository.Get(id)
}
