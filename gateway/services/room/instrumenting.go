package room

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
	next           models.RoomService
}

// GetRoom implements models.RoomService.
func (i *instrumentingService) GetRoom(ctx *gin.Context, roomID uint64) *common.BaseResult {
	defer func(begin time.Time) {
		i.requestCount.With("method", "GetRoom").Add(1)
		i.requestLatency.With("method", "GetRoom").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.GetRoom(ctx, roomID)
}

// Create implements models.RoomService.
func (i *instrumentingService) Create(ctx *gin.Context, in *models.RoomInput) *common.BaseResult {
	defer func(begin time.Time) {
		i.requestCount.With("method", "Create").Add(1)
		i.requestLatency.With("method", "Create").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.Create(ctx, in)
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, service models.RoomService) models.RoomService {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		next:           service,
	}
}
