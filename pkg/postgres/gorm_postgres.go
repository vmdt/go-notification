package postgres

import (
	"errors"
	"fmt"
	"time"

	"github.com/cenkalti/backoff/v4"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

type Gorm struct {
	DB       *gorm.DB
	config   *PostgresConfig
}

func NewGorm(config *PostgresConfig) (*gorm.DB, error) {
	var dataSourceName string

	if config.DBName == "" {
		return nil, errors.New("DBName is required in the config")
	}

	err := createDB(config)
	if err != nil {
		return nil, err
	}

	dataSourceName = fmt.Sprintf("host=%s port=%d user=%s dbname=%s password=%s",
		config.Host,
		config.Port,
		config.User,
		config.DBName,
		config.Password,
	)

	bo := backoff.NewExponentialBackOff()
	bo.MaxElapsedTime = 10 * time.Second
	maxRetries := 5

	var gormDB *gorm.DB

	err = backoff.Retry(func () error {
		gormDB, err = gorm.Open(postgres.Open(dataSourceName), &gorm.Config{})
		
		if err != nil {
			return errors.New(fmt.Sprintf("error connecting to the database: %v", err))
		}

		return nil
	}, backoff.WithMaxRetries(bo, uint64(maxRetries)))

	return gormDB, err
}

func (db *Gorm) Close() {
	sqlDB, _ := db.DB.DB()
	_ = sqlDB.Close()
}

func Migrate(gorm *gorm.DB, types ...interface{}) error {
	for _, t := range types {
		err := gorm.AutoMigrate(t)
		if err != nil {
			return err
		}
	}
	return nil
}