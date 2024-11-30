package models

type MonthlySummary struct {
	Month     string `json:"month"`
	DeadCount int    `json:"dead_count"`
}

var MonthlySummaryData2023 = []MonthlySummary{
	{"january", 10},
	{"february", 12},
	{"march", 15},
	{"april", 11},
	{"may", 14},
	{"june", 13},
	{"july", 17},
	{"august", 16},
	{"september", 18},
	{"october", 20},
	{"november", 19},
	{"december", 22},
}

var MonthlySummaryData2024 = []MonthlySummary{
	{"january", 15},
	{"february", 16},
	{"march", 14},
	{"april", 10},
	{"may", 10},
	{"june", 13},
	{"july", 21},
	{"august", 13},
	{"september", 21},
	{"october", 19},
	{"november", 20},
	{"december", 26},
}

type DailySummary struct {
	Day          string `json:"day"`
	DeadCount    int    `json:"dead_count"`
	CrowdDensity string `json:"crowd_density"`
	ChickenCount int    `json:"chicken_count"`
	IsAlert      bool   `json:"is_alert"`
}

var DailySummaryData3 = []DailySummary{
	{"monday", 7, "low", 101, true},
	{"tuesday", 6, "high", 105, true},
	{"wednesday", 8, "medium", 98, true},
	{"thursday", 5, "low", 102, true},
	{"friday", 6, "high", 100, true},
	{"saturday", 7, "medium", 97, true},
	{"sunday", 9, "medium", 99, true},
}

var DailySummaryData4 = []DailySummary{
	{"monday", 7, "low", 101, true},
	{"tuesday", 6, "high", 105, true},
	{"wednesday", 8, "medium", 98, true},
	{"thursday", 5, "low", 102, true},
	{"friday", 6, "high", 100, true},
	{"saturday", 7, "medium", 97, true},
	{"sunday", 9, "medium", 99, true},
}
