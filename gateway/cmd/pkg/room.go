package pkg

import (
	"log"

	"github.com/go-kit/kit/metrics/prometheus"
	kitLog "github.com/go-kit/log"
	"github.com/nats-io/nats.go"
	"github.com/ppeymann/Planora/gateway/models"
	"github.com/ppeymann/Planora/gateway/server"
	"github.com/ppeymann/Planora/gateway/services/room"
	stdprometheus "github.com/prometheus/client_golang/prometheus"
)

func InitRoomService(nc *nats.Conn, logger kitLog.Logger, s *server.Server) models.RoomService {
	roomService := room.NewService(nc)

	// getting path
	path := getSchemaPath("room")
	roomService, err := room.NewValidationService(path, roomService)
	if err != nil {
		log.Fatal(err)
	}

	// @Injection logging service to chain
	roomService = room.NewLoggerService(kitLog.With(logger, "component", "room"), roomService)

	// @Injection Instrumenting service to chain
	roomService = room.NewInstrumentingService(
		prometheus.NewCounterFrom(stdprometheus.CounterOpts{
			Namespace: "api",
			Subsystem: "room",
			Name:      "request_count",
			Help:      "number of request received. ",
		}, fieldKeys),
		prometheus.NewSummaryFrom(stdprometheus.SummaryOpts{
			Namespace: "api",
			Subsystem: "room",
			Name:      "request_latency_microseconds",
			Help:      "total duration of request (ms). ",
		}, fieldKeys),
		roomService,
	)

	// @Injection Authorization service to chain
	roomService = room.NewAuthService(roomService)

	_ = room.NewHandler(roomService, s)

	return roomService
}
