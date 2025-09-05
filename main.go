package main

import (
	"swaggo-poc/handlers"

	"github.com/gin-gonic/gin"
	swaggerFiles "github.com/swaggo/files"
	ginSwagger "github.com/swaggo/gin-swagger"

	_ "swaggo-poc/docs" // This line is necessary for go-swagger to find your docs!
)

// @title Swaggo POC API with Nested Structures
// @version 1.0
// @description This is a sample API server using Swaggo for Swagger documentation generation with nested JSON structures for user management.
// @termsOfService http://swagger.io/terms/

// @contact.name API Support
// @contact.url http://www.swagger.io/support
// @contact.email support@swagger.io

// @license.name Apache 2.0
// @license.url http://www.apache.org/licenses/LICENSE-2.0.html

// @host localhost:8080
// @BasePath /api/v1
// @schemes http
func main() {
	r := gin.Default()

	// API v1 routes
	v1 := r.Group("/api/v1")
	{
		// User routes
		v1.GET("/users", handlers.GetUsers)
		v1.GET("/users/:id", handlers.GetUser)
		v1.POST("/users", handlers.CreateUser)
		v1.PUT("/users/:id", handlers.UpdateUser)
		v1.DELETE("/users/:id", handlers.DeleteUser)

		// Nested user data routes
		v1.GET("/users/:id/address", handlers.GetUserAddress)
		v1.PUT("/users/:id/address", handlers.UpdateUserAddress)
		v1.GET("/users/:id/preferences", handlers.GetUserPreferences)
		v1.PUT("/users/:id/preferences", handlers.UpdateUserPreferences)
		v1.GET("/users/:id/profile", handlers.GetUserProfile)
		v1.PUT("/users/:id/profile", handlers.UpdateUserProfile)

		// User action endpoints (routes based on 'action' query param)
		v1.GET("/user-action", handlers.UserActionHandler)    // ?action=list or ?action=get&id=1
		v1.POST("/user-action", handlers.UserActionHandler)   // ?action=create
		v1.PUT("/user-action", handlers.UserActionHandler)    // ?action=update&id=1
		v1.DELETE("/user-action", handlers.UserActionHandler) // ?action=delete&id=1

		// Health check
		v1.GET("/health", handlers.HealthCheck)
	}

	// Swagger documentation endpoint
	r.GET("/swagger/*any", ginSwagger.WrapHandler(swaggerFiles.Handler))

	// Start server on port 8080
	r.Run(":8080")
}
