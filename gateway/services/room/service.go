package room

import (
	"encoding/json"
	"net/http"
	"time"

	"github.com/gin-gonic/gin"
	"github.com/nats-io/nats.go"
	"github.com/ppeymann/Planora.git/pkg/common"
	roompb "github.com/ppeymann/Planora.git/proto/room"
	"github.com/ppeymann/Planora/gateway/models"
)

type service struct {
	nc *nats.Conn
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
