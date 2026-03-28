package model

// info:
// Gross: (b) Brutto
// Net: (b) netto
// TODO:
//  replace customer: apply by id and store the customer data (id, name, phone:unique, notes)

type (
	Order struct {
		ID            int64  `json:"id"`
		CashierID     int64  `json:"cashier_id"`
		ShiftID       int64  `json:"shift_id"`
		TableID       int64  `json:"table_id"`
		TimeOpen      int64  `json:"time_open"`
		TimeClose     int64  `json:"time_close"`
		Customer      string `json:"customer"`
		Gross         int64  `json:"gross"`
		Discount      int64  `json:"discount"`
		Net           int64  `json:"net"`
		Tax           int64  `json:"tax"`
		Total         int64  `json:"total"`
		PaymentMethod string `json:"payment_method"`
		PaymentAmount int64  `json:"payment_amount"`
		Change        int64  `json:"change"`
		Notes         string `json:"notes"`
		CancelReason  string `json:"cancel_reason"`
		Status        string `json:"status"`
		CreatedAt     int64  `json:"created_at"`
		UpdatedAt     int64  `json:"updated_at"`
	}

	OrderProduct struct {
		ID            int64   `json:"id"`
		OrderID       int64   `json:"order_id"`
		ProductID     int64   `json:"product_id"`
		CategoryID    int64   `json:"category_id"`
		SubcategoryID int64   `json:"subcategory_id"`
		Name          string  `json:"name"`
		Quantity      int     `json:"quantity"`
		Price         float64 `json:"price"`
		Net           float64 `json:"net"`
		Notes         string  `json:"notes"`
		CreatedAt     int64   `json:"created_at"`
		UpdatedAt     int64   `json:"updated_at"`
	}

	OrderProductOption struct {
		ID             int64  `json:"id"`
		OrderID        int64  `json:"order_id"`
		OrderProductID int64  `json:"order_product_id"`
		VariantID      int64  `json:"variant_id"`
		Name           string `json:"name"`
		Value          string `json:"value"`
		Price          string `json:"price"`
	}

	OrderProductAddon struct {
		ID             int64   `json:"id"`
		OrderID        int64   `json:"order_id"`
		OrderProductID int64   `json:"order_product_id"`
		AddonID        int64   `json:"addon_id"`
		Name           string  `json:"name"`
		Quantity       int     `json:"quantity"`
		Price          float64 `json:"price"`
		Net            float64 `json:"net"`
		Notes          string  `json:"notes"`
		CreatedAt      int64   `json:"created_at"`
		UpdatedAt      int64   `json:"updated_at"`
	}
)
