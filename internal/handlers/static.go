package handlers

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"path/filepath"
)

func VideoHandler(c *gin.Context) {
	filename := c.Param("filename")
	videoPath := filepath.Join("statics/videos", filename)

	if _, err := filepath.Abs(videoPath); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Video not found"})
		return
	}

	c.File(videoPath)
}

func ImageHandler(c *gin.Context) {
	filename := c.Param("filename")
	imagePath := filepath.Join("statics/images", filename)

	if _, err := filepath.Abs(imagePath); err != nil {
		c.JSON(http.StatusNotFound, gin.H{"error": "Image not found"})
		return
	}

	c.File(imagePath)
}
