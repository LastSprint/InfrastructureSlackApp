package slack

import (
	"bytes"
	"encoding/json"
	"fmt"
	"net/http"

	"github.com/LastSprint/InfrastructureSlackApp/models/slack"
	"github.com/LastSprint/InfrastructureSlackApp/utils"
)

// SendMessage отправляет сообщение в Slack
// Parameters:
//  - message: сообщение, которое будет отправлено в slack
func SendMessage(message slack.PostChatMessage) error {

	data, err := json.Marshal(message)

	if err != nil {
		return err
	}

	request, err := http.NewRequest(
		"POST",
		"https://slack.com/api/chat.postMessage",
		bytes.NewBuffer(data),
	)

	if err != nil {
		return err
	}

	request.Header.Add("Authorization", "Bearer "+utils.Config.SlackToken)
	request.Header.Add("Content-Type", "application/json")
	fmt.Println(request)

	client := http.Client{}

	resp, err := client.Do(request)

	if err != nil {
		return err
	}

	fmt.Println(resp.Status)
	resp.Body.Close()

	return nil
}
