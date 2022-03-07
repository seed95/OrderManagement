package carpet

import (
	"Product/internal"
	"Product/internal/repo/carpet/schema"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

func NewCarpetRepoMock() (*carpetRepo, error) {

	mock := carpetRepo{
		config: &internal.PostgresConfig{
			PostgresUri: "host=localhost user=dadashi password=dadashi@1400 dbname=dev_db port=5432 sslmode=disable",
		},
	}

	postgresDB, err := gorm.Open(postgres.Open(mock.config.PostgresUri), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Silent), // Disable default gorm log
	})

	if err != nil {
		return nil, err
	}

	if err = postgresDB.Migrator().DropTable(&schema.Carpet{}); err != nil {
		return nil, err
	}

	if err = postgresDB.AutoMigrate(&schema.Carpet{}); err != nil {
		return nil, err
	}

	mock.db = postgresDB

	return &mock, nil
}
