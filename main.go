package main

import (
	"fmt"
	"log"
	"os"
	"time"

	"github.com/LastSprint/InfrastructureSlackApp/logging"
	rep "github.com/LastSprint/InfrastructureSlackApp/repositories"
	"github.com/LastSprint/InfrastructureSlackApp/services/pipelines"
	"github.com/LastSprint/InfrastructureSlackApp/utils"
	"github.com/sirupsen/logrus"
)

type commandName string

const (
	help                       commandName = "help"
	notifyAllNeedsRate         commandName = "notifyAllNeedsRate"
	notifyManagerAboutBlockers commandName = "notifyManagerAboutBlockers"
	figmaVersionCheck          commandName = "figmaVersionCheck"
)

func main() {

	if err := utils.LoadConfig(); err != nil {
		log.Fatal(err)
	}

	if len(os.Args) < 3 {
		fmt.Println("***********")
		fmt.Println("Ошибка! Не указана команда!")
		fmt.Println("***********")
		fmt.Println(getHelpMsg())
		return
	}

	command := commandName(os.Args[2])

	if command == help {
		fmt.Println(getHelpMsg())
		return
	}

	db, err := initialize()

	if err != nil {
		logging.Loger.WithFields(logrus.Fields{
			"Error": err,
		}).Fatal("Cant initialize DB context")
	}

	pipeline := startPipiline(commandName(os.Args[2]), db)

	if pipeline == nil {
		logging.Loger.WithFields(logrus.Fields{
			"args": os.Args[1:],
		}).Debug("Pipileni is nil")
		return
	}

	startTime := time.Now()

	pipeline.InitPipeline()

	logging.Loger.WithFields(logrus.Fields{
		"pipeline":  os.Args[2],
		"startTime": startTime,
		"endTime":   time.Now(),
	}).Info("ANALYTICS_SYSTEM_TIME")

	err = db.Close()

	if err == nil {
		return
	}

	logging.Loger.WithFields(logrus.Fields{
		"Error": err,
	}).Error("Can't close DB connection")
}

func initialize() (*rep.DBContext, error) {
	args := os.Args[1:]

	if args[0] == "release" {
		logging.ConfigureLog(logging.Release)
	} else {
		logging.ConfigureLog(logging.Test)
	}

	return getDbContext(args)
}

func getDbContext(args []string) (*rep.DBContext, error) {
	if args[0] == "release" {
		return rep.NewDB()
	}
	return rep.NewTestDB()
}

func startPipiline(command commandName, db *rep.DBContext) pipelines.Pipeline {

	switch command {
	case notifyAllNeedsRate:
		rep := &rep.UserDBRepository{DB: db}
		return &pipelines.NotifyAllNeedsRate{Repo: rep}
	case notifyManagerAboutBlockers:
		rep := &rep.UserDBRepository{DB: db}
		return &pipelines.NotifyManagersAboutBlocked{Repo: rep}
	case figmaVersionCheck:
		rep := &rep.FigmaFileDBRepository{DB: db}
		return &pipelines.FigmaVersionCheck{Repo: rep}
	}

	return nil
}

func getHelpMsg() string {
	return "Работает в двух разных стратегиях:\n" +
		"`test` - испольняет работу на тестовой базе данных\n" +
		"`release` - исполняет работу на релизной базе данных\n" +
		"*Первым должен идти параметр стратегии*\n" +
		"-------\n" +
		"Пайплайны:\n" +
		"`notifyAllNeedsRate` - оповещает всех лидов и разработчиков о задачах с 0 Estimate и Remaining\n" +
		"`notifyManagerAboutBlockers` - оповещает всех менеджеров о задачах в статусе \"Blocked\"\n" +
		"`figmaVersionCheck` - проверяет все зарегистрированный figma проекты (файлы) и в случае, если там были изменения сообщает о них в slack"
}
