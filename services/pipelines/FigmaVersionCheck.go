package pipelines

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	models "github.com/LastSprint/InfrastructureSlackApp/models/figma"
	slackModel "github.com/LastSprint/InfrastructureSlackApp/models/slack"
	"github.com/LastSprint/InfrastructureSlackApp/repositories"
	"github.com/LastSprint/InfrastructureSlackApp/services/slack"
	"github.com/LastSprint/InfrastructureSlackApp/utils"
	"github.com/sirupsen/logrus"
)

// FigmaVersionCheck пайплайн для отправки уведомления о изменения в Figma.
type FigmaVersionCheck struct {
	// Репозиторий для чтения пользователей.
	Repo repositories.FigmaFilesRepository
}

// InitPipeline запускает FigmaVersionCheck пайплайн
func (pipeline *FigmaVersionCheck) InitPipeline() (bool, error) {

	files, err := pipeline.Repo.ReadAllFiles()

	if err != nil {

		utils.Loger.WithFields(logrus.Fields{
			"pipeline": "FigmaVersionCheck",
			"isSended": false,
			"error":    err,
			"reason":   0,
		}).Info("ANALYTICS_SYSTEM")

		return false, err
	}

	if len(files) == 0 {
		utils.Loger.WithFields(logrus.Fields{
			"pipeline": "FigmaVersionCheck",
			"isSended": false,
			"reason":   1,
		}).Info("ANALYTICS_SYSTEM")
		return false, err
	}

	for _, file := range files {
		loadFigmaData(file, pipeline.Repo)
	}

	return true, nil
}

func loadFigmaData(file *models.FigmaProjectFileModel, rep repositories.FigmaFilesRepository) {
	request, err := http.NewRequest(
		"GET",
		"https://api.figma.com/v1/files/"+file.FileKey+"/versions",
		bytes.NewBuffer([]byte{}),
	)

	if err != nil {
		utils.Loger.WithFields(logrus.Fields{
			"pipeline":    "FigmaVersionCheck",
			"isSended":    false,
			"error":       err,
			"reason":      -1,
			"description": "Cant create request to figma",
		}).Info("ANALYTICS_SYSTEM")
		return
	}

	request.Header.Add("X-Figma-Token", utils.Config.FigmaToken)

	client := http.Client{}

	resp, err := client.Do(request)

	if err != nil {
		utils.Loger.WithFields(logrus.Fields{
			"pipeline":    "FigmaVersionCheck",
			"isSended":    false,
			"error":       err,
			"reason":      -1,
			"description": "request failed",
		}).Info("ANALYTICS_SYSTEM")
		return
	}

	defer resp.Body.Close()

	bd, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		utils.Loger.WithFields(logrus.Fields{
			"pipeline":    "FigmaVersionCheck",
			"isSended":    false,
			"error":       err,
			"reason":      -1,
			"description": "body reading failed",
		}).Info("ANALYTICS_SYSTEM")
		return
	}

	var figmaData models.FigmaResponseModel

	fmt.Println(resp)
	err = json.Unmarshal(bd, &figmaData)

	if err != nil {
		utils.Loger.WithFields(logrus.Fields{
			"pipeline":    "FigmaVersionCheck",
			"isSended":    false,
			"error":       err,
			"reason":      -1,
			"description": "parsing failed",
		}).Info("ANALYTICS_SYSTEM")
		return
	}

	if file.FileVersion.CreatedAt == figmaData.Versions[0].CreatedAt {
		utils.Loger.WithFields(logrus.Fields{
			"pipeline": "FigmaVersionCheck",
			"isSended": false,
			"reason":   3,
		}).Info("ANALYTICS_SYSTEM")
		return
	}

	file.FileVersion = &figmaData.Versions[0]
	rep.UpdateFile(file)

	msg := slackModel.PostChatMessage{
		Text:       file.FileVersion.User.Handle + " что-то поменял в макете https://www.figma.com/file/" + file.FileKey + "\n" + "Последнее изменение: " + file.FileVersion.CreatedAt,
		Channel:    file.SlackID,
		IsMarkdown: false,
	}

	err = slack.SendMessage(msg)

	utils.Loger.WithFields(logrus.Fields{
		"version":  file.FileVersion,
		"pipeline": "NotifyManagersAboutBlocked",
		"isSended": err == nil,
		"Error":    err,
	}).Info("ANALYTICS")
}
