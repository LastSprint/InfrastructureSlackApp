package models

// UserDepartment отдел пользователя.
type UserDepartment string

const (
	// IOS - iOS отдел.
	IOS UserDepartment = "iOS"
	// Android отдел.
	Android UserDepartment = "Android"
	// Flutter отдел.
	Flutter UserDepartment = "Flutter"
	// Managers отдел менеджеров.
	Managers UserDepartment = "managers"
)

// UserRole роль пользователя в его отделе.
type UserRole string

const (
	// Lead руководитель.
	Lead UserRole = "lead"
	// Developer разработчик.
	Developer UserRole = "developer"
	// Manager менеджер.
	Manager UserRole = "manager"
)

// UserMember это описание группы, к которой принадлежит пользователь.
// Обозначает скорее группу рассылки.
type UserMember struct {
	Department UserDepartment
	Role       UserRole
}

// ToString делает человекочитаемую строку из объекта.
func (m UserMember) ToString() string {
	return "department: " + string(m.Department) + "\n" + "role: " + string(m.Role)
}
