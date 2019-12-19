package pipelines

import (
	// log "github.com/LastSprint/InfrastructureSlackApp/logging"
	"github.com/LastSprint/InfrastructureSlackApp/repositories"
)

// AddReasonToOverwork выбирает из БД разработчиков, затем загружает задачи из Jira. 
// Проверяет разницу между logged и estimate и если произведение одного на второе больше 1.5, то разработчику отправляется просьба оценить задача. 
// При этом из Jira выбираются задачи, у которых поле reason ни чем не заполнено. 
// Из всего списка полученных задач случайным образом выбирается одна задача.
type AddReasonToOverwork struct {
	// Репозиторий для чтения пользователей.
	Repo repositories.UserRepository
}
