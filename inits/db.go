package inits

import (
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"os"
)

var DB *gorm.DB

func DBInit() {
	var dsn string
	if os.Getenv("GO_ENV") == "test" {
		dsn = os.Getenv("TEST_DB_URL")
	} else {
		dsn = os.Getenv("DB_URL")
	}

	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		panic("failed to connect database")
	}

	DB = db
}
