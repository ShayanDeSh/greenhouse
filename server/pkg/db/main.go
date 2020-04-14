package db

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/postgres"
	"github.com/nothingrealm/greenhouse/server/pkg/db/models"
)

func Init() (db *gorm.DB, err error) {
	db, err = gorm.Open("postgres",
		"host=localhost port=5432 user=postgres dbname=greenhouse password=1234")
	migrate(db)
	return
}

func migrate(db *gorm.DB) {
	db.AutoMigrate(&models.Plant{})
}
