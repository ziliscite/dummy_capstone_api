package handlers

import (
	"dummy_capstone/internal/services"
	"io"
	"net/http"
	"os"

	"github.com/gin-gonic/gin"
)

func UploadHandler(c *gin.Context) {
	file, header, err := c.Request.FormFile("file")
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Failed to retrieve file"})
		return
	}
	defer file.Close()

	var path string
	if header.Header.Get("Content-Type") == "image/*" || header.Header.Get("Content-Type") == "image/jpg" {
		path = "images/"
	} else {
		path = "videos/"
	}

	out, err := os.Create("./statics/" + path + header.Filename)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Unable to save the file"})
		return
	}
	defer out.Close()

	_, err = io.Copy(out, file)
	if err != nil {
		c.JSON(http.StatusInternalServerError, gin.H{"error": "Failed to save the file"})
		return
	}

	fileInfo, err := services.ProcessClassification(header)
	if err != nil {
		c.JSON(http.StatusInternalServerError, err.Error())
		return
	}

	c.JSON(http.StatusOK, gin.H{"data": fileInfo})
}

func GetHistory(c *gin.Context) {
	data := services.GetClassificationHistory()
	c.JSON(http.StatusOK, gin.H{"data": data})
}
