package pipelines

import (
	models "github.com/LastSprint/InfrastructureSlackApp/models"
	mj "github.com/LastSprint/InfrastructureSlackApp/models/jira"
	m_s "github.com/LastSprint/InfrastructureSlackApp/models/slack"
	"github.com/LastSprint/InfrastructureSlackApp/services/jira"
	"github.com/LastSprint/InfrastructureSlackApp/services/slack"
	"github.com/LastSprint/InfrastructureSlackApp/utils"

	log "github.com/LastSprint/InfrastructureSlackApp/logging"
)

// ShowIssueToRatePipeline пайплайн для отправки сообщения о задачах, которые нужно оценить.
type ShowIssueToRatePipeline struct {
	// Идентификатор пользователя в slack.
	User *models.User
}

var statuses = []string{
	mj.ToDo,
	mj.Reopened,
	mj.InProgress,
}

// InitPipeline отсылает сообщение в Slack.
// Сообщение не отправится, если ни одной ошибки не было найдено.
// Returns:
//	- bool: Указывает, было ли отправлено сообщение или нет.
//	- error: Был ли ошибка или нет.
func (pipeline *ShowIssueToRatePipeline) InitPipeline() (bool, error) {

	request := jira.SearchRequest{
		ProjectID:        pipeline.User.ProjectID,
		Assignee:         pipeline.User.JiraID,
		IncludedStatuses: statuses,
		ExcludedTypes:    []string{mj.IssueTypeServiceTask, mj.IssueTypeTestExecuted},
	}

	issues, err := jira.LoadIssues(request)

	if err != nil {
		log.PipelineByName(log.ShowIssueToRatePipeline, err, false, log.DataReading, pipeline.User)
		return false, err
	}

	if len(issues.Issues) == 0 {
		log.PipelineByName(log.ShowIssueToRatePipeline, err, false, log.ContentIsEmpty, pipeline.User)
		return false, nil
	}

	message := "*" + pipeline.User.FirstName + "*" + "\n"
	message += "Тебе нужно оценить эти задачи:\n"

	unemMsg, unMsgLen := createUnestimatedMessage(issues.Issues)

	zrMsg, zrMsgLen := createZeroRemainingMessage(issues.Issues)

	slackMsg := m_s.PostChatMessage{
		Text:       "*" + pipeline.User.FirstName + "*" + " :beb:" + "\n",
		IsMarkdown: true,
		Channel:    pipeline.User.SlackID,
	}

	if unemMsg == nil {
		if zrMsg == nil {
			log.PipelineByName(log.ShowIssueToRatePipeline, err, false, log.AnalyzedDataIsEmpty, pipeline.User)
			return false, nil
		}

		slackMsg.Text += *zrMsg
	}

	if unemMsg != nil {
		slackMsg.Text += *unemMsg

		if zrMsg != nil {
			slackMsg.Text += "\n" + *zrMsg
		}
	}

	err = slack.SendMessage(slackMsg)

	payload := map[string]interface{}{"unestimatedCount": unMsgLen, "zeroRemainingCount": zrMsgLen}

	log.PipelineByName(log.ShowIssueToRatePipeline, err, err == nil, log.Successful, payload)

	if err != nil {
		return false, err
	}

	return true, nil
}

func createUnestimatedMessage(issues []mj.IssueEntity) (*string, int) {
	filtred := jira.GetAllUnestimated(issues)

	if len(filtred) == 0 {
		return nil, 0
	}

	message := "Тебе нужно оценить эти задачи:\n"

	for _, issue := range filtred {
		message += utils.Config.JiraBaseHost + issue.Key + "\n"
	}

	return &message, len(filtred)
}

func createZeroRemainingMessage(issues []mj.IssueEntity) (*string, int) {
	filtred := jira.GetAllZeroRemaining(issues)

	if len(filtred) == 0 {
		return nil, 0
	}

	message := "*Оценку этих задачек нужно актуализировать:*\n"

	for _, issue := range filtred {
		message += utils.Config.JiraBaseHost + issue.Key + "\n"
	}

	return &message, len(filtred)
}
