package transport

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/nats-io/nats.go"
	"github.com/ppeymann/Planora.git/pkg/common"
	todopb "github.com/ppeymann/Planora.git/proto/todo"
	userpb "github.com/ppeymann/Planora.git/proto/user"
	"github.com/ppeymann/Planora/room/models"
	"github.com/ppeymann/Planora/room/service"
	"github.com/ppeymann/Planora/room/utils"
)

func HandleCreate(m *nats.Msg, r *service.RoomServiceServer, nc *nats.Conn) {
	resp, err := r.CreateService(m.Data)
	replyData := common.BuildResponse(resp, err)

	if m.Reply != "" {
		nc.Publish(m.Reply, replyData)
	}
}

func HandleGetUsers(m *nats.Msg, r *service.RoomServiceServer, nc *nats.Conn) {
	resp, err := r.GetUsersService(m.Data)
	replyData := common.BuildResponse(resp, err)

	if m.Reply != "" {
		nc.Publish(m.Reply, replyData)
	}
}

func HandleGetRoom(m *nats.Msg, r *service.RoomServiceServer, nc *nats.Conn) {
	resp, err := r.GetRoomService(m.Data)
	if err != nil {
		utils.ReturnError(err, nc, m)
		return
	}

	result := &models.RoomResponse{
		Room: resp.Room,
	}

	// <----- GET USER ----->
	userReq := &userpb.GetRoomUsersRequest{
		Ids: resp.Room.UserIds,
	}

	userData, err := json.Marshal(userReq)
	if err != nil {
		utils.ReturnError(err, nc, m)
		return
	}

	userMsg, err := nc.Request(string(models.SubjectGetRoomUsers), userData, 2*time.Second)
	if err != nil {
		utils.ReturnError(err, nc, m)
		return
	}

	userRes := &userpb.GetRoomUsersResponse{}
	err = json.Unmarshal(userMsg.Data, userRes)
	if err != nil {
		utils.ReturnError(err, nc, m)
		return
	}

	result.Users = userRes.Users

	// <----- GET TODO ---->
	todoReq := &todopb.RoomTodosRequest{
		RoomId: resp.Room.Model.Id,
	}

	todoData, err := json.Marshal(todoReq)
	if err != nil {
		utils.ReturnError(err, nc, m)
		return
	}

	todoMsg, err := nc.Request(string(models.SubjectGetTodoGrpc), todoData, 2*time.Second)
	if err != nil {
		utils.ReturnError(err, nc, m)
		return
	}

	todoRes := &todopb.RoomTodosResponse{}
	err = json.Unmarshal(todoMsg.Data, todoRes)
	if err != nil {
		utils.ReturnError(err, nc, m)
		return
	}

	result.Todos = todoRes.Todos

	// send result to client
	replyData := &common.BaseResult{
		Result: result,
		Status: http.StatusOK,
	}

	data, err := json.Marshal(replyData)
	if err != nil {
		utils.ReturnError(err, nc, m)
		return
	}

	if m.Reply != "" {
		nc.Publish(m.Reply, data)
	}
}
