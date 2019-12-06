package models

import "fmt"

// IssueFieldsEntity contains needed issue fields
type IssueFieldsEntity struct {
	Assignee UserReferenceEntity `json:"assignee"`
	Creator  UserReferenceEntity `json:"creator"`
	Reporter UserReferenceEntity `json:"reporter"`

	Issuetype IssueTypeEntity `json:"issuetype"`
	Status    StatusEntity    `json:"status"`
	Priority  *IssuePriority  `json:"priority"`

	Summary   string `json:"summary"`
	Remaining int    `json:"timeestimate"`
	Estimate  int    `json:"timeoriginalestimate"`
	TimeSpend int    `json:"timespent"`
}

func (model IssueFieldsEntity) FormatedRemaining() string {
	return formateTime(model.Remaining)
}

func (model IssueFieldsEntity) FormatedEstimate() string {
	return formateTime(model.Estimate)
}

func (model IssueFieldsEntity) FormatedTimeSpend() string {
	return formateTime(model.TimeSpend)
}

func formateTime(seconds int) string {
	var minutes = seconds / 60

	if minutes >= 60 {
		hours := float32(minutes) / 60.0
		return fmt.Sprintf("%.2f ч.", hours)
	}
	return fmt.Sprintf("%d мин.", minutes)
}
