package jira

const (
	// JiraFieldKey Ключ Jira-итема.
	JiraFieldKey string = "key"
	// JiraFieldAssignee тот, на кого назначен Jira-итем.
	JiraFieldAssignee string = "assignee"
	// JiraFieldCreator хз что это.
	JiraFieldCreator string = "creator"
	// JiraFieldReporter тот, кто создал Jira-итем.
	JiraFieldReporter string = "reporter"
	// JiraFieldIssueType тип Jira-итема.
	JiraFieldIssueType string = "issuetype"
	// JiraFieldStatus статус Jira-итема.
	JiraFieldStatus string = "status"
	// JiraFieldSummary Название Jira-итема.
	JiraFieldSummary string = "summary"
	// JiraFieldRemaining оценка Jira-итема.
	JiraFieldRemaining string = "timeestimate"
	// JiraFieldTimeoriginalestimate оригинальная (первая) оценка Jira-итема.
	JiraFieldTimeoriginalestimate string = "timeoriginalestimate"
	// JiraFieldTimespent затраченное время.
	JiraFieldTimespent string = "timespent"
	// JiraFieldPriority приоритет Jira-итема.
	JiraFieldPriority string = "priority"
)

var acceptedFields = []string{
	JiraFieldKey,
	JiraFieldAssignee,
	JiraFieldCreator,
	JiraFieldReporter,
	JiraFieldIssueType,
	JiraFieldStatus,
	JiraFieldSummary,
	JiraFieldRemaining,
	JiraFieldTimeoriginalestimate,
	JiraFieldTimespent,
	JiraFieldPriority,
}
