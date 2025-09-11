package models

import (
	roompb "github.com/ppeymann/Planora.git/proto/room"
	todopb "github.com/ppeymann/Planora.git/proto/todo"
	userpb "github.com/ppeymann/Planora.git/proto/user"
	"gorm.io/gorm"
)

type (
	RoomRepository interface {
		Create(in *roompb.CreateRoomRequest) (*RoomEntity, error)
	}

	RoomEntity struct {
		gorm.Model
		Name      string
		CreatorID uint64
		Creator   userpb.User   `gorm:"foreignKey:CreatorID;references:ID"`
		Users     []userpb.User `gorm:"many2many:room_users;"`
		Todos     []todopb.Todo `gorm:"foreignKey:RoomID;references:ID"`
	}
)
