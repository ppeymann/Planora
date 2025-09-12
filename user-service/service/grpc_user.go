package service

import (
	"context"
	"errors"

	userpb "github.com/ppeymann/Planora.git/proto/user"
	"github.com/ppeymann/Planora/user/models"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type UserServiceServer struct {
	userpb.UnimplementedUserServiceServer
	repo models.UserRepository
}

func NewUserServiceServer(r models.UserRepository) *UserServiceServer {
	return &UserServiceServer{
		repo: r,
	}
}

func (s *UserServiceServer) SignUp(ctx context.Context, in *userpb.SignUpRequest) (*userpb.User, error) {
	user, err := s.repo.Create(in)
	if err != nil {
		return nil, err
	}

	return &userpb.User{
		Model: &userpb.BaseModel{
			Id:         uint64(user.ID),
			CreatedAt:  timestamppb.New(user.CreatedAt),
			UpdatedeAt: timestamppb.New(user.UpdatedAt),
		},
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}

func (s *UserServiceServer) Login(ctx context.Context, in *userpb.LoginRequest) (*userpb.User, error) {
	user, err := s.repo.Find(in.GetUsername())
	if err != nil {
		return nil, err
	}

	if user.Password != in.GetPassword() {
		return nil, errors.New("permission denied")
	}

	return &userpb.User{
		Model: &userpb.BaseModel{
			Id:         uint64(user.ID),
			CreatedAt:  timestamppb.New(user.CreatedAt),
			UpdatedeAt: timestamppb.New(user.UpdatedAt),
		},
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}

func (s *UserServiceServer) Account(ctx context.Context, in *userpb.AccountRequest) (*userpb.User, error) {
	user, err := s.repo.FindByID(uint(in.GetId()))
	if err != nil {
		return nil, err
	}

	return &userpb.User{
		Model: &userpb.BaseModel{
			Id:         uint64(user.ID),
			CreatedAt:  timestamppb.New(user.CreatedAt),
			UpdatedeAt: timestamppb.New(user.UpdatedAt),
		},
		Username:  user.Username,
		Email:     user.Email,
		FirstName: user.FirstName,
		LastName:  user.LastName,
	}, nil
}

func (s *UserServiceServer) GetRoomUsers(ctx context.Context, in *userpb.GetRoomUsersRequest) (*userpb.GetRoomUsersResponse, error) {
	users, err := s.repo.GetRoomUsers(in.GetIds())
	if err != nil {
		return nil, err
	}

	var usersResponse []*userpb.User
	for _, user := range users {
		u := &userpb.User{
			Model:     models.ToBaseModel(user),
			Username:  user.Username,
			Email:     user.Email,
			FirstName: user.FirstName,
			LastName:  user.LastName,
		}

		usersResponse = append(usersResponse, u)
	}

	return &userpb.GetRoomUsersResponse{
		Users: usersResponse,
	}, nil
}
