package pipelines

import (
	log "github.com/LastSprint/InfrastructureSlackApp/logging"
	"github.com/LastSprint/InfrastructureSlackApp/repositories"
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
		log.PipelineByName(log.NotifyAllNeedsRate, err, false, log.DataReading, nil)
		return false, err
	}

	if len(users) == 0 {
		log.PipelineByName(log.NotifyAllNeedsRate, err, false, log.ContentIsEmpty, nil)
		return false, nil
	}

	notUserPipe := ShowIssueToRatePipeline{User: nil}

	for _, user := range users {
		notUserPipe.User = user
		notUserPipe.InitPipeline()
	}

	return true, nil
}
