package transport

import (
	"github.com/nats-io/nats.go"
	"github.com/ppeymann/Planora/user/models"
	"github.com/ppeymann/Planora/user/service"
)

func RegisterUserSubscriber(nc *nats.Conn, userService *service.UserServiceServer) {
	handler := map[models.EventType]func(*nats.Msg){
		models.Signup: func(m *nats.Msg) { HandlerCreate(m, userService, nc) },
	}

	nc.Subscribe("user.*", func(m *nats.Msg) {
		if h, ok := handler[models.EventType(m.Subject)]; ok {
			h(m)
		}
	})
}
