package schema

import (
	"Product/internal/model"
	"gorm.io/gorm"
)

type (
	Product struct {
		gorm.Model
		CompanyId   uint   `gorm:"uniqueIndex:product_unique_id"`
		DesignCode  string `gorm:"uniqueIndex:product_unique_id"`
		Description string
		Dimensions  []Dimension
		Themes      []Theme
	}

	Dimension struct {
		Id        uint
		ProductId uint   `gorm:"uniqueIndex:dimension_unique_id"`
		Size      string `gorm:"uniqueIndex:dimension_unique_id"`
	}

	Theme struct {
		Id        uint   `gorm:"column:id"`
		ProductId uint   `gorm:"uniqueIndex:theme_unique_id"`
		Color     string `gorm:"uniqueIndex:theme_unique_id"`
	}
)

func GetProduct(p *model.Product) *Product {
	return &Product{
		Model: gorm.Model{
			ID: p.Id,
		},
		CompanyId:   p.CompanyId,
		DesignCode:  p.DesignCode,
		Description: p.Description,
	}
}

func GetDimensions(product *model.Product) []Dimension {

	result := make([]Dimension, len(product.Dimensions))

	for i, d := range product.Dimensions {
		result[i] = Dimension{
			ProductId: product.Id,
			Size:      d,
		}
	}

	return result
}

func GetThemes(product *model.Product) []Theme {

	result := make([]Theme, len(product.Colors))

	for i, c := range product.Colors {
		result[i] = Theme{
			ProductId: product.Id,
			Color:     c,
		}
	}

	return result
}

//func ProductToSchema(c *Product) *model.Product {
//	return &model.Product{
//		DesignCode: c.DesignCode,
//		Colors:     c.Color,
//		Dimensions: []string{c.Dimension},
//	}
//}
//
//func ProductToCarpets(c *model.Product) []Product {
//	if len(c.Dimensions) == 0 {
//		return []Product{{}}
//	}
//
//	result := make([]Product, len(c.Dimensions))
//
//	for i, s := range c.Dimensions {
//		result[i] = Product{
//			DesignCode: c.DesignCode,
//			Color:      c.Colors,
//			Dimension:  s,
//		}
//	}
//
//	return result
//}
