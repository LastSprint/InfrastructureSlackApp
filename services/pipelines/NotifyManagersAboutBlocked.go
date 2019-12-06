package pipelines

import (
	models "github.com/LastSprint/InfrastructureSlackApp/models/jira"
	"github.com/LastSprint/InfrastructureSlackApp/models/slack"
	"github.com/LastSprint/InfrastructureSlackApp/repositories"
	"github.com/LastSprint/InfrastructureSlackApp/services/jira"
	"github.com/LastSprint/InfrastructureSlackApp/utils"
	"github.com/sirupsen/logrus"

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
		utils.Loger.WithFields(logrus.Fields{
			"pipeline": "NotifyManagersAboutBlocked",
			"isSended": false,
			"error":    err,
			"reason":   0,
		}).Info("ANALYTICS_SYSTEM")
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
			utils.Loger.WithFields(logrus.Fields{
				"pipeline": "NotifyManagersAboutBlocked",
				"isSended": false,
				"reason":   1,
				"user":     user,
			}).Info("ANALYTICS")
			continue
		}

		if len(issues.Issues) == 0 {
			utils.Loger.WithFields(logrus.Fields{
				"pipeline": "NotifyManagersAboutBlocked",
				"isSended": false,
				"reason":   2,
				"user":     user,
			}).Info("ANALYTICS")
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

		utils.Loger.WithFields(logrus.Fields{
			"user":        user,
			"pipeline":    "NotifyManagersAboutBlocked",
			"isSended":    err == nil,
			"Error":       err,
			"IssuesCount": len(issues.Issues),
		}).Info("ANALYTICS")
	}
	return true, err
}
