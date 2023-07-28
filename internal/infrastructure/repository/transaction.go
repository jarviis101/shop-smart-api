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

func (r *transactionRepository) GetByOwner(owner int64) ([]*entity.Transaction, error) {
	return nil, nil
}
