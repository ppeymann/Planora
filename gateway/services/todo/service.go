package todo

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/ppeymann/Planora.git/pkg/auth"
	"github.com/ppeymann/Planora.git/pkg/common"
	"github.com/ppeymann/Planora.git/pkg/utils"
	todopb "github.com/ppeymann/Planora.git/proto/todo"
	"github.com/ppeymann/Planora/gateway/models"
)

type service struct {
	nc *nats.Conn
}

// ChangeStatus implements models.TodoService.
func (s *service) ChangeStatus(ctx *gin.Context, status models.StatusType, id uint64) *common.BaseResult {
	claims := &auth.Claims{}
	_ = utils.CatchClaims(ctx, claims)

	req := &todopb.ChangeStatusRequest{
		Status: string(status),
		Id:     id,
		UserId: uint64(claims.Subject),
	}

	data, err := json.Marshal(req)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}
	}

	msg, err := s.nc.Request(string(models.ChangeStatus), data, 2*time.Second)
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

// GetAllTodos implements models.TodoService.
func (s *service) GetAllTodos(ctx *gin.Context) *common.BaseResult {
	claims := &auth.Claims{}
	_ = utils.CatchClaims(ctx, claims)

	req := &todopb.GetAllTodoRequest{
		UserId: uint64(claims.Subject),
	}

	data, err := json.Marshal(req)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}
	}

	msg, err := s.nc.Request(string(models.GetAll), data, 2*time.Second)
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

// UpdateTodo implements models.TodoService.
func (s *service) UpdateTodo(ctx *gin.Context, in *models.TodoInput, todoID uint64) *common.BaseResult {
	req := &todopb.UpdateTodoRequest{
		Todo: &todopb.AddTodoRequest{
			Title:       in.Title,
			Description: in.Description,
			UserId:      uint64(in.UserID),
		},
		Id: todoID,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}
	}

	msg, err := s.nc.Request(string(models.Update), data, 2*time.Second)
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
