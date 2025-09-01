package models

import "errors"

// Errors
var (
	ErrAccountExist            error = errors.New("account with specified params already exists")
	ErrSignInFailed            error = errors.New("account not found or password error")
	ErrPermissionDenied        error = errors.New("specified role is not available for user")
	ErrAccountNotExist         error = errors.New("specified account does not exist")
	ErrProvideRequiredJsonBody error = errors.New("please provide required JSON body")
)
