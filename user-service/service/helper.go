package service

import (
	"context"
	"encoding/json"
	"time"

	userpb "github.com/ppeymann/Planora.git/proto/user"
)

func (s *UserServiceServer) SignUpService(data []byte) (*userpb.User, error) {
	req := &userpb.SignUpRequest{}

	_ = json.Unmarshal(data, req)

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	return s.SignUp(ctx, req)
}
