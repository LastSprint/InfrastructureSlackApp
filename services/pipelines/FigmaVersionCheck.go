package pipelines

import (
	"bytes"
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	log "github.com/LastSprint/InfrastructureSlackApp/logging"
	models "github.com/LastSprint/InfrastructureSlackApp/models/figma"
	slackModel "github.com/LastSprint/InfrastructureSlackApp/models/slack"
	"github.com/LastSprint/InfrastructureSlackApp/repositories"
	"github.com/LastSprint/InfrastructureSlackApp/services/slack"
	"github.com/LastSprint/InfrastructureSlackApp/utils"
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
		log.PipelineByName(log.FigmaVersionCheck, err, false, log.DataReading, nil)
		return false, err
	}

	if len(files) == 0 {
		log.PipelineByName(log.FigmaVersionCheck, nil, false, log.ContentIsEmpty, nil)
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
		log.PipelineByName(log.FigmaVersionCheck, err, false, log.CantCreateRequest, file)
		return
	}

	request.Header.Add("X-Figma-Token", utils.Config.FigmaToken)

	client := http.Client{}

	resp, err := client.Do(request)

	if err != nil {
		log.PipelineByName(log.FigmaVersionCheck, err, false, log.DataReading, file)
		return
	}

	defer resp.Body.Close()

	bd, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		log.PipelineByName(log.FigmaVersionCheck, err, false, log.DataReading, resp.Body)
		return
	}

	var figmaData models.FigmaResponseModel

	fmt.Println(resp)
	err = json.Unmarshal(bd, &figmaData)

	if err != nil {
		log.PipelineByName(log.FigmaVersionCheck, err, false, log.DataReading, nil)
		return
	}

	if file.FileVersion.CreatedAt == figmaData.Versions[0].CreatedAt {
		log.PipelineByName(log.FigmaVersionCheck, nil, false, log.FigmaHistoryNotChanged, file)
		return
	}

	file.FileVersion = &figmaData.Versions[0]
	rep.UpdateFile(file)

	slackMessage := file.FileVersion.User.Handle

	slackMessage += " что-то поменял в макете https://www.figma.com/file/"

	slackMessage += file.FileKey + "\n" + "Последнее изменение: " + file.FileVersion.CreatedAt

	msg := slackModel.PostChatMessage{
		Text:       slackMessage,
		Channel:    file.SlackID,
		IsMarkdown: false,
	}

	err = slack.SendMessage(msg)

	log.PipelineByName(log.FigmaVersionCheck, err, err == nil, log.Successful, file)
}
