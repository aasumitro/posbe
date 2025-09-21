package model

type (
	Unit struct {
		ID        int64  `json:"id"`
		Magnitude string `json:"magnitude"` // e.g: mass [length, time]
		Name      string `json:"name"`      // e.g: kilogram [metre, second]
		Symbol    string `json:"symbol"`    // e.g: kg [m, s]
		Usage     int64  `json:"usage,omitempty"`
	}

	Category struct {
		ID          int64          `json:"id"`
		Name        string         `json:"name"`
		Subcategory []*Subcategory `json:"subcategories,omitempty"`
		Usage       int64          `json:"usage,omitempty"`
	}

	Subcategory struct {
		ID         int64     `json:"id"`
		CategoryID int64     `json:"category_id"`
		Name       string    `json:"name"`
		Usage      int64     `json:"usage,omitempty"`
		Category   *Category `json:"category,omitempty"`
	}
)
