package controllers

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"vue-go-users.com/config"
	"vue-go-users.com/db"
)

var ()

func TestMain(m *testing.M) {
	testSetup()

	code := m.Run()

	testTeardown()

	os.Exit(code)
}

func testSetup() {
	gin.SetMode(gin.TestMode)

	// Setup database
	config.Init("testing")
	db.Init()
}

func testTeardown() {
	db.Close()
}
