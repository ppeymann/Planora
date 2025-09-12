package models

import (
	"github.com/lib/pq"
	"github.com/ppeymann/Planora.git/pkg/common"
	roompb "github.com/ppeymann/Planora.git/proto/room"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type (
	EventType string

	RoomRepository interface {
		// Create new room
		Create(in *roompb.CreateRoomRequest) (*RoomEntity, error)

		// GetUsers with specific room ID
		GetUsers(roomID uint64) ([]uint64, error)

		common.BaseRepository
	}

	RoomEntity struct {
		gorm.Model
		Name      string        `gorm:"column:name;index"`
		CreatorID uint64        `gorm:"column:creator_id"`
		UserIDs   pq.Int64Array `gorm:"column:user_ids;type:bigint[]"`
		TodosIDs  pq.Int64Array `gorm:"column:todos_ids;type:bigint[]"`
	}
)

func ToBaseModel(t *RoomEntity) *roompb.BaseModel {
	return &roompb.BaseModel{
		Id:         uint64(t.ID),
		CreatedAt:  timestamppb.New(t.CreatedAt),
		UpdatedeAt: timestamppb.New(t.UpdatedAt),
	}
}

const (
	SubjectCreate   EventType = "room.CREATE"
	SubjectGetUsers EventType = "room.GET_USERS"
)
