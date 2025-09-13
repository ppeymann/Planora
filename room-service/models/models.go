package models

import (
	"github.com/lib/pq"
	"github.com/ppeymann/Planora.git/pkg/common"
	roompb "github.com/ppeymann/Planora.git/proto/room"
	todopb "github.com/ppeymann/Planora.git/proto/todo"
	userpb "github.com/ppeymann/Planora.git/proto/user"
	"gorm.io/gorm"
)

type (
	EventType string

	RoomRepository interface {
		// Create new room
		Create(in *roompb.CreateRoomRequest) (*RoomEntity, error)

		// GetUsers with specific room ID
		GetUsers(roomID uint64) ([]uint64, error)

		// GetRoom with room ID
		GetRoom(roomID uint) (*RoomEntity, error)

		common.BaseRepository
	}

	RoomEntity struct {
		gorm.Model
		Name      string        `gorm:"column:name;index"`
		CreatorID uint64        `gorm:"column:creator_id"`
		UserIDs   pq.Int64Array `gorm:"column:user_ids;type:bigint[]"`
		TodosIDs  pq.Int64Array `gorm:"column:todos_ids;type:bigint[]"`
	}

	RoomResponse struct {
		Room  *roompb.Room
		Users []*userpb.User
		Todos []*todopb.Todo
	}
)

const (
	SubjectCreate       EventType = "room.CREATE"
	SubjectGetUsers     EventType = "room.GET_USERS"
	SubjectGetRoomUsers EventType = "user.ROOM_USERS"
	SubjectGetTodoGrpc  EventType = "todo.GET_TODO_GRPC"
	SubjectGetRoom      EventType = "room.GET_ROOM"
)
