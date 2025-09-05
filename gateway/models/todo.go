package models

import (
	"github.com/gin-gonic/gin"
	"github.com/ppeymann/Planora.git/pkg/common"
)

type (
	// TodoService represents method signatures for api todo endpoint.
	// so any object that stratifying this interface can be used as todo service for api endpoint.
	TodoService interface {
		// AddTodo service for create new todo
		AddTodo(ctx *gin.Context, in *TodoInput) *common.BaseResult

		// UpdateTodo service for update a todo
		UpdateTodo(ctx *gin.Context, in *TodoInput, todoID uint64) *common.BaseResult

		// GetAllTodos service with specific UserID
		GetAllTodos(ctx *gin.Context) *common.BaseResult
	}

	// TodoHandler represents method signatures for todo handlers.
	// so any object that stratifying this interface can be used as todo handlers.
	TodoHandler interface {
		// AddTodo handler for create new todo
		AddTodo(ctx *gin.Context)

		// UpdateTodo handler for update a todo
		UpdateTodo(ctx *gin.Context)

		// GetAllTodos handler with specific UserID
		GetAllTodos(ctx *gin.Context)
	}

	// TodoInput for create or update todo
	//
	// swagger: model TodoInput
	TodoInput struct {
		Title       string `json:"title"`
		Description string `json:"description"`
		UserID      uint   `json:"-"`
	}
)
