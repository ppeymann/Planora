package models

type (
	EventName string
)

const (
	SignUp EventName = "user.SIGNUP"
	Login  EventName = "user.LOGIN"
)
