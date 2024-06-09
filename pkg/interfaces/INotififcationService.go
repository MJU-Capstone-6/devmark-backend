package interfaces

import "github.com/MJU-Capstone-6/devmark-backend/internal/repository"

//go:generate mockery --name INotificationService
type INotificationService interface {
	SendUnreadBookmarkNotification() error
	CreateNotificationHistory(repository.CreateNotificationHistoryParams) (*repository.NotificationHistory, error)
	FindUnreadNotificationHistory(int64) (*[]repository.FindUnreadNotificationHistoryRow, error)
}
