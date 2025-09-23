package model

import (
	"database/sql"
)

const (
	StoreSettingCacheKey = "store_settings"
)

type (
	// StorePref store setting
	StorePref struct {
		Key       string `json:"key" form:"key" binding:"required"`
		Value     string `json:"value" form:"value" binding:"required"`
		CreatedAt sql.NullInt64
		UpdatedAt sql.NullInt64
	}

	StoreSetting map[string]interface{}
)
