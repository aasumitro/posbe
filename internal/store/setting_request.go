package store

import (
	"fmt"

	"github.com/aasumitro/posbe/internal/model"
	"github.com/gin-gonic/gin"
	"github.com/golodash/galidator/v2"
)

type SettingForm struct {
	Name            string  `form:"name" json:"name"`
	Phone           string  `form:"phone" json:"phone"`
	Email           string  `form:"email" json:"email"`
	Type            string  `form:"type" json:"type"`
	Address         string  `form:"address" json:"address"`
	Currency        string  `form:"currency" json:"currency"`
	ServiceCategory string  `form:"service_category" json:"service_category"`
	ServiceRate     float64 `form:"service_rate" json:"service_rate"`
	TaxCategory     string  `form:"tax_category" json:"tax_category"`
	TaxRate         float64 `form:"tax_rate" json:"tax_rate"`
	FeatureFloor    string  `form:"feature_floor" json:"feature_floor"`
}

func (f *SettingForm) Validate(ctx *gin.Context) interface{} {
	g := galidator.New()
	r := galidator.Rules{
		"Name":  g.R("name").Optional(),
		"Phone": g.R("phone").Optional().Phone(),
		"Email": g.R("email").Optional().Email(),
		"Type": g.R("type").Optional().Choices(
			"bar", "coffee", "restaurant"),
		"Address": g.R("address").Optional(),
		"Currency": g.R("currency").Optional().Choices(
			"IDR", "USD"),
		"ServiceCategory": g.R("service_category").Optional().Choices(
			"standard", "nominal"),
		"TaxCategory": g.R("tax_category").Optional().Choices(
			"standard", "nominal"),
	}
	if f.ServiceCategory == "standard" {
		r["ServiceRate"] = g.R("service_rate").Optional().Min(1).Max(100)
	}
	if f.TaxCategory == "standard" {
		r["TaxRate"] = g.R("tax_rate").Optional().Min(1).Max(100)
	}
	return g.ComplexValidator(r).Validate(ctx, f)
}

func (f *SettingForm) ToStoreModelMap() model.StoreSetting {
	settings := make(model.StoreSetting)

	assignIfNotEmpty := func(field string, value interface{}) {
		switch v := value.(type) {
		case string:
			if v != "" {
				settings[field] = v
			}
		case float64:
			if v != 0 {
				settings[field] = fmt.Sprintf("%g", v)
			}
		}
	}

	assignIfNotEmpty("name", f.Name)
	assignIfNotEmpty("phone", f.Phone)
	assignIfNotEmpty("email", f.Email)
	assignIfNotEmpty("type", f.Type)
	assignIfNotEmpty("address", f.Address)
	assignIfNotEmpty("currency", f.Currency)
	assignIfNotEmpty("service_category", f.ServiceCategory)
	assignIfNotEmpty("service_rate", f.ServiceRate)
	assignIfNotEmpty("tax_category", f.TaxCategory)
	assignIfNotEmpty("tax_rate", f.TaxRate)
	assignIfNotEmpty("feature_floor", f.FeatureFloor)

	return settings
}
