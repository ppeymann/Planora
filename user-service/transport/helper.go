package transport

import (
	"encoding/json"
	"net/http"

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
	resp, err := u.LoginService(m.Data)
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

func HandleGetRoomUsers(m *nats.Msg, u *service.UserServiceServer, nc *nats.Conn) {
	resp, err := u.GetRoomUsersService(m.Data)
	if err != nil {
		replyData := &common.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}

		data, _ := json.Marshal(replyData)

		if m.Reply != "" {
			nc.Publish(m.Reply, data)
		}
	}

	data, _ := json.Marshal(resp)

	if m.Reply != "" {
		nc.Publish(m.Reply, data)
	}
}
