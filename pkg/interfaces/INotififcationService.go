package interfaces

//go:generate mockery --name INotificationService
type INotificationService interface {
	SendUnreadBookmarkNotification() error
}
