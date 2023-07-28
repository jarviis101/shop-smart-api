package transformers

import (
	"shop-smart-api/internal/controller/graphql/graph/model"
	"shop-smart-api/internal/entity"
	"strconv"
	"time"
)

type OrganizationTransformer interface {
	TransformToModel(o *entity.Organization) *model.Organization
}

type organizationTransformer struct{}

func CreateOrganizationTransformer() OrganizationTransformer {
	return &organizationTransformer{}
}

func (t *organizationTransformer) TransformToModel(o *entity.Organization) *model.Organization {
	return &model.Organization{
		ID:        strconv.Itoa(int(o.ID)),
		Name:      o.Name,
		Orgn:      o.ORGN,
		Kpp:       o.KPP,
		Inn:       o.INN,
		OwnerID:   strconv.Itoa(int(o.OwnerID)),
		CreatedAt: o.CreatedAt.Format(time.RFC822),
		UpdatedAt: o.UpdatedAt.Format(time.RFC822),
	}
}
