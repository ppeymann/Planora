package utils

import (
	"encoding/json"
	"net/http"

	"github.com/lib/pq"
	"github.com/nats-io/nats.go"
	"github.com/ppeymann/Planora.git/pkg/common"
	roompb "github.com/ppeymann/Planora.git/proto/room"
	"github.com/ppeymann/Planora/room/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

func ToBaseModel(t *models.RoomEntity) *roompb.BaseModel {
	return &roompb.BaseModel{
		Id:         uint64(t.ID),
		CreatedAt:  timestamppb.New(t.CreatedAt),
		UpdatedeAt: timestamppb.New(t.UpdatedAt),
	}
}

func ToProtoUintIDs(IDs pq.Int64Array) []uint64 {
	result := make([]uint64, len(IDs))
	for i, v := range IDs {
		result[i] = uint64(v)
	}
	return result
}

func ReturnError(err error, nc *nats.Conn, m *nats.Msg) {
	replyData := &common.BaseResult{
		Errors: []string{err.Error()},
		Status: http.StatusOK,
	}

	data, _ := json.Marshal(replyData)

	if m.Reply != "" {
		nc.Publish(m.Reply, data)
	}
}
