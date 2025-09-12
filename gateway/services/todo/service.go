package todo

import (
	"encoding/json"
	"fmt"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/ppeymann/Planora.git/pkg/auth"
	"github.com/ppeymann/Planora.git/pkg/common"
	"github.com/ppeymann/Planora.git/pkg/utils"
	roompb "github.com/ppeymann/Planora.git/proto/room"
	todopb "github.com/ppeymann/Planora.git/proto/todo"
	"github.com/ppeymann/Planora/gateway/models"
	"github.com/thoas/go-funk"
)

type service struct {
	nc *nats.Conn
}

// GetRoomTodos implements models.TodoService.
func (s *service) GetRoomTodos(ctx *gin.Context, roomID uint64) *common.BaseResult {
	claims := &auth.Claims{}
	_ = utils.CatchClaims(ctx, claims)

	roomReq := &roompb.GetUsersRequest{
		RoomId: roomID,
	}

	data, err := json.Marshal(roomReq)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}
	}

	roomMsg, err := s.nc.Request(string(models.GetUsers), data, 2*time.Second)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}
	}

	roomOutput := &common.BaseResult{}

	err = json.Unmarshal(roomMsg.Data, roomOutput)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}
	}

	resultMap, ok := roomOutput.Result.(map[string]interface{})
	if !ok {
		return &common.BaseResult{
			Errors: []string{"invalid result format"},
			Status: http.StatusOK,
		}
	}

	userIDsIface, ok := resultMap["user_ids"].([]interface{})
	if !ok {
		return &common.BaseResult{
			Errors: []string{"invalid user_ids format"},
			Status: http.StatusOK,
		}
	}

	var userIDs []uint64
	for _, v := range userIDsIface {
		if id, ok := v.(float64); ok { // JSON اعداد رو float64 می‌کنه
			userIDs = append(userIDs, uint64(id))
		}
	}

	fmt.Println(userIDs)

	if !funk.Contains(userIDs, uint64(claims.Subject)) {
		return &common.BaseResult{
			Errors: []string{"permission denied"},
			Status: http.StatusOK,
		}
	}

	todoReq := &todopb.RoomTodosRequest{
		RoomId: roomID,
	}

	todoData, err := json.Marshal(todoReq)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}
	}

	todoMsg, err := s.nc.Request(string(models.RoomTodo), todoData, 2*time.Second)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}
	}

	op := &common.BaseResult{}
	err = json.Unmarshal(todoMsg.Data, op)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}
	}

	return op

}

// DeleteTodo implements models.TodoService.
func (s *service) DeleteTodo(ctx *gin.Context, id uint64) *common.BaseResult {
	claims := &auth.Claims{}
	_ = utils.CatchClaims(ctx, claims)

	req := &todopb.DeleteTodoRequest{
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

	msg, err := s.nc.Request(string(models.Delete), data, 2*time.Second)
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
		RoomId:      uint64(in.RoomID),
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
