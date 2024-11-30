package handlers

import (
	"dummy_capstone/internal/services"
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
)

func GetAnnualSummaries(c *gin.Context) {
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Year must be a valid integer"})
		return
	}

	data := services.GetAnnualSummaryData(year)
	c.JSON(http.StatusOK, gin.H{"data": data})
}

func GetWeeklySummaries(c *gin.Context) {
	year, err := strconv.Atoi(c.Query("year"))
	if err != nil {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid year parameter"})
		return
	}

	month, err := strconv.Atoi(c.Query("month"))
	if err != nil || month < 1 || month > 12 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid month parameter"})
		return
	}

	week, err := strconv.Atoi(c.Query("week"))
	if err != nil || week < 1 || week > 5 {
		c.JSON(http.StatusBadRequest, gin.H{"error": "Invalid week parameter"})
		return
	}

	data := services.GetWeeklySummaryData(year, month, week)
	c.JSON(http.StatusOK, gin.H{"data": data})
}
