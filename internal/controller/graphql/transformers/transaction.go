package transformers

import (
	"shop-smart-api/internal/controller/graphql/graph/model"
	"shop-smart-api/internal/entity"
	"strconv"
	"time"
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

func (t *transactionTransformer) TransformToModel(m *entity.Transaction) *model.Transaction {
	actionedAt := t.resolveActionedAt(m.ActionedAt)

	return &model.Transaction{
		ID:         strconv.Itoa(int(m.ID)),
		OwnerID:    strconv.Itoa(int(m.OwnerID)),
		Value:      m.Value,
		TrxNumber:  m.TrxNumber,
		Status:     m.Status,
		CreatedAt:  m.CreatedAt.Format(time.RFC822),
		UpdatedAt:  m.UpdatedAt.Format(time.RFC822),
		ActionedAt: actionedAt,
	}
}

func (t *transactionTransformer) resolveActionedAt(actionedAt *time.Time) *string {
	if actionedAt == nil {
		return nil
	}

	timestamp := actionedAt.Format(time.RFC822)
	return &timestamp
}
