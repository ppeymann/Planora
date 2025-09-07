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

// DeleteTodo implements models.TodoService.
func (i *instrumentingService) DeleteTodo(ctx *gin.Context, id uint64) *common.BaseResult {
	defer func(begin time.Time) {
		i.requestCount.With("method", "DeleteTodo").Add(1)
		i.requestLatency.With("method", "DeleteTodo").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.DeleteTodo(ctx, id)
}

// ChangeStatus implements models.TodoService.
func (i *instrumentingService) ChangeStatus(ctx *gin.Context, status models.StatusType, id uint64) *common.BaseResult {
	defer func(begin time.Time) {
		i.requestCount.With("method", "ChangeStatus").Add(1)
		i.requestLatency.With("method", "ChangeStatus").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.ChangeStatus(ctx, status, id)
}

// GetAllTodos implements models.TodoService.
func (i *instrumentingService) GetAllTodos(ctx *gin.Context) *common.BaseResult {
	defer func(begin time.Time) {
		i.requestCount.With("method", "GetAllTodos").Add(1)
		i.requestLatency.With("method", "GetAllTodos").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.GetAllTodos(ctx)
}

// UpdateTodo implements models.TodoService.
func (i *instrumentingService) UpdateTodo(ctx *gin.Context, in *models.TodoInput, todoID uint64) *common.BaseResult {
	defer func(begin time.Time) {
		i.requestCount.With("method", "UpdateTodo").Add(1)
		i.requestLatency.With("method", "UpdateTodo").Observe(time.Since(begin).Seconds())
	}(time.Now())

	return i.next.UpdateTodo(ctx, in, todoID)
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
