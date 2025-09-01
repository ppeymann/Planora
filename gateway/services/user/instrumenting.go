package user

import (
	"time"

	"github.com/gin-gonic/gin"
	"github.com/go-kit/kit/metrics"
	"github.com/ppeymann/Planora.git/pkg/common"
	userpb "github.com/ppeymann/Planora.git/proto/user"
	"github.com/ppeymann/Planora/gateway/models"
)

type instrumentingService struct {
	requestCount   metrics.Counter
	requestLatency metrics.Histogram
	next           models.UserService
}

// SignUp implements models.UserService.
func (i *instrumentingService) SignUp(ctx *gin.Context, in *userpb.SignUpRequest) *common.BaseResult {
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
