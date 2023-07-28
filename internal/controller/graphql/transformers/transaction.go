package transformers

import (
	"shop-smart-api/internal/controller/graphql/graph/model"
	"shop-smart-api/internal/entity"
	"strconv"
)

type TransactionTransformer interface {
	TransformToModel(u *entity.Transaction) *model.Transaction
	TransformManyToModel(u []*entity.Transaction) []*model.Transaction
}

type transactionTransformer struct{}

func CreateTransactionTransformer() TransactionTransformer {
	return &transactionTransformer{}
}

func (t *transactionTransformer) TransformManyToModel(u []*entity.Transaction) []*model.Transaction {
	var transactions []*model.Transaction
	for _, user := range u {
		m := t.TransformToModel(user)

		transactions = append(transactions, m)
	}

	return transactions
}

func (t *transactionTransformer) TransformToModel(u *entity.Transaction) *model.Transaction {
	return &model.Transaction{
		ID: strconv.Itoa(int(u.ID)),
	}
}
