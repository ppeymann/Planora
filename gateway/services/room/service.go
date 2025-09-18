package room

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/ppeymann/Planora.git/pkg/auth"
	"github.com/ppeymann/Planora.git/pkg/common"
	"github.com/ppeymann/Planora.git/pkg/utils"
	roompb "github.com/ppeymann/Planora.git/proto/room"
	"github.com/ppeymann/Planora/gateway/models"
)

type service struct {
	nc *nats.Conn
}

// AddUser implements models.RoomService.
func (s *service) AddUser(ctx *gin.Context, in *models.AddUserInput) *common.BaseResult {
	req := &roompb.AddUserRequest{
		Username:  in.Username,
		CreatorId: in.CreatorID,
		RoomId:    in.RoomID,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}
	}

	msg, err := s.nc.Request(string(models.AddUser), data, 2*time.Second)
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

// GetRoom implements models.RoomService.
func (s *service) GetRoom(ctx *gin.Context, roomID uint64) *common.BaseResult {
	claims := &auth.Claims{}
	_ = utils.CatchClaims(ctx, claims)

	req := &roompb.GetRoomRequest{
		RoomId:    roomID,
		CreatorId: uint64(claims.Subject),
	}

	data, err := json.Marshal(req)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}
	}

	msg, err := s.nc.Request(string(models.GetRoom), data, 2*time.Second)
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

// Create implements models.RoomService.
func (s *service) Create(ctx *gin.Context, in *models.RoomInput) *common.BaseResult {
	req := &roompb.CreateRoomRequest{
		Name:      in.Name,
		CreatorId: in.CreatorID,
	}

	data, err := json.Marshal(req)
	if err != nil {
		return &common.BaseResult{
			Errors: []string{err.Error()},
			Status: http.StatusOK,
		}
	}

	msg, err := s.nc.Request(string(models.CreateRoom), data, 2*time.Second)
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

func NewService(nc *nats.Conn) models.RoomService {
	return &service{
		nc: nc,
	}
}
