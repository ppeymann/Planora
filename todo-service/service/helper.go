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

func (s *TodoServiceServer) GetAllTodoService(data []byte) (*todopb.GetAllTodoResponse, error) {
	req := &todopb.GetAllTodoRequest{}

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	todos, err := s.GetAllTodo(ctx, req)
	if err != nil {
		return nil, err
	}

	return todos, nil
}

func (s *TodoServiceServer) ChangeStatusService(data []byte) (*todopb.Todo, error) {
	req := &todopb.ChangeStatusRequest{}

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	todo, err := s.ChangeStatus(ctx, req)
	if err != nil {
		return nil, err
	}

	return todo, err
}

func (s *TodoServiceServer) DeleteTodoService(data []byte) (*todopb.DeleteTodoResponse, error) {
	req := &todopb.DeleteTodoRequest{}

	err := json.Unmarshal(data, req)
	if err != nil {
		return nil, err
	}

	ctx, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	res, err := s.DeleteTodo(ctx, req)
	if err != nil {
		return nil, err
	}

	return res, nil
}
