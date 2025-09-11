package models

type (
	UserEventName string
	TodoEventName string
	RoomEventName string
)

const (
	SignUp  UserEventName = "user.SIGNUP"
	Login   UserEventName = "user.LOGIN"
	Account UserEventName = "user.ACCOUNT"
)

const (
	Add          TodoEventName = "todo.ADD"
	Update       TodoEventName = "todo.UPDATE"
	GetAll       TodoEventName = "todo.GETALL"
	ChangeStatus TodoEventName = "todo.CHANGE_STATUS"
	Delete       TodoEventName = "todo.DELETE"
)

const (
	CreateRoom RoomEventName = "room.CREATE"
)
