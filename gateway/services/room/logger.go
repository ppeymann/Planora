package room

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
	next   models.RoomService
}

// GetRoom implements models.RoomService.
func (l *loggerService) GetRoom(ctx *gin.Context, roomID uint64) (result *common.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "GetRoom",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"id", roomID,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.GetRoom(ctx, roomID)
}

// Create implements models.RoomService.
func (l *loggerService) Create(ctx *gin.Context, in *models.RoomInput) (result *common.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "Create",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"input", in,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.Create(ctx, in)
}

func NewLoggerService(logger log.Logger, service models.RoomService) models.RoomService {
	return &loggerService{
		logger: logger,
		next:   service,
	}
}
