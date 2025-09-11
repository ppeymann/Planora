package transport

import (
	"github.com/nats-io/nats.go"
	"github.com/ppeymann/Planora.git/pkg/common"
	"github.com/ppeymann/Planora/room/service"
)

func HandleCreate(m *nats.Msg, r *service.RoomServiceServer, nc *nats.Conn) {
	resp, err := r.CreateService(m.Data)
	replyData := common.BuildResponse(resp, err)

	if m.Reply != "" {
		nc.Publish(m.Reply, replyData)
	}
}
