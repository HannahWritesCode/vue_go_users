package controllers

import (
	"net/http"
	"net/http/httptest"
	"testing"

	"github.com/gin-gonic/gin"
	"github.com/stretchr/testify/assert"
)

// Test HealthCheck
func TestHealthController_HealthCheck(t *testing.T) {
	testRecorder := httptest.NewRecorder()
	testContext, _ := gin.CreateTestContext(testRecorder)
	HealthCheck(testContext)

	assert.Equal(t, http.StatusOK, testContext.Writer.Status())
}
