package models

type (
	UserEventName string
	TodoEventName string
)

const (
	SignUp  UserEventName = "user.SIGNUP"
	Login   UserEventName = "user.LOGIN"
	Account UserEventName = "user.ACCOUNT"
)

const (
	Add    TodoEventName = "todo.ADD"
	Update TodoEventName = "todo.UPDATE"
)
