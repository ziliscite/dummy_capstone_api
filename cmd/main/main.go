package main

import (
	"dummy_capstone/internal/handlers"
	"github.com/gin-gonic/gin"
	"github.com/joho/godotenv"
	"log"
)

func main() {
	err := godotenv.Load()
	if err != nil {
		log.Fatalf("Error loading .env file")
	}

	r := gin.Default()

	r.POST("/upload", handlers.UploadHandler)
	r.GET("/histories", handlers.GetHistory)

	r.GET("/video/:filename", handlers.VideoHandler)
	r.GET("/image/:filename", handlers.ImageHandler)

	r.GET("/notification", handlers.GetAllNotification)
	r.PATCH("/notification/:id", handlers.UpdateNotificationStatus)

	r.GET("/summaries/weekly", handlers.GetWeeklySummaries)
	r.GET("/summaries/monthly", handlers.GetAnnualSummaries)

	if err := r.Run(":8082"); err != nil {
		log.Fatal(err)
	}
}
