package repo

import (
	"Product/internal/model"
)

type ProductRepo interface {
	CreateProduct(p *model.Product) (*model.Product, error)
	GetAllCarpet(companyId uint) ([]model.Carpet, error)
	DeleteProduct(productId uint) error
}
