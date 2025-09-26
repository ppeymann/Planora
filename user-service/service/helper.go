package service

import (
	"context"
	"encoding/json"
	"time"

	"github.com/ppeymann/Planora.git/pkg/auth"
	"github.com/ppeymann/Planora.git/pkg/env"
	userpb "github.com/ppeymann/Planora.git/proto/user"
	"github.com/ppeymann/Planora/user/models"
)

func (s *UserServiceServer) SignUpService(data []byte) (*models.TokenBundlerOutput, error) {
	req := &userpb.SignUpRequest{}

	_ = json.Unmarshal(data, req)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	user, err := s.SignUp(ctx, req)
	if err != nil {
		return nil, err
	}

	paseto, err := auth.NewPasetoMaker(env.GetEnv("JWT", ""))
	if err != nil {
		return nil, err
	}

	tokenClaims := &auth.Claims{
		Subject:   uint(user.Model.Id),
		Issuer:    "www.planora.com",
		Audience:  "www.planora.com",
		IssuedAt:  time.Unix(518400, 0),
		ExpiredAt: time.Now().Add(1036800 * time.Minute).UTC(),
	}

	tokenStr, err := paseto.CreateToken(tokenClaims)
	if err != nil {
		return nil, err
	}

	return &models.TokenBundlerOutput{
		Token:  tokenStr,
		Expire: tokenClaims.ExpiredAt,
	}, nil
}

func (s *UserServiceServer) LoginService(data []byte) (*models.TokenBundlerOutput, error) {
	req := &userpb.LoginRequest{}

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	user, err := s.Login(ctx, req)
	if err != nil {
		return nil, err
	}

	paseto, err := auth.NewPasetoMaker(env.GetEnv("JWT", ""))
	if err != nil {
		return nil, err
	}

	tokenClaims := &auth.Claims{
		Subject:   uint(user.Model.Id),
		Issuer:    "www.planora.com",
		Audience:  "www.planora.com",
		IssuedAt:  time.Unix(518400, 0),
		ExpiredAt: time.Now().Add(1036800 * time.Minute).UTC(),
	}

	tokenStr, err := paseto.CreateToken(tokenClaims)
	if err != nil {
		return nil, err
	}

	return &models.TokenBundlerOutput{
		Token:  tokenStr,
		Expire: tokenClaims.ExpiredAt,
	}, nil
}

func (s *UserServiceServer) AccountService(data []byte) (*userpb.User, error) {
	req := &userpb.AccountRequest{}

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	user, err := s.Account(ctx, req)
	if err != nil {
		return nil, err
	}

	return user, nil
}

func (s *UserServiceServer) GetRoomUsersService(data []byte) (*userpb.GetRoomUsersResponse, error) {
	req := &userpb.GetRoomUsersRequest{}

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	users, err := s.GetRoomUsers(ctx, req)
	if err != nil {
		return nil, err
	}

	return users, nil
}

func (s *UserServiceServer) GetByUsernameService(data []byte) (*userpb.User, error) {
	req := &userpb.GetByUsernameRequest{}

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	user, err := s.GetByUsername(ctx, req)
	if err != nil {
		return nil, err
	}

	return user, nil
}
