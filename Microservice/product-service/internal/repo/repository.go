package repo

import (
	"Product/internal/model"
)

type ProductRepo interface {
	CreateCarpet(p *model.Product) error
	GetAllCarpets(companyId uint) ([]model.Product, error)
}
