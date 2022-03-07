package carpet

import (
	"Product/internal"
	"Product/internal/derror"
	"Product/internal/model"
	"Product/internal/repo"
	"Product/internal/repo/carpet/schema"
	"errors"
	"fmt"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

type (
	carpetRepo struct {
		db     *gorm.DB
		config *internal.PostgresConfig
	}

	Setting struct {
		Config *internal.PostgresConfig
	}
)

var _ repo.CarpetRepo = (*carpetRepo)(nil)

func New(s *Setting) (repo.CarpetRepo, error) {
	productRepo := &carpetRepo{
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

func (r *carpetRepo) connect() error {

	postgresDB, err := gorm.Open(postgres.Open(r.config.PostgresUri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Disable default gorm log
	})

	if err != nil {
		return errors.New(fmt.Sprintf(derror.CreateProductRepoErrorFormat, err))
	}

	r.db = postgresDB

	return nil
}

func (r *carpetRepo) migration() error {

	if err := r.db.AutoMigrate(&schema.Carpet{}); err != nil {
		return errors.New(fmt.Sprintf(derror.CreateProductRepoErrorFormat, err))
	}

	return nil

}

func (r *carpetRepo) NewCarpet(c *model.Carpet) error {

	carpets := schema.CarpetToSchema(c)

	if err := r.db.Create(carpets).Error; err != nil {
		return derror.New(derror.InternalServer, err.Error())
	}
	return nil
}
