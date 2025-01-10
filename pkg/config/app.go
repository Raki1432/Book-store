package config

import (
	"github.com/jinzhu/gorm"
	_ "github.com/jinzhu/gorm/dialects/mysql"
)

var (
	db *gorm.DB
)

// Connect initializes the database connection
func Connect() {
	// Define the connection string
	d, err := gorm.Open("mysql", "root:raki1432@/testdb?charset=utf8&parseTime=True&loc=Local")
	if err != nil {
		panic(err)
	}
	db = d
}

// GetDB returns the database instance
func GetDB() *gorm.DB {
	return db
}
