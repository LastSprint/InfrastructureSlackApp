package jira

import models "github.com/LastSprint/InfrastructureSlackApp/models/jira"

// Filter фильтрует тикеты по определенному предикату.
func Filter(issues []models.IssueEntity, predicate func(models.IssueEntity) bool) []models.IssueEntity {

	result := []models.IssueEntity{}

	for _, item := range issues {
		if predicate(item) {
			result = append(result, item)
		}
	}

	return result
}

// GetAllUnestimated возвращает все тикеты, у которых нет оценки (или она 0)
func GetAllUnestimated(issues []models.IssueEntity) []models.IssueEntity {
	return Filter(issues, func(model models.IssueEntity) bool {
		return model.Fields.Estimate == 0 && model.Fields.Status.IsToDo()
	})
}

// GetAllZeroRemaining возвращает все тикеты, которые
func GetAllZeroRemaining(issues []models.IssueEntity) []models.IssueEntity {
	return Filter(issues, func(model models.IssueEntity) bool {
		return model.Fields.Remaining == 0 && model.Fields.Status.ID == models.InProgressID
	})
}

// GetAllBlocked возвращает все заблокированные тикеты.
func GetAllBlocked(issues []models.IssueEntity) []models.IssueEntity {
	return Filter(issues, func(model models.IssueEntity) bool {
		return model.Fields.Status.ID == models.BlockedID
	})
}
