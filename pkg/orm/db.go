package orm

import (
	"errors"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/gorm/logger"
)

var dbClient *gorm.DB

func GetDbClient() *gorm.DB {
	if dbClient == nil {
		panic(errors.New("db client is not initialized"))
	}
	return dbClient
}

func InitDbClient(dsn string) error {
	if dbClient != nil {
		return nil
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{
		Logger: logger.Default.LogMode(logger.Info),
	})
	if err != nil {
		return err
	}
	dbClient = db
	return nil
}

func NewDbClient(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}


