package product

import (
	"Product/internal"
	"Product/internal/derror"
	"Product/internal/model"
	"Product/internal/repo"
	"Product/internal/repo/product/schema"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"

	gorm_schema "gorm.io/gorm/schema"
)

type (
	productRepo struct {
		db     *gorm.DB
		config *internal.PostgresConfig
	}

	Setting struct {
		Config *internal.PostgresConfig
	}
)

var _ repo.ProductRepo = (*productRepo)(nil)

func New(s *Setting) (repo.ProductRepo, error) {
	productRepo := &productRepo{
		config: s.Config,
	}

	if err := productRepo.connect(); err != nil {
		return nil, err
	}

	if err := productRepo.migration(); err != nil {
		return nil, err
	}

	return productRepo, nil
}

func (r *productRepo) connect() error {

	postgresDB, err := gorm.Open(postgres.Open(r.config.PostgresUri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Disable default gorm log
		NamingStrategy: gorm_schema.NamingStrategy{
			TablePrefix:   "tbl_",
			SingularTable: true,
		},
	})

	if err != nil {
		return errors.New(fmt.Sprintf(derror.CreateProductRepoErrorFormat, err))
	}

	r.db = postgresDB

	return nil
}

func (r *productRepo) migration() error {

	if err := r.db.AutoMigrate(&schema.Product{}, &schema.Dimension{}, &schema.Theme{}); err != nil {
		return errors.New(fmt.Sprintf(derror.CreateProductRepoErrorFormat, err))
	}

	return nil

}

// CreateCarpet add one row for each size of product
func (r *productRepo) CreateCarpet(product *model.Product) error {

	if product == nil {
		return derror.NilProduct
	}

	err := r.db.Transaction(func(tx *gorm.DB) error {

		schemaProduct := schema.GetProduct(product)

		if err := tx.Create(schemaProduct).Error; err != nil {
			return derror.New(derror.InternalServer, err.Error())
		}
		product.Id = schemaProduct.ID

		if len(product.Dimensions) != 0 {
			schemaDimension := schema.GetDimensions(product)
			if err := tx.Create(schemaDimension).Error; err != nil {
				return derror.New(derror.InternalServer, err.Error())
			}
		}

		if len(product.Colors) != 0 {
			schemaThemes := schema.GetThemes(product)
			if err := tx.Create(schemaThemes).Error; err != nil {
				return derror.New(derror.InternalServer, err.Error())
			}
		}

		return nil
	})

	if err != nil {
		return derror.New(derror.InternalServer, err.Error())
	}

	return nil
}

func (r *productRepo) GetAllCarpets(companyId uint) ([]model.Product, error) {

	var schemaProducts []schema.Product

	if err := r.db.Where("company_id = ?", companyId).Find(&schemaProducts).Error; err != nil {
		return nil, derror.New(derror.InternalServer, err.Error())
	}

	result := make([]model.Product, len(schemaProducts))

	return result, nil
}
