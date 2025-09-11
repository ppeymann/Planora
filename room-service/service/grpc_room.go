package service

import (
	"context"

	roompb "github.com/ppeymann/Planora.git/proto/room"
	"github.com/ppeymann/Planora/room/models"
)

type RoomServiceServer struct {
	roompb.UnimplementedRoomServiceServer
	repo models.RoomRepository
}

func NewRoomServiceServer(r models.RoomRepository) *RoomServiceServer {
	return &RoomServiceServer{
		repo: r,
	}
}

func (s *RoomServiceServer) Create(_ context.Context, in *roompb.CreateRoomRequest) (*roompb.Room, error) {
	// room, err := s.repo
	return nil, nil
}
