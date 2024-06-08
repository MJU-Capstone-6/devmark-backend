package scheduler

import (
	"log"
	"time"

	"github.com/MJU-Capstone-6/devmark-backend/pkg/interfaces"
	"github.com/go-co-op/gocron/v2"
)

type UnreadBookmarkSchedulerService struct {
	NotificationService interfaces.INotificationService
}

func (u *UnreadBookmarkSchedulerService) Run() (gocron.Scheduler, error) {
	unreadBookmarkScheduler, err := gocron.NewScheduler()
	if err != nil {
		return unreadBookmarkScheduler, err
	}

	job, err := unreadBookmarkScheduler.NewJob(gocron.DurationJob(time.Hour),
		gocron.NewTask(func() {
			err := u.NotificationService.SendUnreadBookmarkNotification()
			if err != nil {
				log.Println(err)
			}
		}))
	log.Println(job.ID())
	if err != nil {
		return unreadBookmarkScheduler, err
	}
	unreadBookmarkScheduler.Start()
	return unreadBookmarkScheduler, nil
}

func InitUnreadBookmarkSchedulerService() *UnreadBookmarkSchedulerService {
	return &UnreadBookmarkSchedulerService{}
}

func (u UnreadBookmarkSchedulerService) WithNotificationService(service interfaces.INotificationService) UnreadBookmarkSchedulerService {
	u.NotificationService = service
	return u
}
