package controllers

import "github.com/gin-gonic/gin"

// HealthCheck tests if the server is healthy
func HealthCheck(c *gin.Context) {
	c.String(200, "OK")
}
