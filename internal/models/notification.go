package models

import (
	"github.com/joho/godotenv"
	"os"
)

type Notification struct {
	ID        string `json:"id"`
	MediaURL  string `json:"media_url"`
	CreatedAt string `json:"created_at"`
	Category  string `json:"category"`
	IsRead    bool   `json:"is_read"`
}

var _ = godotenv.Load()
var BaseUrl = os.Getenv("BASE_URL")

var NotificationData = []*Notification{
	{ID: "1", MediaURL: BaseUrl + "/video/video_2.mkv", CreatedAt: "2024-11-20T08:00:00Z", Category: "Dead Chicken", IsRead: false},
	{ID: "2", MediaURL: BaseUrl + "/video/video_5.mp4", CreatedAt: "2024-11-20T09:15:00Z", Category: "High Crowd Density", IsRead: true},
	{ID: "3", MediaURL: BaseUrl + "/video/video_4.mkv", CreatedAt: "2024-11-20T11:45:00Z", Category: "Suspicious Activity", IsRead: false},
	{ID: "4", MediaURL: BaseUrl + "/video/video_1.mp4", CreatedAt: "2024-11-20T12:30:00Z", Category: "Fire Alert", IsRead: true},
	{ID: "5", MediaURL: BaseUrl + "/video/video_2.mkv", CreatedAt: "2024-11-20T13:00:00Z", Category: "Intrusion Detected", IsRead: false},
	{ID: "6", MediaURL: BaseUrl + "/video/video_3.mp4", CreatedAt: "2024-11-20T15:10:00Z", Category: "Dead Chicken", IsRead: true},
	{ID: "7", MediaURL: BaseUrl + "/video/video_4.mkv", CreatedAt: "2024-11-20T16:45:00Z", Category: "High Crowd Density", IsRead: false},
	{ID: "8", MediaURL: BaseUrl + "/video/video_1.mp4", CreatedAt: "2024-11-21T08:20:00Z", Category: "Suspicious Activity", IsRead: true},
	{ID: "9", MediaURL: BaseUrl + "/video/video_2.mkv", CreatedAt: "2024-11-21T09:30:00Z", Category: "Fire Alert", IsRead: false},
	{ID: "10", MediaURL: BaseUrl + "/video/video_3.mp4", CreatedAt: "2024-11-21T11:50:00Z", Category: "Intrusion Detected", IsRead: true},
	{ID: "11", MediaURL: BaseUrl + "/video/video_4.mkv", CreatedAt: "2024-11-21T14:00:00Z", Category: "Dead Chicken", IsRead: false},
	{ID: "12", MediaURL: BaseUrl + "/video/video_5.mp4", CreatedAt: "2024-11-21T16:15:00Z", Category: "High Crowd Density", IsRead: true},
	{ID: "13", MediaURL: BaseUrl + "/video/video_2.mkv", CreatedAt: "2024-11-21T17:45:00Z", Category: "Suspicious Activity", IsRead: false},
	{ID: "14", MediaURL: BaseUrl + "/video/video_3.mp4", CreatedAt: "2024-11-22T07:10:00Z", Category: "Fire Alert", IsRead: true},
	{ID: "15", MediaURL: BaseUrl + "/video/video_4.mkv", CreatedAt: "2024-11-22T09:25:00Z", Category: "Intrusion Detected", IsRead: false},
	{ID: "16", MediaURL: BaseUrl + "/video/video_1.mp4", CreatedAt: "2024-11-22T11:40:00Z", Category: "Dead Chicken", IsRead: true},
	{ID: "17", MediaURL: BaseUrl + "/video/video_2.mkv", CreatedAt: "2024-11-22T12:55:00Z", Category: "High Crowd Density", IsRead: false},
	{ID: "18", MediaURL: BaseUrl + "/video/video_5.mp4", CreatedAt: "2024-11-22T15:00:00Z", Category: "Suspicious Activity", IsRead: true},
	{ID: "19", MediaURL: BaseUrl + "/video/video_4.mkv", CreatedAt: "2024-11-22T16:20:00Z", Category: "Fire Alert", IsRead: false},
	{ID: "20", MediaURL: BaseUrl + "/video/video_1.mp4", CreatedAt: "2024-11-23T08:15:00Z", Category: "Intrusion Detected", IsRead: true},
	{ID: "21", MediaURL: BaseUrl + "/video/video_2.mkv", CreatedAt: "2024-11-23T10:30:00Z", Category: "Dead Chicken", IsRead: false},
	{ID: "22", MediaURL: BaseUrl + "/video/video_3.mp4", CreatedAt: "2024-11-23T12:45:00Z", Category: "High Crowd Density", IsRead: true},
	{ID: "23", MediaURL: BaseUrl + "/video/video_4.mkv", CreatedAt: "2024-11-23T14:50:00Z", Category: "Suspicious Activity", IsRead: false},
	{ID: "24", MediaURL: BaseUrl + "/video/video_1.mp4", CreatedAt: "2024-11-23T16:00:00Z", Category: "Fire Alert", IsRead: true},
	{ID: "25", MediaURL: BaseUrl + "/video/video_2.mkv", CreatedAt: "2024-11-23T17:10:00Z", Category: "Intrusion Detected", IsRead: false},
	{ID: "26", MediaURL: BaseUrl + "/video/video_5.mp4", CreatedAt: "2024-11-24T09:00:00Z", Category: "Dead Chicken", IsRead: true},
	{ID: "27", MediaURL: BaseUrl + "/video/video_4.mkv", CreatedAt: "2024-11-24T10:20:00Z", Category: "High Crowd Density", IsRead: false},
	{ID: "28", MediaURL: BaseUrl + "/video/video_1.mp4", CreatedAt: "2024-11-24T12:35:00Z", Category: "Suspicious Activity", IsRead: true},
	{ID: "29", MediaURL: BaseUrl + "/video/video_2.mkv", CreatedAt: "2024-11-24T14:45:00Z", Category: "Fire Alert", IsRead: false},
	{ID: "30", MediaURL: BaseUrl + "/video/video_3.mp4", CreatedAt: "2024-11-24T16:55:00Z", Category: "Intrusion Detected", IsRead: true},
}
