package repo

import (
	"Product/internal/model"
)

type ProductRepo interface {
	CreateProduct(p *model.Product) error
	GetAllCarpet(companyId uint) ([]model.Carpet, error)
}
