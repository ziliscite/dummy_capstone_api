package services

import "dummy_capstone/internal/models"

func GetAnnualSummaryData(year int) []models.MonthlySummary {
	if year == 2023 {
		return models.MonthlySummaryData2023
	}

	return models.MonthlySummaryData2024
}

func GetWeeklySummaryData(year, month, week int) []models.DailySummary {
	if week == 3 {
		return models.DailySummaryData3
	}

	return models.DailySummaryData4
}
