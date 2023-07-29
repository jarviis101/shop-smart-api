package repository

import (
	"database/sql"
	"shop-smart-api/internal/entity"
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
