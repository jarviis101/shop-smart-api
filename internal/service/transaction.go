package service

import (
	"shop-smart-api/internal/entity"
	"shop-smart-api/internal/service/transaction"
)

type transactionService struct {
	finder transaction.Finder
}

func CreateTransactionService(f transaction.Finder) TransactionService {
	return &transactionService{f}
}

func (s *transactionService) GetTransactions(owner *entity.User) ([]*entity.Transaction, error) {
	return s.finder.FindByOwner(owner.ID)
}
