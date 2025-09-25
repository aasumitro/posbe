package model

type (
	Addon struct {
		ID          int64   `json:"id"`
		Price       float64 `json:"price"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Usage       int64   `json:"usage"`
	}

	Product struct {
		ID            int64  `json:"id"`
		CategoryID    int64  `json:"category_id" `
		SubcategoryID int64  `json:"subcategory_id"`
		SKU           string `json:"sku"`
		Image         string `json:"image"`
		Name          string `json:"name"`
		Description   string `json:"description"`
		// embedding data
		Category    *Category    `json:"category,omitempty"`
		Subcategory *Subcategory `json:"subcategory,omitempty"`
		Variants    []*Variant   `json:"variants,omitempty"`
		// Sales counts by period
		SalesToday    int64 `json:"sales_today"`
		SalesThisWeek int64 `json:"sales_this_week"`
		SalesThisYear int64 `json:"sales_this_year"`
	}

	Variant struct {
		ID          int64   `json:"id"`
		ProductID   int64   `json:"product_id"`
		UnitID      int64   `json:"unit_id"`
		UnitSize    float64 `json:"unit_size"`
		Type        string  `json:"type"`
		Name        string  `json:"name"`
		Description string  `json:"description"`
		Price       float64 `json:"price"`
		// embedding data
		Unit *Unit `json:"unit,omitempty"`
	}
)
