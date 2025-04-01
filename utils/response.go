package utils

import "github.com/gin-gonic/gin"

func JSONResponse(c *gin.Context, status int, payload gin.H) {
	c.JSON(status, payload)
}
