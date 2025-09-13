package transport

import (
	"github.com/nats-io/nats.go"

	"github.com/ppeymann/Planora/todo/models"
	"github.com/ppeymann/Planora/todo/service"
)

func RegisterTodoSubscriber(nc *nats.Conn, todoService *service.TodoServiceServer) {
	handler := map[models.EventType]func(*nats.Msg){
		models.SubjectAddTodo:      func(m *nats.Msg) { HandleAddTodo(m, todoService, nc) },
		models.SubjectUpdateTodo:   func(m *nats.Msg) { HandleUpdateTodo(m, todoService, nc) },
		models.SubjectGetAllTodo:   func(m *nats.Msg) { HandleGetAllTodo(m, todoService, nc) },
		models.SubjectChangeStatus: func(m *nats.Msg) { HandleChangeStatus(m, todoService, nc) },
		models.SubjectDeleteTodo:   func(m *nats.Msg) { HandleDeleteTodo(m, todoService, nc) },
		models.SubjectGetRoomTodo:  func(m *nats.Msg) { HandleGetRoomTodo(m, todoService, nc) },
		models.SubjectGetTodoGrpc:  func(m *nats.Msg) { HandleGetTodosGrpc(m, todoService, nc) },
	}

	nc.Subscribe("todo.*", func(m *nats.Msg) {
		if h, ok := handler[models.EventType(m.Subject)]; ok {
			h(m)
		}
	})
}
