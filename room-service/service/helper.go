package service

import (
	"context"
	"encoding/json"
	"time"

	roompb "github.com/ppeymann/Planora.git/proto/room"
)

func (s *RoomServiceServer) CreateService(data []byte) (*roompb.Room, error) {
	req := &roompb.CreateRoomRequest{}

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	room, err := s.Create(ctx, req)
	if err != nil {
		return nil, err
	}

	return room, nil
}

func (s *RoomServiceServer) GetUsersService(data []byte) (*roompb.GetUsersResponse, error) {
	req := &roompb.GetUsersRequest{}

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	ids, err := s.GetUsers(ctx, req)
	if err != nil {
		return nil, err
	}

	return ids, nil
}
