package service

import (
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/service/organization"
)

type organizationService struct {
	finder organization.Finder
}

func CreateOrganizationService(f organization.Finder) OrganizationService {
	return &organizationService{f}
}

func (s *organizationService) Get(id int64) (*entity.Organization, error) {
	return s.finder.Find(id)
}
