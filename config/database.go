package config

import (
	"github.com/karledenstal/golang-todo/entities"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var Database *gorm.DB
var DatabaseName string = "todo.db"

func Connect() error {
	var err error

	Database, err = gorm.Open((sqlite.Open(DatabaseName)), &gorm.Config{})

	if err != nil {
		panic(err)
	}

	Database.AutoMigrate(&entities.Todo{})

	return nil
}
