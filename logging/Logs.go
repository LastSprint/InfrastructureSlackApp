package logging

import (
	"os"

	"github.com/sirupsen/logrus"
)

// LogerConfig тип конфигурации логера.
type LogerConfig string

const (
	// Release Релизная конфигурация.
	Release LogerConfig = "/var/log/surf_proj_infr/log"
	// Test Тестовая конфигурация.
	Test LogerConfig = "/var/log/surf_proj_infr/test_log"
)

// Loger инстанс логера по-умолчанию.
var Loger = logrus.New()

// ConfigureLog Конфигурирует логер.
// - Patameters:
//	- logType: Тип конфигурации.
func ConfigureLog(logType LogerConfig) {
	configure(string(logType))
}

func configure(path string) {
	file, err := os.OpenFile(path, os.O_CREATE|os.O_WRONLY|os.O_APPEND, 0666)

	if err != nil {
		Loger.Fatal(err)
	}

	Loger.SetFormatter(new(logrus.JSONFormatter))
	Loger.Out = file
}
