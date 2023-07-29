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
	return r.executeQueryRow("SELECT * FROM users WHERE id = $1", id)
}

func (r *userRepository) GetByPhone(phone string) (*entity.User, error) {
	return r.executeQueryRow("SELECT * FROM users WHERE phone = $1", phone)
}

func (r *userRepository) GetByOrganization(id int64) ([]*entity.User, error) {
	return r.executeQuery("SELECT * FROM users WHERE organization_id = $1", id)
}

func (r *userRepository) GetAll() ([]*entity.User, error) {
	return r.executeQuery("SELECT * FROM users")
}

func (r *userRepository) Store(
	phone, firstName, lastName, middleName string,
	roles []entity.Role,
) (*entity.User, error) {
	return r.executeQueryRow(`
		INSERT INTO users (first_name, last_name, middle_name, phone, roles) VALUES ($1, $2, $3, $4, $5)
		RETURNING id, first_name, last_name, middle_name, phone, created_at, updated_at, organization_id, roles
	`, firstName, lastName, middleName, phone, pq.Array(roles))
}

func (r *userRepository) UpdateUser(id int64, firstName, lastName, middleName string) (*entity.User, error) {
	return r.executeQueryRow(`
		UPDATE users SET first_name = $1, last_name = $2, middle_name = $3 WHERE id = $4
		RETURNING id, first_name, last_name, middle_name, phone, created_at, updated_at, organization_id, roles
	`, firstName, lastName, middleName, id)
}

func (r *userRepository) AddOrganization(id, organization int64, role *entity.Role) (*entity.User, error) {
	if role == nil {
		return r.executeQueryRow(`
			UPDATE users SET organization_id = $1 WHERE id = $2
			RETURNING id, first_name, last_name, middle_name, phone, created_at, updated_at, organization_id, roles
		`, organization, id)
	}

	return r.executeQueryRow(`
		UPDATE users SET organization_id = $1, roles = array_append(roles, $2) WHERE id = $3
		RETURNING id, first_name, last_name, middle_name, phone, created_at, updated_at, organization_id, roles
		`, organization, role, id)
}

func (r *userRepository) executeQuery(query string, args ...any) ([]*entity.User, error) {
	rows, err := r.database.Query(query, args...)
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

func (r *userRepository) executeQueryRow(query string, args ...any) (*entity.User, error) {
	var user entity.User

	err := r.database.QueryRow(query, args...).Scan(
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
