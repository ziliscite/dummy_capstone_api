package handlers

import (
	"dummy_capstone/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
)

func GetAllNotification(c *gin.Context) {
	media := services.GetAllNotification()
	c.JSON(http.StatusOK, gin.H{"data": media})
}

type UpdateStatusRequest struct {
	IsRead bool `json:"is_read"`
}

func UpdateNotificationStatus(c *gin.Context) {
	id := c.Param("id")
	var request UpdateStatusRequest
	if err := c.ShouldBindJSON(&request); err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid request body"})
		return
	}

	updatedNotification, err := services.UpdateNotificationStatus(id, request.IsRead)
	if err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Media not found"})
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": updatedNotification})
}
