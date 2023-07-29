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

func (r *organizationRepository) Store(name, kpp, orgn, inn string, owner int64) (*entity.Organization, error) {
	var organization entity.Organization

	err := r.database.QueryRow(
		`INSERT INTO organizations (name, kpp, orgn, inn, owner_id) VALUES ($1, $2, $3, $4, $5) 
		RETURNING id, name, kpp, orgn, inn, created_at, updated_at, owner_id
		`,
		name, kpp, orgn, inn, owner,
	).Scan(
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
