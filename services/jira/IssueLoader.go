package jira

import (
	"encoding/json"
	"fmt"
	"io/ioutil"
	"net/http"

	mj "github.com/LastSprint/InfrastructureSlackApp/models/jira"
	"github.com/LastSprint/InfrastructureSlackApp/utils"
)

// LoadIssues Загружает все задачи для конкретного пользователя
// Parameters:
// 	- user: Пользователь, для которого нужно выгрузить задачи
//	- statuses: Статусы, в которых должны быть задачи
func LoadIssues(request RequestConvertible) (mj.IssueSearchWrapperEntity, error) {
	result := mj.IssueSearchWrapperEntity{}

	req, err := http.NewRequest(
		"GET", utils.Config.JiraAPISearchURL, nil,
	)

	if err != nil {
		return result, err
	}

	req.SetBasicAuth(utils.Config.JiraLogin, utils.Config.JiraPassword)

	qr := req.URL.Query()

	jql := request.MakeJiraRequest()
	fmt.Println(jql)
	qr.Set("jql", jql)
	qr.Set("startAt", "0")
	qr.Set("maxResults", "500")
	qr.Set("fields", utils.JoinByCharacter(acceptedFields, ",", ""))

	req.URL.RawQuery = qr.Encode()
	client := http.Client{}
	resp, err := client.Do(req)

	if err != nil {
		return result, err
	}

	defer resp.Body.Close()

	bd, err := ioutil.ReadAll(resp.Body)

	if err != nil {
		return result, err
	}

	err = json.Unmarshal(bd, &result)

	if err != nil {
		return result, err
	}

	return result, err
}
