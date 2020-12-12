package sqlite

import (
	model "core-backend/pkg/models"

	"github.com/jinzhu/gorm"
	// Register some standard stuff
	_ "github.com/mattn/go-sqlite3"
)

// MigrateDatabase migrate all table in database
func MigrateDatabase() {
	database := ConnectDatabase()

	if !database.HasTable(&model.User{}) {
		database.CreateTable(&model.User{})
		database.Set("gorm:table_options", "ENGINE=InnoDB").CreateTable(&model.User{})
	}
}

// ConnectDatabase connect database
func ConnectDatabase() *gorm.DB {
	database, err := gorm.Open("sqlite3", "./test.db")
	database.LogMode(true)

	if err != nil {
		panic(err)
	}

	return database
}
