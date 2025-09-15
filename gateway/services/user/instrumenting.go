package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/metrics"
	"github.com/ppeymann/Planora.git/pkg/common"
	"github.com/ppeymann/Planora/gateway/models"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           models.UserService
}

// Account implements models.UserService.
func (i *instrumentingService) Account(ctx *gin.Context) *common.BaseResult {
	defer func(begin time.Time) {
		i.requestCount.With("method", "Account").Add(1)
		i.requestLatency.With("method", "Account").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.Account(ctx)
}

// Login implements models.UserService.
func (i *instrumentingService) Login(ctx *gin.Context, in *models.LoginInput) *common.BaseResult {
	defer func(begin time.Time) {
		i.requestCount.With("method", "Login").Add(1)
		i.requestLatency.With("method", "Login").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.Login(ctx, in)
}

// SignUp implements models.UserService.
func (i *instrumentingService) SignUp(ctx *gin.Context, in *models.SignUpInput) *common.BaseResult {
	defer func(begin time.Time) {
		i.requestCount.With("method", "SignUp").Add(1)
		i.requestLatency.With("method", "SignUp").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.SignUp(ctx, in)
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, service models.UserService) models.UserService {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		next:           service,
	}
}
