package models

import (
	"github.com/ppeymann/Planora.git/pkg/common"
	roompb "github.com/ppeymann/Planora.git/proto/room"
	"google.golang.org/protobuf/types/known/timestamppb"
	"gorm.io/gorm"
)

type (
	EventType string

	RoomRepository interface {
		Create(in *roompb.CreateRoomRequest) (*RoomEntity, error)

		common.BaseRepository
	}

	RoomEntity struct {
		gorm.Model
		Name      string   `gorm:"column:name;index"`
		CreatorID uint64   `gorm:"column:creator_id"`
		UserIDs   []uint64 `gorm:"column:user_ids;type:int[]"`
		TodosIDs  []uint64 `gorm:"column:todos_ids;type:int[]"`
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
	SubjectCreate EventType = "room.CREATE"
)
