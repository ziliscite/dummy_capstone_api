package services

import (
	"dummy_capstone/internal/models"
	"mime/multipart"
	"os"
	"time"
)

func ProcessClassification(header *multipart.FileHeader) (*models.Classification, error) {
	var baseUrl = os.Getenv("BASE_URL")

	var contentUrl string
	if header.Header.Get("Content-Type") == "image/*" || header.Header.Get("Content-Type") == "image/jpg" {
		contentUrl = baseUrl + "/image/" + header.Filename
	} else {
		contentUrl = baseUrl + "/video/" + header.Filename
	}

	currentTime := time.Now().UTC()
	formattedTime := currentTime.Format("2006-01-02T15:04:05Z")

	data := &models.Classification{
		MediaURL:     contentUrl,
		DeadCount:    5,
		ChickenCount: int(header.Size),
		CrowdDensity: header.Header.Get("Content-Type"),
		CreatedAt:    formattedTime,
		IsAlert:      false,
	}

	return data, nil
}

func GetClassificationHistory() []models.Classification {
	return models.ClassificationHistory
}
