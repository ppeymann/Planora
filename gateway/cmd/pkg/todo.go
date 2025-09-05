package pkg

import (
	"log"

	"github.com/go-kit/kit/metrics/prometheus"
	kitLog "github.com/go-kit/log"
	"github.com/nats-io/nats.go"
	"github.com/ppeymann/Planora/gateway/models"
	"github.com/ppeymann/Planora/gateway/server"
	"github.com/ppeymann/Planora/gateway/services/todo"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func InitTodoService(nc *nats.Conn, logger kitLog.Logger, s *server.Server) models.TodoService {
	todoService := todo.NewService(nc)

	// getting path
	path := getSchemaPath("todo")
	todoService, err := todo.NewValidationService(path, todoService)
	if err != nil {
		log.Fatal(err)
	}

	// @Injection logging service to chain
	todoService = todo.NewLoggerService(kitLog.With(logger, "component", "todo"), todoService)

	// @Injection Instrumenting service to chain
	todoService = todo.NewInstrumentingService(
		prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "todo",
			Name:      "request_count",
			Help:      "number of request received. ",
		}, fieldKeys),
		prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "todo",
			Name:      "request_latency_microseconds",
			Help:      "total duration of request (ms). ",
		}, fieldKeys),
		todoService,
	)

	// @Injection Authorization service to chain
	todoService = todo.NewAuthService(todoService)

	_ = todo.NewHandler(todoService, s)

	return todoService
}
