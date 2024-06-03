package notification

import (
	"context"
	"fmt"
	"log"

	"firebase.google.com/go/messaging"
	"github.com/MJU-Capstone-6/devmark-backend/internal/constants"
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
			resp, err := n.MessagingClient.SendMulticast(context.Background(), &messaging.MulticastMessage{
				Notification: &messaging.Notification{
					Title: fmt.Sprintf(constants.UNREAD_BOOKMARK_NOTIFIACTION_TITLE, *unreadBookmark.WorkspaceName, len(unreadBookmark.Bookmarks)),
					Body:  fmt.Sprintf(constants.UNREAD_BOOKMARK_NOTIFIACTION_BODY, *bookmark.Title),
				},
				Tokens: tokens,
			})
			if err != nil {
				log.Println(err)
			}
			log.Println(resp)
		}
	}
	return nil
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
