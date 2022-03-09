package schema

import (
	"Product/internal/model"
	"gorm.io/gorm"
)

type (
	Carpet struct {
		gorm.Model
		DesignCode string `gorm:"column:design_code;uniqueIndex:carpet_unique_id"`
		Color      string `gorm:"column:color;uniqueIndex:carpet_unique_id"`
		Dimension  string `gorm:"column:dimension;uniqueIndex:carpet_unique_id"`
	}
)

func CarpetToProduct(c *Carpet) *model.Product {
	return &model.Product{
		DesignCode: c.DesignCode,
		Color:      c.Color,
		Sizes:      []string{c.Dimension},
	}
}

func ProductToCarpets(c *model.Product) []Carpet {
	if len(c.Sizes) == 0 {
		return []Carpet{{}}
	}

	result := make([]Carpet, len(c.Sizes))

	for i, s := range c.Sizes {
		result[i] = Carpet{
			DesignCode: c.DesignCode,
			Color:      c.Color,
			Dimension:  s,
		}
	}

	return result
}
