package todo

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	"github.com/ppeymann/Planora.git/pkg/common"
	"github.com/ppeymann/Planora/gateway/models"
)

type loggerService struct {
	logger log.Logger
	next   models.TodoService
}

// DeleteTodo implements models.TodoService.
func (l *loggerService) DeleteTodo(ctx *gin.Context, id uint64) (result *common.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "DeleteTodo",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"id", id,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.DeleteTodo(ctx, id)
}

// ChangeStatus implements models.TodoService.
func (l *loggerService) ChangeStatus(ctx *gin.Context, status models.StatusType, id uint64) (result *common.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "ChangeStatus",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"id", id,
			"status", status,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.ChangeStatus(ctx, status, id)
}

// GetAllTodos implements models.TodoService.
func (l *loggerService) GetAllTodos(ctx *gin.Context) (result *common.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "GetAllTodos",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.GetAllTodos(ctx)
}

// UpdateTodo implements models.TodoService.
func (l *loggerService) UpdateTodo(ctx *gin.Context, in *models.TodoInput, todoID uint64) (result *common.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "UpdateTodo",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"input", in,
			"todo_id", todoID,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.UpdateTodo(ctx, in, todoID)
}

// AddTodo implements models.TodoService.
func (l *loggerService) AddTodo(ctx *gin.Context, in *models.TodoInput) (result *common.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "AddTodo",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"input", in,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.AddTodo(ctx, in)
}

func NewLoggerService(logger log.Logger, service models.TodoService) models.TodoService {
	return &loggerService{
		logger: logger,
		next:   service,
	}
}
