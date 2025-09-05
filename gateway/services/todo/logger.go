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
