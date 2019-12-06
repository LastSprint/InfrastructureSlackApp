package jira

import (
	"github.com/LastSprint/InfrastructureSlackApp/utils"
)

// SearchRequest Модель для запроса в Jira.
type SearchRequest struct {
	// IncludedStatuses статусы Jira-итемов, которые НУЖНО получить в ответ на этот запрос.
	// Константны описаны в пакете models/jira
	IncludedStatuses []string
	// ExcludedStatuses статусы Jira-итемов, которые НЕ НУЖНО получать в ответ на этот запрос.
	// Константны описаны в пакете models/jira
	ExcludedStatuses []string
	// IncludedTypes типы Jira-итемов, которые НУЖНО получить в ответ на этот запрос.
	// Константны описаны в пакете models/jira
	IncludedTypes []string
	// ExcludedTypes типы Jira-итемов, которые НЕ НУЖНО получать в этот на этот запрос.
	// Константны описаны в пакете models/jira
	ExcludedTypes []string
	// Assignee тот, на кого назначен Jira-итем.
	Assignee string
	// Ordering опциональная сортировка.
	Ordering *OrderingModel
	// Priorities приортеты Jira-итемов, которые нужно включить в выдачу.
	Priorities []string
}

// MakeJiraRequest конвертирует структуру в строку JQL запроса.
func (req SearchRequest) MakeJiraRequest() string {

	result := []string{}

	if len(req.Assignee) != 0 {
		result = append(result, JiraFieldAssignee+" = "+req.Assignee)
	}

	if str := utils.JoinByCharacter(req.IncludedStatuses, ",", "\""); len(str) != 0 {
		result = append(result, JiraFieldStatus+" in ("+str+")")
	}

	if str := utils.JoinByCharacter(req.ExcludedStatuses, ",", "\""); len(str) != 0 {
		result = append(result, JiraFieldStatus+" not in ("+str+")")
	}

	if str := utils.JoinByCharacter(req.IncludedTypes, ",", "\""); len(str) != 0 {
		result = append(result, JiraFieldIssueType+" in ("+str+")")
	}

	if str := utils.JoinByCharacter(req.ExcludedTypes, ",", "\""); len(str) != 0 {
		result = append(result, JiraFieldIssueType+" not in ("+str+")")
	}

	if str := utils.JoinByCharacter(req.Priorities, ",", "\""); len(str) != 0 {
		result = append(result, JiraFieldPriority+" in ("+str+")")
	}

	str := utils.JoinByCharacter(result, " and ", "")

	if order := req.Ordering; order != nil {
		str += " " + order.MakeJiraRequest()
	}

	return str
}
