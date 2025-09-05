package todo

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
	next           models.TodoService
}

// AddTodo implements models.TodoService.
func (i *instrumentingService) AddTodo(ctx *gin.Context, in *models.TodoInput) *common.BaseResult {
	defer func(begin time.Time) {
		i.requestCount.With("method", "AddTodo").Add(1)
		i.requestLatency.With("method", "AddTodo").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.AddTodo(ctx, in)
}

func NewInstrumentingService(counter metrics.Counter, latency metrics.Histogram, service models.TodoService) models.TodoService {
	return &instrumentingService{
		requestCount:   counter,
		requestLatency: latency,
		next:           service,
	}
}
