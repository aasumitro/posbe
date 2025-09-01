package model

type (
	Order struct {
		ID           int64  `json:"id"`
		CashierID    int64  `json:"cashier_id"`
		ShiftID      int64  `json:"shift_id"`
		TableID      int64  `json:"table_id"`
		TimeOpen     int64  `json:"time_open"`
		TimeClose    int64  `json:"time_close"`
		Customer     string `json:"customer"` // TODO: apply by id and store the customer data (id, name, phone:unique, notes)
		Gross        int64  `json:"gross"`    // Bruto
		Discount     int64  `json:"discount"`
		Net          int64  `json:"net"` // neto
		Tax          int64  `json:"tax"`
		Total        int64  `json:"total"`
		Type         string `json:"type"`
		Payment      int64  `json:"payment"`
		Change       int64  `json:"change"`
		Notes        string `json:"notes"`
		CancelReason string `json:"cancel_reason"`
		Status       string `json:"status"`
		CreatedAt    int64  `json:"created_at"`
		UpdatedAt    int64  `json:"updated_at"`
	}

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
