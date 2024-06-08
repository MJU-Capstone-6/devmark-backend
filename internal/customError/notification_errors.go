package customerror

import "net/http"

var NOTIFICATION_HISTORY_NOT_FOUND = NewError(http.StatusNotFound, "Notification History Not Found.", "1301")

func NotificationHistoryNotFound(err error) error {
	notificationHistoryNotFound := NOTIFICATION_HISTORY_NOT_FOUND
	notificationHistoryNotFound.Data = err
	return notificationHistoryNotFound
}
