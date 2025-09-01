package pkg

import (
	"log"

	"github.com/go-kit/kit/metrics/prometheus"
	kitLog "github.com/go-kit/log"
	"github.com/nats-io/nats.go"
	"github.com/ppeymann/Planora/gateway/models"
	"github.com/ppeymann/Planora/gateway/server"
	"github.com/ppeymann/Planora/gateway/services/user"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func InitUserService(nc *nats.Conn, logger kitLog.Logger, s *server.Server) models.UserService {
	userService := user.NewService(nc)

	// getting path
	path := getSchemaPath("user")
	userService, err := user.NewValidationService(userService, path)
	if err != nil {
		log.Fatal(err)
	}

	// @Injection logging service to chain
	userService = user.NewLoggerService(kitLog.With(logger, "component", "user"), userService)

	// @Injection Instrumenting service to chain
	userService = user.NewInstrumentingService(
		prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "user",
			Name:      "request_count",
			Help:      "number of request received. ",
		}, fieldKeys),
		prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "user",
			Name:      "request_latency_microseconds",
			Help:      "total duration of request (ms). ",
		}, fieldKeys),
		userService,
	)

	// @Injection Authorization service to chain
	userService = user.NewAuthService(userService)

	_ = user.NewHandler(userService, s)

	return userService
}
