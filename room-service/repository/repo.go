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

// Update implements models.RoomRepository.
func (r *roomRepo) Update(room *models.RoomEntity) error {
	return r.pg.Save(room).Error
}

// GetRoom implements models.RoomRepository.
func (r *roomRepo) GetRoom(roomID uint) (*models.RoomEntity, error) {
	room := &models.RoomEntity{}

	if err := r.Model().Where("id = ?", roomID).First(room).Error; err != nil {
		return nil, err
	}

	return room, nil
}

// GetUsers implements models.RoomRepository.
func (r *roomRepo) GetUsers(roomID uint64) ([]uint64, error) {
	room := &models.RoomEntity{}

	err := r.Model().Where("id = ?", uint(roomID)).First(room).Error
	if err != nil {
		return nil, err
	}

	userIDs := make([]uint64, len(room.UserIDs))
	for i, id := range room.UserIDs {
		userIDs[i] = uint64(id)
	}

	return userIDs, nil
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
