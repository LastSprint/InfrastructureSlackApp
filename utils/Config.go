package utils

import (
	"encoding/json"
	"io/ioutil"
)

// DBConfigModel конфигурация БД.
type DBConfigModel struct {
	ConnectionString     string `json:"connectionString"`
	DataBaseString       string `json:"dataBaseString"`
	TestConnectionString string `json:"testConnectionString"`
	TestDataBaseString   string `json:"testDataBaseString"`
}

// ConfigModel конфигурация.
type ConfigModel struct {
	JiraLogin        string        `json:"jiraLogin"`
	JiraPassword     string        `json:"jiraPassword"`
	SlackToken       string        `json:"slackToken"`
	FigmaToken       string        `json:"figmaToken"`
	MongoDBConfig    DBConfigModel `json:"mongodb"`
	JiraBaseHost     string        `json:"jiraBaseHost"`
	JiraAPISearchURL string        `json:"jiraApiSearchUrl"`
}

// Config Глобальный конфиг
var Config ConfigModel

// LoadConfig загружает конфиг в переменную `Config` или возвращает ошибку.
func LoadConfig() error {
	data, err := ioutil.ReadFile("config/Config.json")
	if err != nil {
		return err
	}

	var conf = ConfigModel{}

	err = json.Unmarshal(data, &conf)

	if err != nil {
		return err
	}

	Config = conf

	return nil
}
