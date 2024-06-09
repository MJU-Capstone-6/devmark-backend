package notification

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/messaging"
	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
	customerror "github.com/MJU-Capstone-6/devmark-backend/internal/customError"
	"github.com/MJU-Capstone-6/devmark-backend/internal/repository"
	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
)

type NotificationService struct {
	MessagingClient *messaging.Client
	Repository      interfaces.IRepository
}

func (n *NotificationService) SendUnreadBookmarkNotification() error {
	unreadBookmarks, err := n.Repository.FindUnreadBookmark(context.Background())
	if err != nil {
		return err
	}
	for _, unreadBookmark := range unreadBookmarks {
		tokens := []string{}
		for _, deviceInfo := range unreadBookmark.DeviceInfos {
			tokens = append(tokens, *deviceInfo.RegistrationToken)
		}
		for _, bookmark := range unreadBookmark.Bookmarks {
			title := fmt.Sprintf(constants.UNREAD_BOOKMARK_NOTIFIACTION_TITLE, *unreadBookmark.WorkspaceName, len(unreadBookmark.Bookmarks))
			_, err := n.MessagingClient.SendMulticast(context.Background(), &messaging.MulticastMessage{
				Notification: &messaging.Notification{
					Title: title,
					Body:  fmt.Sprintf(constants.UNREAD_BOOKMARK_NOTIFIACTION_BODY, *bookmark.Title),
				},
				Tokens: tokens,
			})
			if err != nil {
				log.Println(err)
			}
			param := repository.CreateNotificationHistoryParams{
				UserID:            bookmark.UserID,
				NotificationTitle: &title,
				BookmarkID:        &bookmark.ID,
			}
			_, err = n.CreateNotificationHistory(param)
			if err != nil {
				log.Println(err)
			}
		}
	}
	return nil
}

func (n *NotificationService) CreateNotificationHistory(param repository.CreateNotificationHistoryParams) (*repository.NotificationHistory, error) {
	history, err := n.Repository.CreateNotificationHistory(context.Background(), param)
	if err != nil {
		return nil, err
	}
	return &history, nil
}

func (n *NotificationService) FindUnreadNotificationHistory(userID int64) (*repository.UnreadNotification, error) {
	historys, err := n.Repository.FindUnreadNotificationHistory(context.Background(), userID)
	if err != nil {
		log.Println(err)
		return nil, customerror.NotificationHistoryNotFound(err)
	}
	return &historys, nil
}

func InitNotificationService() *NotificationService {
	return &NotificationService{}
}

func (n NotificationService) WithClient(client *messaging.Client) NotificationService {
	n.MessagingClient = client
	return n
}

func (n NotificationService) WithRepository(repo interfaces.IRepository) NotificationService {
	n.Repository = repo
	return n
}
