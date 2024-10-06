package model

type (
	Order struct{}

	OrderProduct struct{}

	OrderProductAddon struct{}

	IOrderService interface {
		// Make()
		// Update()
		// RemoveProduct()
		// RemoveAddon()
	}

	ITransactionService interface {
		// List()
		// Detail()
	}
)
