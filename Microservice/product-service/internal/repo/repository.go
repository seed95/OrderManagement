package repo

import (
	"Product/internal/model"
)

type CarpetRepo interface {
	NewCarpet(p *model.Carpet) error
}
