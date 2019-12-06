package models

// User пользователь системы. Например разработчик.
type User struct {
	// Имя пользователя.
	FirstName string
	// Фамилия пользователя.
	LastName string
	// Идентификатор пользователя в Slack.
	SlackID string
	// Идентификатор пользователя в Jira.
	JiraID string
	// Идентификатор проекта в Jira.
	ProjectID string
	// Группа которой принадлежит пользователь.
	Member UserMember
}

// ToString делает человекочитаемую строку из объекта.
func (user *User) ToString() string {
	return "First name: " + user.FirstName + "\n" +
		"Last name: " + user.LastName + "\n" +
		"Slack ID: " + user.SlackID + "\n" +
		"Jira ID: " + user.JiraID + "\n" +
		"Project ID: " + user.ProjectID
}
