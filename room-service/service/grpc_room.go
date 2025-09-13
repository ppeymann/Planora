package service

import (
	"context"

	roompb "github.com/ppeymann/Planora.git/proto/room"
	"github.com/ppeymann/Planora/room/models"
	"github.com/ppeymann/Planora/room/utils"
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
	room, err := s.repo.Create(in)
	if err != nil {
		return nil, err
	}

	return &roompb.Room{
		Model:     utils.ToBaseModel(room),
		Name:      room.Name,
		CreatorId: in.CreatorId,
	}, nil
}

func (s *RoomServiceServer) GetUsers(_ context.Context, in *roompb.GetUsersRequest) (*roompb.GetUsersResponse, error) {
	ids, err := s.repo.GetUsers(in.GetRoomId())
	if err != nil {
		return nil, err
	}

	return &roompb.GetUsersResponse{
		UserIds: ids,
	}, err
}

func (s *RoomServiceServer) GetRoom(_ context.Context, in *roompb.GetRoomRequest) (*roompb.GetRoomResponse, error) {
	room, err := s.repo.GetRoom(uint(in.GetRoomId()))
	if err != nil {
		return nil, err
	}

	return &roompb.GetRoomResponse{
		Room: &roompb.Room{
			Model:     utils.ToBaseModel(room),
			Name:      room.Name,
			CreatorId: room.CreatorID,
			UserIds:   utils.ToProtoUintIDs(room.UserIDs),
			TodoIds:   utils.ToProtoUintIDs(room.TodosIDs),
		},
	}, nil
}
