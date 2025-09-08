package repository

import (
	"github.com/ppeymann/Planora/room/models"
	"gorm.io/gorm"
)

type roomRepo struct {
	pg       *gorm.DB
	database string
	table    string
}

func NewRoomRepo(db *gorm.DB, database string) models.RoomRepository {
	return &roomRepo{
		pg:       db,
		database: database,
		table:    "room_entities",
	}
}
