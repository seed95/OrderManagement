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

func CarpetToModel(c *Carpet) *model.Carpet {
	return &model.Carpet{
		DesignCode: c.DesignCode,
		Color:      c.Color,
		Sizes:      []string{c.Dimension},
	}
}

func CarpetToSchema(c *model.Carpet) []Carpet {
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
