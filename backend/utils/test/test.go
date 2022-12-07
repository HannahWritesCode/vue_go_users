package test

import (
	"github.com/gin-gonic/gin"

	"vue-go-users.com/config"
	"vue-go-users.com/db"
)

// Setup setups testing env
func Setup() {
	config.Init("testing")
	gin.SetMode(gin.TestMode)

	// Setup and clean database
	db.Init()
}

// Teardown cleans up after the test
func Teardown() {
	db.Close()
}

// CleanTable does a quick truncate of the table provided
func CleanTable(table interface{}) {
	db.GetDB().Exec("SET foreign_key_checks=0")
	db.GetDB().Unscoped().Where("1 = 1").Delete(table)
	db.GetDB().Exec("SET foreign_key_checks=1")
}
