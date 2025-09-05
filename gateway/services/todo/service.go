package todo

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/ppeymann/Planora.git/pkg/common"
	todopb "github.com/ppeymann/Planora.git/proto/todo"
	"github.com/ppeymann/Planora/gateway/models"
)

type service struct {
	nc *nats.Conn
}

// AddTodo implements models.TodoService.
func (s *service) AddTodo(ctx *gin.Context, in *models.TodoInput) *common.BaseResult {
	req := &todopb.AddTodoRequest{
		Title:       in.Title,
		Description: in.Description,
		UserId:      uint64(in.UserID),
	}

	data, err := json.Marshal(req)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}
	}

	msg, err := s.nc.Request(string(models.Add), data, 2*time.Second)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}
	}

	op := &common.BaseResult{}
	err = json.Unmarshal(msg.Data, op)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}
	}

	return op
}

func NewService(nc *nats.Conn) models.TodoService {
	return &service{
		nc: nc,
	}
}
