package models

import (
	"github.com/gin-gonic/gin"
	"github.com/ppeymann/Planora.git/pkg/common"
	userpb "github.com/ppeymann/Planora.git/proto/user"
)

type (
	// UserService represents method signatures for api user endpoint.
	// so any object that stratifying this interface can be used as user service for api endpoint.
	UserService interface {
		// SignUp service is for signing up user if never sign in
		SignUp(ctx *gin.Context, in *userpb.SignUpRequest) *common.BaseResult

		// Login service is for log in user if signed up
		Login(ctx *gin.Context, in *userpb.LoginRequest) *common.BaseResult
	}

	// UserHandler represents method signatures for user handlers.
	// so any object that stratifying this interface can be used as user handlers.
	UserHandler interface {
		// SignUp handler is for signing up user if never sign in
		SignUp(ctx *gin.Context)

		// Login handler is for log in user if signed up
		Login(ctx *gin.Context)
	}
)
