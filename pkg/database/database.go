package database

import (
	"github.com/TeerapatChan/inventory-management-api/config"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New() (*gorm.DB, error) {
	conf := config.Load()
	db, err := gorm.Open(postgres.Open(conf.Database.Url), &gorm.Config{})
	if err != nil {
		return nil, err
	}
	return db, nil
}
