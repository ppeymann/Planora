package user

import (
	"strings"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/log"
	"github.com/ppeymann/Planora.git/pkg/common"
	userpb "github.com/ppeymann/Planora.git/proto/user"
	"github.com/ppeymann/Planora/gateway/models"
)

type loggerService struct {
	logger log.Logger
	next   models.UserService
}

// SignUp implements models.UserService.
func (l *loggerService) SignUp(ctx *gin.Context, in *userpb.SignUpRequest) (result *common.BaseResult) {
	defer func(begin time.Time) {
		_ = l.logger.Log(
			"method", "SignUp",
			"errors", strings.Join(result.Errors, " ,"),
			"result", result,
			"input", in,
			"client_ip", ctx.ClientIP(),
			"took", time.Since(begin),
		)
	}(time.Now())

	return l.next.SignUp(ctx, in)
}

func NewLoggerService(logger log.Logger, service models.UserService) models.UserService {
	return &loggerService{
		logger: logger,
		next:   service,
	}
}
