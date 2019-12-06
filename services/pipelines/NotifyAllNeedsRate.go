package pipelines

import (
	"github.com/LastSprint/InfrastructureSlackApp/repositories"
	"github.com/LastSprint/InfrastructureSlackApp/utils"
	"github.com/sirupsen/logrus"
)

// NotifyAllNeedsRate пайплайн для отправки уведомления о неоценненых задачах для всех пользователей.
type NotifyAllNeedsRate struct {
	// Репозиторий для чтения пользователей.
	Repo repositories.UserRepository
}

// InitPipeline считывает всех пользователей и отправляет всем сообщения.
func (pipeline *NotifyAllNeedsRate) InitPipeline() (bool, error) {
	users, err := pipeline.Repo.ReadAllDevelopers()

	if err != nil {


		utils.Loger.WithFields(logrus.Fields{
			"pipeline": "NotifyAllNeedsRate",
			"isSended": false,
			"error": err,
			"reason": 0,
		}).Info("ANALYTICS_SYSTEM")

		return false, err
	}

	if len(users) == 0 {

		utils.Loger.WithFields(logrus.Fields{
			"pipeline": "NotifyAllNeedsRate",
			"isSended": false,
			"reason": 1,
		}).Info("ANALYTICS_SYSTEM")

		return false, nil
	}

	notUserPipe := ShowIssueToRatePipeline{User: nil}

	for _, user := range users {
		notUserPipe.User = user
		notUserPipe.InitPipeline()
	}

	return true, nil
}
