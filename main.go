package main

import (
	"flowtime-backend/routes"

	"os"

	"github.com/gin-gonic/gin"
)

func main() {
	// Set Gin mode
	gin.SetMode(os.Getenv("GIN_MODE"))

	r := gin.Default()

	// Health check route
	r.GET("/health", routes.HealthCheck)

	// Session routes
	r.POST("/session/start", routes.StartSession)
	r.POST("/session/end", routes.EndSession)

	// Start the server
	port := os.Getenv("PORT")
	if port == "" {
		port = "8080" // Default port
	}
	r.Run(":" + port)

	r.SetTrustedProxies(nil) // Set trusted proxies to nil for local development
}
