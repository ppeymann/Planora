package repository

import (
	roompb "github.com/ppeymann/Planora.git/proto/room"
	"github.com/ppeymann/Planora/room/models"
	"gorm.io/gorm"
)

type roomRepo struct {
	pg       *gorm.DB
	database string
	table    string
}

// Create implements models.RoomRepository.
func (r *roomRepo) Create(in *roompb.CreateRoomRequest) (*models.RoomEntity, error) {
	_ = &models.RoomEntity{
		Name:      in.GetName(),
		CreatorID: in.GetCreatorId(),
	}

	return nil, nil
}

func NewRoomRepo(db *gorm.DB, database string) models.RoomRepository {
	return &roomRepo{
		pg:       db,
		database: database,
		table:    "room_entities",
	}
}
