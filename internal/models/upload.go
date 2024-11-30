package models

type FileInfo struct {
	Filename string `json:"filename"`
	Size     int64  `json:"size"`
	Type     string `json:"type"`
}

type Classification struct {
	MediaURL     string `json:"media_url"`
	DeadCount    int    `json:"dead_count"`
	ChickenCount int    `json:"chicken_count"`
	CrowdDensity string `json:"crowd_density"`
	CreatedAt    string `json:"created_at"`
	IsAlert      bool   `json:"is_alert"`
}

var ClassificationHistory = []Classification{
	{
		MediaURL:     BaseUrl + "/video/video_1.mp4",
		DeadCount:    3,
		ChickenCount: 10,
		CrowdDensity: "High",
		CreatedAt:    "2024-11-21T08:20:00Z",
		IsAlert:      true,
	},
	{
		MediaURL:     BaseUrl + "/image/sticker_11.png",
		DeadCount:    0,
		ChickenCount: 20,
		CrowdDensity: "Low",
		CreatedAt:    "2024-11-21T09:30:00Z",
		IsAlert:      false,
	},
	{
		MediaURL:     BaseUrl + "/image/sticker_9.png",
		DeadCount:    5,
		ChickenCount: 15,
		CrowdDensity: "Medium",
		CreatedAt:    "2024-11-22T11:50:00Z",
		IsAlert:      true,
	},
	{
		MediaURL:     BaseUrl + "/video/video_4.mkv",
		DeadCount:    8,
		ChickenCount: 5,
		CrowdDensity: "High",
		CreatedAt:    "2024-11-22T14:00:00Z",
		IsAlert:      true,
	},
	{
		MediaURL:     BaseUrl + "/video/video_5.mp4",
		DeadCount:    2,
		ChickenCount: 30,
		CrowdDensity: "Low",
		CreatedAt:    "2024-11-22T15:50:00Z",
		IsAlert:      false,
	},
	{
		MediaURL:     BaseUrl + "/image/sticker_8.png",
		DeadCount:    0,
		ChickenCount: 25,
		CrowdDensity: "Low",
		CreatedAt:    "2024-11-23T17:00:00Z",
		IsAlert:      false,
	},
	{
		MediaURL:     BaseUrl + "/video/video_2.mkv",
		DeadCount:    10,
		ChickenCount: 0,
		CrowdDensity: "High",
		CreatedAt:    "2024-11-24T18:15:00Z",
		IsAlert:      true,
	},
	{
		MediaURL:     BaseUrl + "/video/video_3.mp4",
		DeadCount:    1,
		ChickenCount: 10,
		CrowdDensity: "Medium",
		CreatedAt:    "2024-11-25T19:30:00Z",
		IsAlert:      false,
	},
	{
		MediaURL:     BaseUrl + "/image/sticker_11.png",
		DeadCount:    3,
		ChickenCount: 7,
		CrowdDensity: "High",
		CreatedAt:    "2024-11-25T20:45:00Z",
		IsAlert:      true,
	},
	{
		MediaURL:     BaseUrl + "/image/sticker_12.png",
		DeadCount:    6,
		ChickenCount: 13,
		CrowdDensity: "Medium",
		CreatedAt:    "2024-11-26T22:00:00Z",
		IsAlert:      true,
	},
	{
		MediaURL:     BaseUrl + "/video/video_1.mp4",
		DeadCount:    0,
		ChickenCount: 30,
		CrowdDensity: "Low",
		CreatedAt:    "2024-11-26T23:10:00Z",
		IsAlert:      false,
	},
	{
		MediaURL:     BaseUrl + "/image/sticker_10.png",
		DeadCount:    2,
		ChickenCount: 12,
		CrowdDensity: "Medium",
		CreatedAt:    "2024-11-26T00:20:00Z",
		IsAlert:      false,
	},
}
