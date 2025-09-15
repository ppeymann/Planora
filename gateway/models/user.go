package models

import (
	"github.com/gin-gonic/gin"
	"github.com/ppeymann/Planora.git/pkg/common"
)

type (
	// UserService represents method signatures for api user endpoint.
	// so any object that stratifying this interface can be used as user service for api endpoint.
	UserService interface {
		// SignUp service is for signing up user if never sign in
		SignUp(ctx *gin.Context, in *SignUpInput) *common.BaseResult

		// Login service is for log in user if signed up
		Login(ctx *gin.Context, in *LoginInput) *common.BaseResult

		// Account service is for getting user account info
		Account(ctx *gin.Context) *common.BaseResult
	}

	// UserHandler represents method signatures for user handlers.
	// so any object that stratifying this interface can be used as user handlers.
	UserHandler interface {
		// SignUp handler is for signing up user if never sign in
		SignUp(ctx *gin.Context)

		// Login handler is for log in user if signed up
		Login(ctx *gin.Context)

		// Account handler is for getting user account info
		Account(ctx *gin.Context)
	}

	SignUpInput struct {
		// Username
		Username string `json:"username"`

		// Password
		Password string `json:"password"`

		// Email
		Email string `json:"email"`

		// FirstName
		FirstName string `json:"first_name"`

		// LastName
		LastName string `json:"last_name"`
	}

	LoginInput struct {
		// Username
		Username string `json:"username"`

		// Password
		Password string `json:"password"`
	}
)
