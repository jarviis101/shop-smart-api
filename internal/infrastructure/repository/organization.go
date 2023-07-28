package repository

import (
	"database/sql"
	_ "github.com/lib/pq"
	"shop-smart-api/internal/entity"
)

type organizationRepository struct {
	database *sql.DB
}

func CreateOrganizationRepository(db *sql.DB) OrganizationRepository {
	return &organizationRepository{db}
}

func (r *organizationRepository) Get(id int64) (*entity.Organization, error) {
	var organization entity.Organization

	err := r.database.QueryRow("SELECT * FROM organizations WHERE id = $1", id).Scan(
		&organization.ID,
		&organization.Name,
		&organization.KPP,
		&organization.ORGN,
		&organization.INN,
		&organization.CreatedAt,
		&organization.UpdatedAt,
		&organization.OwnerID,
	)
	if err != nil {
		return nil, err
	}

	return &organization, nil
}
