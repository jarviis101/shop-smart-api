package repository

import (
	"database/sql"
	"shop-smart-api/internal/entity"
	"time"
)

type transactionRepository struct {
	database *sql.DB
}

func CreateTransactionRepository(db *sql.DB) TransactionRepository {
	return &transactionRepository{db}
}

func (r *transactionRepository) GetByOwner(id int64) ([]*entity.Transaction, error) {
	return r.executeQuery("SELECT * FROM transactions WHERE owner_id = $1", id)
}

func (r *transactionRepository) Store(
	owner int64,
	trxNumber string,
	value float64,
	actionedAt *time.Time,
	status bool,
) (*entity.Transaction, error) {
	return r.executeQueryRow(`
		INSERT INTO transactions (trx_number, value, owner_id, actioned_at, status) VALUES ($1, $2, $3, $4, $5)
		RETURNING id, trx_number, value, status, created_at, updated_at, actioned_at, owner_id
	`, trxNumber, value, owner, actionedAt, status)
}

func (r *transactionRepository) executeQuery(query string, args ...any) ([]*entity.Transaction, error) {
	rows, err := r.database.Query(query, args...)
	if err != nil {
		return nil, err
	}
	defer rows.Close()

	var transactions []*entity.Transaction
	for rows.Next() {
		var transaction entity.Transaction

		if err := rows.Scan(
			&transaction.ID,
			&transaction.TrxNumber,
			&transaction.Value,
			&transaction.Status,
			&transaction.CreatedAt,
			&transaction.UpdatedAt,
			&transaction.ActionedAt,
			&transaction.OwnerID,
		); err != nil {
			continue
		}

		transactions = append(transactions, &transaction)
	}

	return transactions, nil
}

func (r *transactionRepository) executeQueryRow(query string, args ...any) (*entity.Transaction, error) {
	var transaction entity.Transaction

	err := r.database.QueryRow(query, args...).Scan(
		&transaction.ID,
		&transaction.TrxNumber,
		&transaction.Value,
		&transaction.Status,
		&transaction.CreatedAt,
		&transaction.UpdatedAt,
		&transaction.ActionedAt,
		&transaction.OwnerID,
	)
	if err != nil {
		return nil, err
	}

	return &transaction, nil
}
