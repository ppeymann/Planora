package transport

import (
	"github.com/nats-io/nats.go"
	"github.com/ppeymann/Planora.git/pkg/common"
	"github.com/ppeymann/Planora/user/service"
)

func HandlerCreate(m *nats.Msg, u *service.UserServiceServer, nc *nats.Conn) {
	resp, err := u.SignUpService(m.Data)
	replyData := common.BuildResponse(resp, err)

	if m.Reply != "" {
		nc.Publish(m.Reply, replyData)
	}
}

func HandleLogin(m *nats.Msg, u *service.UserServiceServer, nc *nats.Conn) {
	resp, err := u.SignUpService(m.Data)
	replyData := common.BuildResponse(resp, err)

	if m.Reply != "" {
		nc.Publish(m.Reply, replyData)
	}
}

func HandleAccount(m *nats.Msg, u *service.UserServiceServer, nc *nats.Conn) {
	resp, err := u.AccountService(m.Data)
	replayData := common.BuildResponse(resp, err)

	if m.Reply != "" {
		nc.Publish(m.Reply, replayData)
	}
}
