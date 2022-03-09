package repo

import (
	"Product/internal/model"
)

type CarpetRepo interface {
	CreateCarpet(p *model.Product) error
	GetAllCarpets() ([]model.Product, error)
}
