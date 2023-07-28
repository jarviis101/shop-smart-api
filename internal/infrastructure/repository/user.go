package repository

import (
	"database/sql"
	"github.com/lib/pq"
	"shop-smart-api/internal/entity"
)

type userRepository struct {
	database *sql.DB
}

func CreateUserRepository(db *sql.DB) UserRepository {
	return &userRepository{db}
}

func (r *userRepository) Get(id int64) (*entity.User, error) {
	var user entity.User

	err := r.database.QueryRow("SELECT * FROM users WHERE id = $1", id).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.MiddleName,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.OrganizationID,
		pq.Array(&user.Roles),
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByPhone(phone string) (*entity.User, error) {
	var user entity.User

	err := r.database.QueryRow(
		"SELECT * FROM users WHERE phone = $1",
		phone,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.MiddleName,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.OrganizationID,
		pq.Array(&user.Roles),
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) GetByOrganization(id int64) ([]*entity.User, error) {
	rows, err := r.database.Query("SELECT * FROM users WHERE organization_id = $1", id)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User

		if err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.MiddleName,
			&user.Phone,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.OrganizationID,
			pq.Array(&user.Roles),
		); err != nil {
			continue
		}

		users = append(users, &user)
	}

	return users, nil
}

func (r *userRepository) GetAll() ([]*entity.User, error) {
	rows, err := r.database.Query("SELECT * FROM users")
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var users []*entity.User
	for rows.Next() {
		var user entity.User

		if err := rows.Scan(
			&user.ID,
			&user.FirstName,
			&user.LastName,
			&user.MiddleName,
			&user.Phone,
			&user.CreatedAt,
			&user.UpdatedAt,
			&user.OrganizationID,
			pq.Array(&user.Roles),
		); err != nil {
			continue
		}

		users = append(users, &user)
	}

	return users, nil
}

func (r *userRepository) Store(
	phone, firstName, lastName, middleName string,
	roles []string,
) (*entity.User, error) {
	var user entity.User

	err := r.database.QueryRow(
		`INSERT INTO users (first_name, last_name, middle_name, phone, roles) VALUES ($1, $2, $3, $4, $5) 
		RETURNING id, first_name, last_name, middle_name, phone, created_at, updated_at, organization_id, roles
		`,
		firstName, lastName, middleName, phone, pq.Array(roles),
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.MiddleName,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.OrganizationID,
		pq.Array(&user.Roles),
	)

	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) UpdateUser(id int64, firstName, lastName, middleName string) (*entity.User, error) {
	var user entity.User

	err := r.database.QueryRow(
		`UPDATE users SET first_name = $1, last_name = $2, middle_name = $3 WHERE id = $4
		RETURNING id, first_name, last_name, middle_name, phone, created_at, updated_at, organization_id, roles
		`,
		firstName, lastName, middleName, id,
	).Scan(
		&user.ID,
		&user.FirstName,
		&user.LastName,
		&user.MiddleName,
		&user.Phone,
		&user.CreatedAt,
		&user.UpdatedAt,
		&user.OrganizationID,
		pq.Array(&user.Roles),
	)
	if err != nil {
		return nil, err
	}

	return &user, nil
}

func (r *userRepository) Truncate() error {
	if _, err := r.database.Exec("TRUNCATE TABLE users CASCADE"); err != nil {
		return err
	}

	return nil
}
