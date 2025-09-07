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

func HandleUpdateTodo(m *nats.Msg, t *service.TodoServiceServer, nc *nats.Conn) {
	resp, err := t.UpdateTodoService(m.Data)
	replayData := common.BuildResponse(resp, err)

	if m.Reply != "" {
		nc.Publish(m.Reply, replayData)
	}
}

func HandleGetAllTodo(m *nats.Msg, t *service.TodoServiceServer, nc *nats.Conn) {
	resp, err := t.GetAllTodoService(m.Data)
	replay := common.BuildResponse(resp, err)

	if m.Reply != "" {
		nc.Publish(m.Reply, replay)
	}
}

func HandleChangeStatus(m *nats.Msg, t *service.TodoServiceServer, nc *nats.Conn) {
	resp, err := t.ChangeStatusService(m.Data)
	replay := common.BuildResponse(resp, err)

	if m.Reply != "" {
		nc.Publish(m.Reply, replay)
	}
}

func HandleDeleteTodo(m *nats.Msg, t *service.TodoServiceServer, nc *nats.Conn) {
	resp, err := t.DeleteTodoService(m.Data)
	reply := common.BuildResponse(resp, err)

	if m.Reply != "" {
		nc.Publish(m.Reply, reply)
	}
}
