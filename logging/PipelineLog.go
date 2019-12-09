package logging

import "github.com/sirupsen/logrus"

// PipelineName имя пайплайна для лога
type PipelineName string

const (
	// NotifyAllNeedsRate см `NotifyAllNeedsRate.go`
	NotifyAllNeedsRate PipelineName = "NotifyAllNeedsRate"
	// ShowIssueToRatePipeline см `ShowIssueToRatePipeline.go`
	ShowIssueToRatePipeline PipelineName = "ShowIssueToRatePipeline"
	// NotifyTeamleadsAboutBlocked см `NotifyTeamleadsAboutBlocked.go`
	NotifyTeamleadsAboutBlocked PipelineName = "NotifyTeamleadsAboutBlocked"
	// NotifyManagersAboutBlocked см `NotifyManagersAboutBlocked.go`
	NotifyManagersAboutBlocked PipelineName = "NotifyManagersAboutBlocked"
	// FigmaVersionCheck см `FigmaVersionCheck.go`
	FigmaVersionCheck PipelineName = "FigmaVersionCheck"
)

// LogReason причина лога.
type LogReason string

const (
	// DataReading причина "Не удалось прочесть данные из репозитория/сети/источника данных "
	DataReading LogReason = "RepoReadingFailed"
	// ContentIsEmpty причина "Когда коллекция данных для анализа пуста"
	ContentIsEmpty LogReason = "ReadedContentIsEmpty"
	// CantCreateRequest причина "Не удалось составить запрос в сеть"
	CantCreateRequest LogReason = "CantCreateRequest"
	// FigmaHistoryNotChanged причина "История изменений файлов в Figma не изменялась"
	FigmaHistoryNotChanged LogReason = "FigmaHistoryNotChanged"
	// Successful причина "Действие выполнено успешно"
	Successful LogReason = "Successful"
	// AnalyzedDataIsEmpty причина "После анализа/фильтрования и прочих операций получились пустые данные"
	AnalyzedDataIsEmpty LogReason = "AnalyzedDataIsEmpty"
)

const analyticsConst = "ANALYTICS"

func PipelineByName(name PipelineName, err error, isSended bool, reason LogReason, payload interface{}) {
	Loger.WithFields(logrus.Fields{
		"pipeline": string(name),
		"payload":  payload,
		"isSended": isSended,
		"error":    err,
		"reason":   string(reason),
	}).Info(analyticsConst)
}
