package order

import "github.com/aasumitro/posbe/internal/model"

type transactionRepository struct {
	db model.IPgxPool
}

func NewTransactionRepository(db model.IPgxPool) ITransactionRepository {
	return &transactionRepository{db: db}
}
