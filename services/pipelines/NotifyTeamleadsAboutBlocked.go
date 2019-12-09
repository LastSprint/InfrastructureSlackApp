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

// NotifyTeamleadsAboutBlocked этот пайплайн оповещает тимлидов о заблокированных задачах.
type NotifyTeamleadsAboutBlocked struct {
	Repo repositories.UserRepository
}

// InitPipeline иниципллизирует пайплайн.
func (pipeline *NotifyTeamleadsAboutBlocked) InitPipeline() (bool, error) {
	leads, err := pipeline.Repo.ReadLeadDevelopers()

	if err != nil {
		log.PipelineByName(log.NotifyTeamleadsAboutBlocked, err, false, log.DataReading, nil)
		return false, err
	}

	for _, user := range leads {

		request := jira.SearchRequest{
			Assignee:         user.JiraID,
			IncludedStatuses: []string{models.Blocked},
			ExcludedTypes:    []string{models.IssueTypeServiceTask, models.IssueTypeTestExecuted},
		}

		issues, err := jira.LoadIssues(request)

		if err != nil {
			log.PipelineByName(log.NotifyTeamleadsAboutBlocked, err, false, log.DataReading, user)
			continue
		}

		if len(issues.Issues) == 0 {
			log.PipelineByName(log.NotifyTeamleadsAboutBlocked, err, false, log.ContentIsEmpty, user)
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

		log.PipelineByName(log.NotifyTeamleadsAboutBlocked, err, err == nil, log.Successful, logPayload)
	}
	return true, err
}
