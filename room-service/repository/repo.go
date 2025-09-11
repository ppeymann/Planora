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

// Migrate implements models.RoomRepository.
func (r *roomRepo) Migrate() error {
	return r.pg.AutoMigrate(&models.RoomEntity{})
}

// Model implements models.RoomRepository.
func (r *roomRepo) Model() *gorm.DB {
	return r.pg.Model(&models.RoomEntity{})
}

// Name implements models.RoomRepository.
func (r *roomRepo) Name() string {
	return r.table
}

// Create implements models.RoomRepository.
func (r *roomRepo) Create(in *roompb.CreateRoomRequest) (*models.RoomEntity, error) {
	room := &models.RoomEntity{
		Name:      in.GetName(),
		CreatorID: in.GetCreatorId(),
	}

	err := r.Model().Create(room).Error
	if err != nil {
		return nil, err
	}

	return room, nil
}

func NewRoomRepo(db *gorm.DB, database string) models.RoomRepository {
	return &roomRepo{
		pg:       db,
		database: database,
		table:    "room_entities",
	}
}
