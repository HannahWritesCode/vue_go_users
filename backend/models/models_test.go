package models

import (
	"os"
	"testing"

	"github.com/gin-gonic/gin"
	"vue-go-users.com/config"
	"vue-go-users.com/db"

	testutils "vue-go-users.com/utils/test"
)

var ()

func ptrString(s string) *string {
	return &s
}

func TestMain(m *testing.M) {
	testutils.Setup()

	code := m.Run()

	testutils.Teardown()

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
