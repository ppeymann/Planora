package service

import (
	"context"
	"encoding/json"
	"time"

	todopb "github.com/ppeymann/Planora.git/proto/todo"
)

func (s *TodoServiceServer) AddTodoService(data []byte) (*todopb.Todo, error) {
	req := &todopb.AddTodoRequest{}

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	todo, err := s.AddTodo(ctx, req)
	if err != nil {
		return nil, err
	}

	return todo, nil
}

func (s *TodoServiceServer) UpdateTodoService(data []byte) (*todopb.Todo, error) {
	req := &todopb.UpdateTodoRequest{}

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	todo, err := s.UpdateTodo(ctx, req)
	if err != nil {
		return nil, err
	}

	return todo, nil
}
