package services

import (
	"dummy_capstone/internal/models"
	"errors"
	"sync"
)

var (
	notificationDataMutex sync.Mutex
	notificationData      []*models.Notification
)

func init() {
	notificationData = models.NotificationData
}

func GetAllNotification() []*models.Notification {
	notificationDataMutex.Lock()
	defer notificationDataMutex.Unlock()
	return notificationData
}

func UpdateNotificationStatus(id string, isRead bool) (*models.Notification, error) {
	notificationDataMutex.Lock()
	defer notificationDataMutex.Unlock()
	for i := range notificationData {
		if notificationData[i].ID == id {
			notificationData[i].IsRead = isRead
			return notificationData[i], nil
		}
	}
	return nil, errors.New("media not found")
}
