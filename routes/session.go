package routes

import (
	"net/http"

	"github.com/gin-gonic/gin"
)

func StartSession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Session started successfully!",
	})
}

func EndSession(c *gin.Context) {
	c.JSON(http.StatusOK, gin.H{
		"message": "Session ended successfully!",
	})
}
