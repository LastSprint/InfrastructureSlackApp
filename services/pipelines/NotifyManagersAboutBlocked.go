package pipelines

import (
	log "github.com/LastSprint/InfrastructureSlackApp/logging"
	models "github.com/LastSprint/InfrastructureSlackApp/models/jira"
	"github.com/LastSprint/InfrastructureSlackApp/models/slack"
	"github.com/LastSprint/InfrastructureSlackApp/repositories"
	"github.com/LastSprint/InfrastructureSlackApp/services/jira"
	"github.com/LastSprint/InfrastructureSlackApp/utils"

	sl_serv "github.com/LastSprint/InfrastructureSlackApp/services/slack"
)

// NotifyManagersAboutBlocked этот пайплайн оповещает менеджеров о заблокированных задачах.
type NotifyManagersAboutBlocked struct {
	Repo repositories.UserRepository
}

// InitPipeline иниципллизирует пайплайн.
func (pipeline *NotifyManagersAboutBlocked) InitPipeline() (bool, error) {
	managers, err := pipeline.Repo.ReadManagers()

	if err != nil {
		log.PipelineByName(log.NotifyManagersAboutBlocked, err, false, log.DataReading, nil)
		return false, err
	}

	for _, user := range managers {

		request := jira.SearchRequest{
			Assignee:         user.JiraID,
			IncludedStatuses: []string{models.Blocked},
			ExcludedTypes:    []string{models.IssueTypeServiceTask},
		}

		issues, err := jira.LoadIssues(request)

		if err != nil {
			log.PipelineByName(log.NotifyManagersAboutBlocked, err, false, log.DataReading, user)
			continue
		}

		if len(issues.Issues) == 0 {
			log.PipelineByName(log.NotifyManagersAboutBlocked, nil, false, log.ContentIsEmpty, user)
			continue
		}

		text := "*" + user.FirstName + "*\n"

		for _, issue := range issues.Issues {
			text += utils.Config.JiraBaseHost + issue.Key + "\n"
		}

		message := slack.PostChatMessage{
			Text:       text,
			Channel:    user.SlackID,
			IsMarkdown: true,
		}

		err = sl_serv.SendMessage(message)

		logPayload := map[string]interface{}{"issueCount": len(issues.Issues), "user": user}

		log.PipelineByName(log.NotifyManagersAboutBlocked, err, err == nil, log.Successful, logPayload)
	}
	return true, err
}
