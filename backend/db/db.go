package db

import (
	"fmt"

	"gorm.io/driver/mysql"
	"gorm.io/gorm"
	"vue-go-users.com/config"
)

var db *gorm.DB

// Init initializes the DB connection
func Init() {
	c := config.GetConfig()
	dsn := fmt.Sprintf(
		"%s:%s@tcp(%s:%s)/%s?charset=utf8mb4&parseTime=true&loc=%s",
		c.GetString("DB_USERNAME"),
		c.GetString("DB_PASSWORD"),
		c.GetString("DB_HOST"),
		c.GetString("DB_PORT"),
		c.GetString("DB_DATABASE"),
		c.GetString("DB_TIMEZONE"),
	)

	var err error
	db, err = gorm.Open(mysql.Open(dsn), &gorm.Config{})
	if err != nil {
		panic(err)
	}
}

// GetDB gets the DB connection
func GetDB() *gorm.DB {
	return db
}

// Close closes the DB connection
func Close() {
	sqlDB, _ := db.DB()
	sqlDB.Close()
}
