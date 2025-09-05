package transport

import (
	"github.com/nats-io/nats.go"
	"github.com/ppeymann/Planora.git/pkg/common"
	"github.com/ppeymann/Planora/todo/service"
)

func HandleAddTodo(m *nats.Msg, t *service.TodoServiceServer, nc *nats.Conn) {
	resp, err := t.AddTodoService(m.Data)
	replayData := common.BuildResponse(resp, err)

	if m.Reply != "" {
		nc.Publish(m.Reply, replayData)
	}
}
