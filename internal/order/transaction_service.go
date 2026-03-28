package order

type transactionService struct {
	repository ITransactionRepository
}

func NewTransactionService(repository ITransactionRepository) ITransactionService {
	return &transactionService{repository: repository}
}
