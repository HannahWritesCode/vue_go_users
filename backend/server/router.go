package server

import (
	"time"

	sentrygin "github.com/getsentry/sentry-go/gin"
	"github.com/gin-contrib/cors"
	"github.com/gin-gonic/gin"

	"vue-go-users.com/controllers/auth"
	"vue-go-users.com/controllers/users"

	"vue-go-users.com/controllers"
	"vue-go-users.com/middleware"
)

// NewRouter creates a new router instance
func NewRouter() *gin.Engine {
	//conf := config.GetConfig()
	router := gin.Default()

	// Add sentry
	router.Use(sentrygin.New(sentrygin.Options{
		Repanic: true,
	}))

	// CORS
	router.Use(cors.New(cors.Config{
		AllowMethods: []string{"GET", "POST", "OPTIONS", "DELETE", "PUT"},
		AllowHeaders: []string{
			"Origin", "Content-Length", "Authorization", "Content-Type",
			"User-Agent", "Referrer", "Host", "Token",
		},
		ExposeHeaders:    []string{"Content-Length"},
		AllowCredentials: true,
		AllowAllOrigins:  true,
		MaxAge:           time.Hour * 12,
	}))

	router.GET("/health-check", controllers.HealthCheck)

	v1 := router.Group("api/v1")
	{
		v1.POST("/users", users.Create)
		v1.POST("/auth", auth.Login)

		authGroup := v1.Group("", middleware.Middleware())
		{
			authGroup.POST("/reset-password", auth.ResetPassword)

			authGroup.GET("/users", users.Get)
			authGroup.PUT("/users", users.Update)
			authGroup.DELETE("/users", users.Delete)
		}

		v1.POST("/request-password", auth.SendResetPasswordEmail)
	}

	return router
}
