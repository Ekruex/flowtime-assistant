package controllers

import (
	"flowtime-backend/models"
	"flowtime-backend/utils"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
)

var sessions []models.Session // In-memory storage for simplicity

func StartSession(c *gin.Context) {
	session := models.Session{
		ID:        time.Now().Format("20060102150405"), // Unique ID based on timestamp
		StartTime: time.Now(),
	}
	sessions = append(sessions, session)
	utils.JSONResponse(c, http.StatusOK, gin.H{"message": "Session started", "session": session})
}

func EndSession(c *gin.Context) {
	var request struct {
		ID string `json:"id"`
	}
	if err := c.BindJSON(&request); err != nil {
		utils.JSONResponse(c, http.StatusBadRequest, gin.H{"error": "Invalid request"})
		return
	}

	// Find the session by ID
	for i, session := range sessions {
		if session.ID == request.ID {
			sessions[i].EndTime = time.Now()
			sessions[i].Duration = int(sessions[i].EndTime.Sub(sessions[i].StartTime).Seconds())
			utils.JSONResponse(c, http.StatusOK, gin.H{"message": "Session ended", "session": sessions[i]})
			return
		}
	}

	utils.JSONResponse(c, http.StatusNotFound, gin.H{"error": "Session not found"})
}
