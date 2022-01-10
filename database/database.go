package database

import (
	"database/sql"
	"fmt"

	"main/models"

	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

var (
	DB   *gorm.DB
	CONN *sql.DB
)

func InitDatabase() {
	var err error
	DB, err = gorm.Open(sqlite.Open("blog.db"), &gorm.Config{})

	if err != nil {
		panic("Cannot connect to db")
	}

	fmt.Println("Connected to database")

	// create tables if not exists
	DB.AutoMigrate(&models.Comment{}, &models.User{}, &models.Session{})
	fmt.Println("Migrated to database")

}
