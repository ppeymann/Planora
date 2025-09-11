package transport

import (
	"github.com/nats-io/nats.go"
	"github.com/ppeymann/Planora/room/models"
	"github.com/ppeymann/Planora/room/service"
)

func RegisterRoomSubscriber(nc *nats.Conn, roomService *service.RoomServiceServer) {
	handler := map[models.EventType]func(*nats.Msg){
		models.SubjectCreate: func(m *nats.Msg) { HandleCreate(m, roomService, nc) },
	}

	nc.Subscribe("room.*", func(m *nats.Msg) {
		if h, ok := handler[models.EventType(m.Subject)]; ok {
			h(m)
		}
	})
}
