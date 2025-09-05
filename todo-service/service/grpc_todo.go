package service

import (
	"context"

	todopb "github.com/ppeymann/Planora.git/proto/todo"
	"github.com/ppeymann/Planora/todo/models"
)

type TodoServiceServer struct {
	todopb.UnimplementedTodoServiceServer
	repo models.TodoRepository
}

func NewTodoServiceServer(r models.TodoRepository) *TodoServiceServer {
	return &TodoServiceServer{
		repo: r,
	}
}

func (s *TodoServiceServer) AddTodo(ctx context.Context, in *todopb.AddTodoRequest) (*todopb.Todo, error) {
	todo, err := s.repo.Create(in)
	if err != nil {
		return nil, err
	}

	return &todopb.Todo{
		Model:       models.ToBaseModel(todo),
		Title:       todo.Title,
		Description: todo.Description,
		Status:      string(todo.Status),
		UserId:      uint64(todo.UserID),
	}, nil
}
